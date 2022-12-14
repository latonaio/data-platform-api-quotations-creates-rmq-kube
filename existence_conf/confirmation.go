package existence_conf

import (
	"context"
	dpfm_api_input_reader "data-platform-api-quotations-creates-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-quotations-creates-rmq-kube/config"
	"data-platform-api-quotations-creates-rmq-kube/sub_func_complementer"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type ExistenceConf struct {
	ctx context.Context

	c             *config.Conf
	queueToMapper exconfQueueMapper
	rmq           *rabbitmq.RabbitmqClient
}

func NewExistenceConf(ctx context.Context, c *config.Conf, rmq *rabbitmq.RabbitmqClient) *ExistenceConf {
	return &ExistenceConf{
		ctx:           ctx,
		c:             c,
		queueToMapper: NewExconfQueueMapper(c),
		rmq:           rmq,
	}
}

// Confirm returns existenceMap, allExist, err
func (c *ExistenceConf) Conf(data *dpfm_api_input_reader.SDC, ssdc *sub_func_complementer.SDC, l *logger.Logger) (allExist bool, errs []error) {
	var res string
	var resMsg string
	var err error
	existenceMap := make([]bool, 0, 5)
	wg := sync.WaitGroup{}
	mtx := &sync.Mutex{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		res, err = c.bpExistenceConf(*data.Quotations.Buyer, data, &existenceMap, mtx, l)
		if err != nil {
			mtx.Lock()
			errs = append(errs, err)
			mtx.Unlock()
		}
		if res != "" {
			resMsg = res
		}
	}()
	go func() {
		defer wg.Done()
		res, err = c.bpExistenceConf(*data.Quotations.Seller, data, &existenceMap, mtx, l)
		if errs != nil {
			mtx.Lock()
			errs = append(errs, err)
			mtx.Unlock()
		}
		if res != "" {
			resMsg = res
		}
	}()

	wg.Wait()

	if len(errs) != 0 {
		return false, errs
	}

	ssdc.ExconfResult = getBoolPtr(true)
	for _, v := range existenceMap {
		if v {
			continue
		}
		ssdc.ExconfResult = getBoolPtr(false)
		ssdc.ExconfError = resMsg
		return false, nil
	}
	return true, nil
}

func confKeyExistence(res map[string]interface{}) bool {
	if res == nil {
		return false
	}
	raw, ok := res["ExistenceConf"]
	exist := fmt.Sprintf("%v", raw)
	if ok {
		return strings.ToLower(exist) == "true"
	}

	return false
}

func (c *ExistenceConf) bpExistenceConf(bpID int, data *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	var resMsg string
	key := "BusinessPartnerGeneral"
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()

	}()
	b, _ := json.Marshal(data)
	req := BusinessPartnerReq{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return resMsg, xerrors.Errorf("Unmarshal error: %w", err)
	}

	req.BusinessPartner.BusinessPartner = bpID
	res, err := c.rmq.SessionKeepRequest(nil, c.queueToMapper[key], req)
	if err != nil {
		return resMsg, xerrors.Errorf("response error: %w", err)
	}
	res.Success()
	exist = confKeyExistence(res.Data())
	log.Info(res.Data())
	if !confKeyExistence(res.Data()) {
		resMsg = fmt.Sprintf("BusinessPartner:%d ???????????????????????????????????????", int(res.Data()["BusinessPartner"].(float64)))
	}

	return resMsg, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}
