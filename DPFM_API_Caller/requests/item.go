package requests

type Item struct {
	Quotation                     int      `json:"Quotation"`
	QuotationItem                 int      `json:"QuotationItem"`
	QuotationItemCategory         *string  `json:"QuotationItemCategory"`
	QuotationItemText             *string  `json:"QuotationItemText"`
	Product                       *string  `json:"Product"`
	QuotationQuantity             *float32 `json:"QuotationQuantity"`
	QuotationQuantityUnit         *string  `json:"QuotationQuantityUnit"`
	ItemOrderProbabilityInPercent *float32 `json:"ItemOrderProbabilityInPercent"`
	ItemGrossWeight               *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                 *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                *string  `json:"ItemWeightUnit"`
	TransactionCurrency           *string  `json:"TransactionCurrency"`
	BussinessParnterCurrency      *string  `json:"BussinessParnterCurrency"`
	NetAmount                     *float32 `json:"NetAmount"`
	ProductGroup                  *string  `json:"ProductGroup"`
	ProductPricingGroup           *string  `json:"ProductPricingGroup"`
	Batch                         *string  `json:"Batch"`
	IssuingPlant                  *string  `json:"IssuingPlant"`
	IssuingPlantStorageLocation   *string  `json:"IssuingPlantStorageLocation"`
	ReceivingPlant                *string  `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation *string  `json:"ReceivingPlantStorageLocation"`
	Incoterms                     *string  `json:"Incoterms"`
	PaymentTerms                  *string  `json:"PaymentTerms"`
	ProductTaxClassification1     *string  `json:"ProductTaxClassification1"`
	Project                       *string  `json:"Project"`
	ProfitCenter                  *string  `json:"ProfitCenter"`
	ReferenceInquiry              *int     `json:"ReferenceInquiry"`
	ReferenceInquiryItem          *int     `json:"ReferenceInquiryItem"`
}
