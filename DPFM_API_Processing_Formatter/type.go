package dpfm_api_processing_formatter

type HeaderUpdates struct {
	QuoationType                   string   `json:"QuoationType"`
	Buyer                          int      `json:"Buyer"`
	Seller                         int      `json:"Seller"`
	BillToParty                    *int     `json:"BillToParty"`
	BillFromParty                  *int     `json:"BillFromParty"`
	Payer                          *int     `json:"Payer"`
	Payee                          *int     `json:"Payee"`
	ContractType                   *string  `json:"ContractType"`
	BindingPeriodValidityStartDate *string  `json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate   *string  `json:"BindingPeriodValidityEndDate"`
	OrderValidityStartDate         *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate           *string  `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate         *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate           *string  `json:"InvoicePeriodEndDate"`
	TotalNetAmount                 float32  `json:"TotalNetAmount"`
	TotalTaxAmount                 float32  `json:"TotalTaxAmount"`
	TotalGrossAmount               float32  `json:"TotalGrossAmount"`
	TransactionCurrency            string   `json:"TransactionCurrency"`
	PricingDate                    string   `json:"PricingDate"`
	PriceDetnExchangeRate          *float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate          string   `json:"RequestedDeliveryDate"`
	OrderProbabilityInPercent      *float32 `json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount         *float32 `json:"ExpectedOrderNetAmount"`
	Incoterms                      *string  `json:"Incoterms"`
	PaymentTerms                   string   `json:"PaymentTerms"`
	PaymentMethod                  string   `json:"PaymentMethod"`
	AccountingExchangeRate         *float32 `json:"AccountingExchangeRate"`
	InvoiceDocumentDate            string   `json:"InvoiceDocumentDate"`
	HeaderText                     *string  `json:"HeaderText"`
	HeaderBlockStatus              *bool    `json:"HeaderBlockStatus"`
}

type PartnerUpdates struct {
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
}

type AddressUpdates struct {
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

type ItemUpdates struct {
	QuotationItemCategory                   string   `json:"QuotationItemCategory"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	QuotationItemText                       string   `json:"QuotationItemText"`
	QuotationItemTextByBuyer                string   `json:"QuotationItemTextByBuyer"`
	QuotationItemTextBySeller               string   `json:"QuotationItemTextBySeller"`
	Product                                 string   `json:"Product"`
	RequestedDeliveryDate                   string   `json:"RequestedDeliveryDate"`
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
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string   `json:"DefinedTaxClassification"`
	ProductAccountAssignmentGroup           string   `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	Project                                 *string  `json:"Project"`
	ItemBlockStatus                         *bool    `json:"ItemBlockStatus"`
}

type ItemPricingElementUpdates struct {
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
