package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string     `json:"connection_key"`
	Result           bool       `json:"result"`
	RedisKey         string     `json:"redis_key"`
	Filepath         string     `json:"filepath"`
	APIStatusCode    int        `json:"api_status_code"`
	RuntimeSessionID string     `json:"runtime_session_id"`
	BusinessPartner  *int       `json:"business_partner"`
	ServiceLabel     string     `json:"service_label"`
	Quotations       Quotations `json:"Quotations"`
	APISchema        string     `json:"api_schema"`
	Accepter         []string   `json:"accepter"`
	Deleted          bool       `json:"deleted"`
}

type Quotations struct {
	Quotation                      *int     `json:"Quotation"`
	QuotationDate                  *string  `json:"QuotationDate"`
	QuoationType                   string   `json:"QuoationType"`
	Buyer                          *int     `json:"Buyer"`
	Seller                         *int     `json:"Seller"`
	CreationDate                   *string  `json:"CreationDate"`
	LastChangeDate                 *string  `json:"LastChangeDate"`
	ContractType                   string   `json:"ContractType"`
	VaridityStartDate              *string  `json:"VaridityStartDate"`
	ValidityEndDate                *string  `json:"ValidityEndDate"`
	InvoiceScheduleStartDate       *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string  `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                 *float32 `json:"TotalNetAmount"`
	TotalTaxAmount                 *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount               *float32 `json:"TotalGrossAmount"`
	TransactionCurrency            string   `json:"TransactionCurrency"`
	PricingDate                    *string  `json:"PricingDate"`
	RequestedDeliveryDate          *string  `json:"RequestedDeliveryDate"`
	BindingPeriodValidityStartDate *string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   *string  `json:"BindingPeriodValidityEndDate"`
	Incoterms                      string   `json:"Incoterms"`
	PaymentTerms                   string   `json:"PaymentTerms"`
	PaymentMethod                  string   `json:"PaymentMethod"`
	AccountingExchangeRate         *float32 `json:"AccountingExchangeRate"`
	BillingDocumentDate            *string  `json:"BillingDocumentDate"`
	HeaderText                     string   `json:"HeaderText"`
	ReferenceInquiry               *int     `json:"ReferenceInquiry"`
}
