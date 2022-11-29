package dpfm_api_output_formatter

type Header struct {
	Quotation                      int     `json:"Quotation"`
	QuotationDate                  string  `json:"QuotationDate"`
	QuoationType                   string  `json:"QuoationType"`
	Buyer                          int     `json:"Buyer"`
	Seller                         int     `json:"Seller"`
	CreationDate                   string  `json:"CreationDate"`
	LastChangeDate                 string  `json:"LastChangeDate"`
	ContractType                   string  `json:"ContractType"`
	VaridityStartDate              string  `json:"VaridityStartDate"`
	ValidityEndDate                string  `json:"ValidityEndDate"`
	InvoiceScheduleStartDate       string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         string  `json:"InvoiceScheduleEndDate"`
	TotalNetAmount                 float32 `json:"TotalNetAmount"`
	TotalTaxAmount                 float32 `json:"TotalTaxAmount"`
	TotalGrossAmount               float32 `json:"TotalGrossAmount"`
	TransactionCurrency            string  `json:"TransactionCurrency"`
	PricingDate                    string  `json:"PricingDate"`
	RequestedDeliveryDate          string  `json:"RequestedDeliveryDate"`
	BindingPeriodValidityStartDate string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   string  `json:"BindingPeriodValidityEndDate"`
	Incoterms                      string  `json:"Incoterms"`
	PaymentTerms                   string  `json:"PaymentTerms"`
	PaymentMethod                  string  `json:"PaymentMethod"`
	AccountingExchangeRate         float32 `json:"AccountingExchangeRate"`
	BillingDocumentDate            string  `json:"BillingDocumentDate"`
	HeaderText                     string  `json:"HeaderText"`
	ReferenceInquiry               int     `json:"ReferenceInquiry"`
}
