package dpfm_api_input_reader

import (
	"data-platform-api-quotations-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Quotations
	return &requests.Header{
		Quotation:                      data.Quotation,
		QuotationDate:                  data.QuotationDate,
		QuoationType:                   data.QuoationType,
		Buyer:                          data.Buyer,
		Seller:                         data.Seller,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		ContractType:                   data.ContractType,
		VaridityStartDate:              data.VaridityStartDate,
		ValidityEndDate:                data.ValidityEndDate,
		InvoiceScheduleStartDate:       data.InvoiceScheduleStartDate,
		InvoiceScheduleEndDate:         data.InvoiceScheduleEndDate,
		TotalNetAmount:                 data.TotalNetAmount,
		TotalTaxAmount:                 data.TotalTaxAmount,
		TotalGrossAmount:               data.TotalGrossAmount,
		TransactionCurrency:            data.TransactionCurrency,
		PricingDate:                    data.PricingDate,
		RequestedDeliveryDate:          data.RequestedDeliveryDate,
		BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
		Incoterms:                      data.Incoterms,
		PaymentTerms:                   data.PaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		AccountingExchangeRate:         data.AccountingExchangeRate,
		BillingDocumentDate:            data.BillingDocumentDate,
		HeaderText:                     data.HeaderText,
		ReferenceInquiry:               data.ReferenceInquiry,
	}
}
