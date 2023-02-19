package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-quotations-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		QuoationType:                   *data.QuoationType,
		Buyer:                          *data.Buyer,
		Seller:                         *data.Seller,
		BillToParty:                    data.BillToParty,
		BillFromParty:                  data.BillFromParty,
		Payer:                          data.Payer,
		Payee:                          data.Payee,
		ContractType:                   data.ContractType,
		BindingPeriodValidityStartDate: data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:   data.BindingPeriodValidityEndDate,
		OrderValidityStartDate:         data.OrderValidityStartDate,
		OrderValidityEndDate:           data.OrderValidityEndDate,
		InvoicePeriodStartDate:         data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:           data.InvoicePeriodEndDate,
		TotalNetAmount:                 *data.TotalNetAmount,
		TotalTaxAmount:                 *data.TotalTaxAmount,
		TotalGrossAmount:               *data.TotalGrossAmount,
		TransactionCurrency:            *data.TransactionCurrency,
		PricingDate:                    *data.PricingDate,
		PriceDetnExchangeRate:          data.PriceDetnExchangeRate,
		RequestedDeliveryDate:          *data.RequestedDeliveryDate,
		OrderProbabilityInPercent:      data.OrderProbabilityInPercent,
		ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
		Incoterms:                      data.Incoterms,
		PaymentTerms:                   *data.PaymentTerms,
		PaymentMethod:                  *data.PaymentMethod,
		InvoiceDocumentDate:            *data.InvoiceDocumentDate,
		HeaderText:                     data.HeaderText,
		HeaderBlockStatus:              data.HeaderBlockStatus,
	}
}

func ConvertToAddressUpdates(address dpfm_api_input_reader.Address) *AddressUpdates {
	data := address

	return &AddressUpdates{
		AddressID:   data.AddressID,
		PostalCode:  data.PostalCode,
		LocalRegion: data.LocalRegion,
		Country:     data.Country,
		District:    data.District,
		StreetName:  data.StreetName,
		CityName:    data.CityName,
		Building:    data.Building,
		Floor:       data.Floor,
		Room:        data.Room,
	}
}

func ConvertToItemUpdates(item dpfm_api_input_reader.Item) *ItemUpdates {
	data := item

	return &ItemUpdates{
		QuotationItemCategory:                   *data.QuotationItemCategory,
		SupplyChainRelationshipID:               *data.SupplyChainRelationshipID,
		QuotationItemText:                       *data.QuotationItemText,
		QuotationItemTextByBuyer:                *data.QuotationItemTextByBuyer,
		QuotationItemTextBySeller:               *data.QuotationItemTextBySeller,
		Product:                                 *data.Product,
		RequestedDeliveryDate:                   *data.RequestedDeliveryDate,
		DeliveryUnit:                            *data.DeliveryUnit,
		QuotationQuantityInBaseUnit:             *data.QuotationQuantityInBaseUnit,
		QuotationQuantityInDeliveryUnit:         *data.QuotationQuantityInDeliveryUnit,
		ItemWeightUnit:                          data.ItemWeightUnit,
		ProductGrossWeight:                      data.ProductGrossWeight,
		ItemGrossWeight:                         data.ItemGrossWeight,
		ProductNetWeight:                        data.ProductNetWeight,
		ItemNetWeight:                           data.ItemNetWeight,
		InternalCapacityQuantity:                data.InternalCapacityQuantity,
		InternalCapacityQuantityUnit:            data.InternalCapacityQuantityUnit,
		TaxAmount:                               data.TaxAmount,
		GrossAmount:                             data.GrossAmount,
		Incoterms:                               data.Incoterms,
		TransactionTaxClassification:            *data.TransactionTaxClassification,
		ProductTaxClassificationBillToCountry:   *data.ProductTaxClassificationBillToCountry,
		ProductTaxClassificationBillFromCountry: *data.ProductTaxClassificationBillFromCountry,
		DefinedTaxClassification:                *data.DefinedTaxClassification,
		ProductAccountAssignmentGroup:           *data.ProductAccountAssignmentGroup,
		PaymentTerms:                            *data.PaymentTerms,
		PaymentMethod:                           *data.PaymentMethod,
		Project:                                 data.Project,
		ItemBlockStatus:                         data.ItemBlockStatus,
	}
}

func ConvertToItemPricingElementUpdates(itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		ConditionRecord:            data.ConditionRecord,
		ConditionSequentialNumber:  data.ConditionSequentialNumber,
		ConditionRateValue:         data.ConditionRateValue,
		TaxCode:                    data.TaxCode,
		ConditionAmount:            data.ConditionAmount,
		ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
	}
}

func ConvertToPartnerUpdates(partner dpfm_api_input_reader.Partner) *PartnerUpdates {
	data := partner

	return &PartnerUpdates{
		BusinessPartnerFullName: data.BusinessPartnerFullName,
		BusinessPartnerName:     data.BusinessPartnerName,
		Organization:            data.Organization,
		Country:                 data.Country,
		Language:                data.Language,
		Currency:                data.Currency,
		ExternalDocumentID:      data.ExternalDocumentID,
	}
}
