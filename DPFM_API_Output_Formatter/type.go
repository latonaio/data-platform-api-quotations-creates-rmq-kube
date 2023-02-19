package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header             *Header               `json:"Header"`
	Partner            *[]Partner            `json:"Partner"`
	Address            *[]Address            `json:"Address"`
	Item               *[]Item               `json:"Item"`
	ItemPricingElement *[]ItemPricingElement `json:"ItemPricingElement"`
}

type Header struct {
	Quotation                        int      `json:"Quotation"`
	QuotationDate                    string   `json:"QuotationDate"`
	QuoationType                     string   `json:"QuoationType"`
	QuoationStatus                   string   `json:"QuoationStatus"`
	SupplyChainRelationshipID        int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int      `json:"Buyer"`
	Seller                           int      `json:"Seller"`
	BillToParty                      *int     `json:"BillToParty"`
	BillFromParty                    *int     `json:"BillFromParty"`
	BillToCountry                    *string  `json:"BillToCountry"`
	BillFromCountry                  *string  `json:"BillFromCountry"`
	Payer                            *int     `json:"Payer"`
	Payee                            *int     `json:"Payee"`
	CreationDate                     string   `json:"CreationDate"`
	LastChangeDate                   string   `json:"LastChangeDate"`
	ContractType                     *string  `json:"ContractType"`
	BindingPeriodValidityStartDate   *string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate     *string  `json:"BindingPeriodValidityEndDate"`
	OrderValidityStartDate           *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate             *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32  `json:"TotalNetAmount"`
	TotalTaxAmount                   float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32  `json:"TotalGrossAmount"`
	TransactionCurrency              string   `json:"TransactionCurrency"`
	PricingDate                      string   `json:"PricingDate"`
	PriceDetnExchangeRate            *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate            string   `json:"RequestedDeliveryDate"`
	OrderProbabilityInPercent        *float32 `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount           *float32 `json:"ExpectedOrderNetAmount"`
	Incoterms                        *string  `json:"Incoterms"`
	PaymentTerms                     string   `json:"PaymentTerms"`
	PaymentMethod                    string   `json:"PaymentMethod"`
	ReferenceDocument                *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem            *int     `json:"ReferenceDocumentItem"`
	AccountAssignmentGroup           string   `json:"AccountAssignmentGroup"`
	AccountingExchangeRate           *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate              string   `json:"InvoiceDocumentDate"`
	IsExportImport                   *bool    `json:"IsExportImport"`
	HeaderText                       *string  `json:"HeaderText"`
	HeaderBlockStatus                *bool    `json:"HeaderBlockStatus"`
	IsCancelled                      *bool    `json:"IsCancelled"`
	IsMarkedForDeletion              *bool    `json:"IsMarkedForDeletion"`
}

type Item struct {
	Quotation                               int      `json:"Quotation"`
	QuotationItem                           int      `json:"QuotationItem"`
	QuotationItemCategory                   string   `json:"QuotationItemCategory"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	QuotationItemText                       string   `json:"QuotationItemText"`
	QuotationItemTextByBuyer                string   `json:"QuotationItemTextByBuyer"`
	QuotationItemTextBySeller               string   `json:"QuotationItemTextBySeller"`
	Product                                 string   `json:"Product"`
	ProductStandardID                       string   `json:"ProductStandardID"`
	ProductGroup                            *string  `json:"ProductGroup"`
	BaseUnit                                string   `json:"BaseUnit"`
	PricingDate                             string   `json:"PricingDate"`
	PriceDetnExchangeRate                   *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                   string   `json:"RequestedDeliveryDate"`
	CreationDate                            string   `json:"CreationDate"`
	LastChangeDate                          string   `json:"LastChangeDate"`
	DeliveryUnit                            string   `json:"DeliveryUnit"`
	ServicesRenderingDate                   *string  `json:"ServicesRenderingDate"`
	QuotationQuantityInBaseUnit             float32  `json:"QuotationQuantityInBaseUnit"`
	QuotationQuantityInDeliveryUnit         float32  `json:"QuotationQuantityInDeliveryUnit"`
	ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
	ProductGrossWeight                      *float32 `json:"ProductGrossWeight"`
	ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
	ProductNetWeight                        *float32 `json:"ProductNetWeight"`
	ItemNetWeight                           *float32 `json:"ItemNetWeight"`
	InternalCapacityQuantity                *float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit            *string  `json:"InternalCapacityQuantityUnit"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string   `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                  string   `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	Project                                 *string  `json:"Project"`
	AccountingExchangeRate                  *float32 `json:"AccountingExchangeRate"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	CountryOfOrigin                         *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string  `json:"CountryOfOriginLanguage"`
	ItemBlockStatus                         *bool    `json:"ItemBlockStatus"`
	IsCancelled                             *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool    `json:"IsMarkedForDeletion"`
}

type ItemPricingElement struct {
	Quotation                  int      `json:"Quotation"`
	QuotationItem              int      `json:"QuotationItem"`
	SupplyChainRelationshipID  int      `json:"SupplyChainRelationshipID"`
	Buyer                      int      `json:"Buyer"`
	Seller                     int      `json:"Seller"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type Partner struct {
	Quotation               int     `json:"Quotation"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
	Quotation   int     `json:"Quotation"`
	AddressID   int     `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}
