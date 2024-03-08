package requests

type Item struct {
	Quotation								int		`json:"Quotation"`
	QuotationItem							int		`json:"QuotationItem"`
	QuotationItemCategory					string	`json:"QuotationItemCategory"`
	SupplyChainRelationshipID				int		`json:"SupplyChainRelationshipID"`
	Buyer									int		`json:"Buyer"`
	Seller									int		`json:"Seller"`
	QuotationItemText						string	`json:"QuotationItemText"`
	QuotationItemTextByBuyer				string	`json:"QuotationItemTextByBuyer"`
	QuotationItemTextBySeller				string	`json:"QuotationItemTextBySeller"`
	Product									string	`json:"Product"`
	SizeOrDimensionText                     *string  `json:"SizeOrDimensionText"`
	ProductStandardID						*string	`json:"ProductStandardID"`
	ProductGroup							*string	`json:"ProductGroup"`
	ProductSpecification                    *string `json:"ProductSpecification"`
	MarkingOfMaterial                       *string `json:"MarkingOfMaterial"`
	BaseUnit								string	`json:"BaseUnit"`
	PricingDate								string	`json:"PricingDate"`
	PriceDetnExchangeRate					*float32 `json:"PriceDetnExchangeRate"`
	RequestedDeliveryDate					string	`json:"RequestedDeliveryDate"`
	DeliveryUnit							string	`json:"DeliveryUnit"`
	ServicesRenderingDate					*string	`json:"ServicesRenderingDate"`
	QuotationQuantityInBaseUnit				float32	`json:"QuotationQuantityInBaseUnit"`
	QuotationQuantityInDeliveryUnit			float32	`json:"QuotationQuantityInDeliveryUnit"`
	ItemWeightUnit							*string	`json:"ItemWeightUnit"`
	ProductGrossWeight						*float32 `json:"ProductGrossWeight"`
	ItemGrossWeight							*float32 `json:"ItemGrossWeight"`
	ProductNetWeight						*float32 `json:"ProductNetWeight"`
	ItemNetWeight							*float32 `json:"ItemNetWeight"`
	InternalCapacityQuantity				*float32 `json:"InternalCapacityQuantity"`
	InternalCapacityQuantityUnit			*string	 `json:"InternalCapacityQuantityUnit"`
	NetAmount								float32 `json:"NetAmount"`
	TaxAmount								float32 `json:"TaxAmount"`
	GrossAmount								float32 `json:"GrossAmount"`
	InspectionPlantBusinessPartner          *int    `json:"InspectionPlantBusinessPartner"`
	InspectionPlant                         *string `json:"InspectionPlant"`
	InspectionPlan                          *int    `json:"InspectionPlan"`
	InspectionLot                           *int    `json:"InspectionLot"`
	Incoterms								*string	`json:"Incoterms"`
	TransactionTaxClassification			string	`json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry	*string	`json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry	*string	`json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification				*string	`json:"DefinedTaxClassification"`
	AccountAssignmentGroup					string	`json:"AccountAssignmentGroup"`
	ProductAccountAssignmentGroup			string	`json:"ProductAccountAssignmentGroup"`
	PaymentTerms							string	`json:"PaymentTerms"`
	PaymentMethod							string	`json:"PaymentMethod"`
	Contract		                 		*int    `json:"Contract"`
	ContractItem	                 		*int    `json:"ContractItem"`
	Project									*int	`json:"Project"`
	WBSElement								*int	`json:"WBSElement"`
	AccountingExchangeRate					*float32 `json:"AccountingExchangeRate"`
	ReferenceDocument						*int	`json:"ReferenceDocument"`
	ReferenceDocumentItem					*int	`json:"ReferenceDocumentItem"`
	TaxCode									*string	`json:"TaxCode"`
	TaxRate									*float32 `json:"TaxRate"`
	CountryOfOrigin							*string	`json:"CountryOfOrigin"`
	CountryOfOriginLanguage					*string	`json:"CountryOfOriginLanguage"`
	ItemBlockStatus							*bool	`json:"ItemBlockStatus"`
	ExternalReferenceDocument               *string `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem           *string `json:"ExternalReferenceDocumentItem"`
	CreationDate							string	`json:"CreationDate"`
	LastChangeDate							string	`json:"LastChangeDate"`
	IsCancelled								*bool	`json:"IsCancelled"`
	IsMarkedForDeletion						*bool	`json:"IsMarkedForDeletion"`
}
