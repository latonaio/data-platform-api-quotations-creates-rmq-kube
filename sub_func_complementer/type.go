package sub_func_complementer

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Header Header `json:"Header"`
}

type Header struct {
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
