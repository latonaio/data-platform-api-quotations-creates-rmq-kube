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
	ConnectionKey    string   `json:"connection_key"`
	Result           bool     `json:"result"`
	RedisKey         string   `json:"redis_key"`
	Filepath         string   `json:"filepath"`
	APIStatusCode    int      `json:"api_status_code"`
	RuntimeSessionID string   `json:"runtime_session_id"`
	BusinessPartner  *int     `json:"business_partner"`
	ServiceLabel     string   `json:"service_label"`
	APIType          string   `json:"api_type"`
	Header           Header   `json:"Quotations"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	Deleted          bool     `json:"deleted"`
}

type Header struct {
	Quotation							int			`json:"Quotation"`
	QuotationDate						*string		`json:"QuotationDate"`
	QuotationType						*string		`json:"QuotationType"`
	QuotationStatus						*string		`json:"QuotationStatus"`
	SupplyChainRelationshipID			*int		`json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID	*int		`json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID	*int		`json:"SupplyChainRelationshipPaymentID"`
	Buyer								*int		`json:"Buyer"`
	Seller								*int		`json:"Seller"`
	BillToParty							*int		`json:"BillToParty"`
	BillFromParty						*int		`json:"BillFromParty"`
	BillToCountry						*int		`json:"BillToCountry"`
	BillFromCountry						*int		`json:"BillFromCountry"`
	Payer								*int		`json:"Payer"`
	Payee								*int		`json:"Payee"`
	ContractType						*string		`json:"ContractType"`
	BindingPeriodValidityStartDate		*string		`json:"BindingPeriodValidityStartDate"`
	BindingPeriodValidityEndDate		*string		`json:"BindingPeriodValidityEndDate"`
	OrderValidityStartDate				*string		`json:"OrderValidityStartDate"`
	OrderValidityEndDate				*string		`json:"OrderValidityEndDate"`
	InvoicePeriodStartDate				*string		`json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate				*string		`json:"InvoicePeriodEndDate"`
	TotalNetAmount						*float32	`json:"TotalNetAmount"`
	TotalTaxAmount						*float32	`json:"TotalTaxAmount"`
	TotalGrossAmount					*float32	`json:"TotalGrossAmount"`
	HeaderOrderIsDefined				*bool		`json:"HeaderOrderIsDefined"`
	TransactionCurrency					*string		`json:"TransactionCurrency"`
	PricingDate							*string		`json:"PricingDate"`
	PriceDetnExchangeRate				*string		`json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate				*string		`json:"RequestedDeliveryDate"`
	OrderProbabilityInPercent			*float32	`json:"OrderProbabilityInPercent"`
	ExpectedOrderNetAmount				*float32	`json:"ExpectedOrderNetAmount"`
	Incoterms							*string		`json:"Incoterms"`
	PaymentTerms						*string		`json:"PaymentTerms"`
	PaymentMethod						*string		`json:"PaymentMethod"`
	Contract		                 	*int 	    `json:"Contract"`
	ContractItem	                 	*int     	`json:"ContractItem"`
	ReferenceDocument					*int		`json:"ReferenceDocument"`
	AccountAssignmentGroup				*string		`json:"AccountAssignmentGroup"`
	AccountingExchangeRate				*float32	`json:"AccountingExchangeRate"`
	InvoiceDocumentDate					*string		`json:"InvoiceDocumentDate"`
	IsExportImport						*bool		`json:"IsExportImport"`
	HeaderText							*bool		`json:"HeaderText"`
	HeaderIsClosed						*bool		`json:"HeaderIsClosed"`
	HeaderBlockStatus					*bool		`json:"HeaderBlockStatus"`
	ExternalReferenceDocument         	*string		`json:"ExternalReferenceDocument"`
	CertificateAuthorityChain        	*string  	`json:"CertificateAuthorityChain"`
	UsageControlChain        		 	*string  	`json:"UsageControlChain"`
	CreationDate						*string		`json:"CreationDate"`
	LastChangeDate						*string		`json:"LastChangeDate"`
	IsCancelled							*bool		`json:"IsCancelled"`
	IsMarkedForDeletion					*bool		`json:"IsMarkedForDeletion"`
	Item                             	[]Item    	`json:"Item"`
	Partner                          	[]Partner 	`json:"Partner"`
	Address                          	[]Address 	`json:"Address"`
}

type Item struct {
	Quotation                               int                  `json:"Quotation"`
	QuotationItem                           int                  `json:"QuotationItem"`
	QuotationItemCategory                   *string              `json:"QuotationItemCategory"`
	SupplyChainRelationshipID               *int                 `json:"SupplyChainRelationshipID"`
	Buyer                                   *int                 `json:"Buyer"`
	Seller                                  *int                 `json:"Seller"`
	QuotationItemText                       *string              `json:"QuotationItemText"`
	QuotationItemTextByBuyer                *string              `json:"QuotationItemTextByBuyer"`
	QuotationItemTextBySeller               *string              `json:"QuotationItemTextBySeller"`
	Product                                 *string              `json:"Product"`
	SizeOrDimensionText                     *string				 `json:"SizeOrDimensionText"`
	ProductStandardID                       *string              `json:"ProductStandardID"`
	ProductGroup                            *string              `json:"ProductGroup"`
	BaseUnit                                *string              `json:"BaseUnit"`
	PricingDate                             *string              `json:"PricingDate"`
	PriceDetnExchangeRate                   *float32             `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate                   *string              `json:"RequestedDeliveryDate"`
	DeliveryUnit                            *float32             `json:"DeliveryUnit"`
	ServicesRenderingDate                   *string              `json:"ServicesRenderingDate"`
	QuotationQuantityInBaseUnit             *string              `json:"QuotationQuantityInBaseUnit"`
	QuotationQuantityInDeliveryUnit         *string              `json:"QuotationQuantityInDeliveryUnit"`
	ItemWeightUnit                          *string              `json:"ItemWeightUnit"`
	ProductGrossWeight                      *float32             `json:"ProductGrossWeight"`
	ItemGrossWeight                         *float32             `json:"ItemGrossWeight"`
	ProductNetWeight                        *float32             `json:"ProductNetWeight"`
	ItemNetWeight                           *float32             `json:"ItemNetWeight"`
	InternalCapacityQuantity                *float32             `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit            *string              `json:"InternalCapacityQuantityUnit"`
	NetAmount                               *float32             `json:"NetAmount"`
	TaxAmount                               *float32             `json:"TaxAmount"`
	GrossAmount                             *float32             `json:"GrossAmount"`
	Incoterms                               *string              `json:"Incoterms"`
	TransactionTaxClassification            *string              `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string              `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string              `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                *string              `json:"DefinedTaxClassification"`
	AccountAssignmentGroup                  *string              `json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup           *string              `json:"ProductAccountAssignmentGroup"`
	PaymentTerms                            *string              `json:"PaymentTerms"`
	PaymentMethod                           *string              `json:"PaymentMethod"`
	Contract		                 		*int     			 `json:"Contract"`
	ContractItem	                 		*int     			 `json:"ContractItem"`
	Project                                 *int                 `json:"Project"`
	WBSElement                              *int                 `json:"WBSElement"`
	AccountingExchangeRate                  *float32             `json:"AccountingExchangeRate"`
	ReferenceDocument                       *int                 `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int                 `json:"ReferenceDocumentItem"`
	TaxCode                                 *string              `json:"TaxCode"`
	TaxRate                                 *float32             `json:"TaxRate"`
	CountryOfOrigin                         *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string              `json:"CountryOfOriginLanguage"`
	ItemBlockStatus                         *bool                `json:"ItemBlockStatus"`
	ExternalReferenceDocument               *string              `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem           *string              `json:"ExternalReferenceDocumentItem"`
	CreationDate                            *string              `json:"CreationDate"`
	LastChangeDate                          *string              `json:"LastChangeDate"`
	IsCancelled                             *bool                `json:"IsCancelled"`
	IsMarkedForDeletion                     *bool                `json:"IsMarkedForDeletion"`
	ItemPricingElement                      []ItemPricingElement `json:"ItemPricingElement"`
}

type ItemPricingElement struct {
	Quotation                  int      `json:"Quotation"`
	QuotationItem              int      `json:"QuotationItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	SupplyChainRelationshipID  *int     `json:"SupplyChainRelationshipID"`
	Buyer                      *int     `json:"Buyer"`
	Seller                     *int     `json:"Seller"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionRateValueUnit     *string  `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity     *int     `json:"ConditionScaleQuantity"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
	CreationDate               *string  `json:"CreationDate"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	IsCancelled                *bool    `json:"IsCancelled"`
	IsMarkedForDeletion        *bool    `json:"IsMarkedForDeletion"`
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
