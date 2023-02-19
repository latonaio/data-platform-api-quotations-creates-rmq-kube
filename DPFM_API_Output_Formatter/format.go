package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-quotations-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-quotations-creates-rmq-kube/DPFM_API_Processing_Formatter"
)

func ConvertToHeaderCreates(SDC *dpfm_api_input_reader.SDC) *Header {
	data := SDC.Header

	header := &Header{
		Quotation:                        data.Quotation,
		QuotationDate:                    *data.QuotationDate,
		QuoationType:                     *data.QuoationType,
		QuoationStatus:                   *data.QuoationStatus,
		SupplyChainRelationshipID:        *data.SupplyChainRelationshipID,
		SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
		Buyer:                            *data.Buyer,
		Seller:                           *data.Seller,
		BillToParty:                      data.BillToParty,
		BillFromParty:                    data.BillFromParty,
		BillToCountry:                    data.BillToCountry,
		BillFromCountry:                  data.BillFromCountry,
		Payer:                            data.Payer,
		Payee:                            data.Payee,
		CreationDate:                     *data.CreationDate,
		LastChangeDate:                   *data.LastChangeDate,
		ContractType:                     data.ContractType,
		BindingPeriodValidityStartDate:   data.BindingPeriodValidityStartDate,
		BindingPeriodValidityEndDate:     data.BindingPeriodValidityEndDate,
		OrderValidityStartDate:           data.OrderValidityStartDate,
		OrderValidityEndDate:             data.OrderValidityEndDate,
		InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
		InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
		TotalNetAmount:                   *data.TotalNetAmount,
		TotalTaxAmount:                   *data.TotalTaxAmount,
		TotalGrossAmount:                 *data.TotalGrossAmount,
		TransactionCurrency:              *data.TransactionCurrency,
		PricingDate:                      *data.PricingDate,
		PriceDetnExchangeRate:            data.PriceDetnExchangeRate,
		RequestedDeliveryDate:            *data.RequestedDeliveryDate,
		OrderProbabilityInPercent:        data.OrderProbabilityInPercent,
		ExpectedOrderNetAmount:           data.ExpectedOrderNetAmount,
		Incoterms:                        data.Incoterms,
		PaymentTerms:                     *data.PaymentTerms,
		PaymentMethod:                    *data.PaymentMethod,
		ReferenceDocument:                data.ReferenceDocument,
		ReferenceDocumentItem:            data.ReferenceDocumentItem,
		AccountAssignmentGroup:           *data.AccountAssignmentGroup,
		AccountingExchangeRate:           data.AccountingExchangeRate,
		InvoiceDocumentDate:              *data.InvoiceDocumentDate,
		IsExportImport:                   data.IsExportImport,
		HeaderText:                       data.HeaderText,
		HeaderBlockStatus:                data.HeaderBlockStatus,
		IsCancelled:                      data.IsCancelled,
		IsMarkedForDeletion:              data.IsMarkedForDeletion,
	}

	return header
}

func ConvertToHeaderUpdates(headerUpdates *dpfm_api_processing_formatter.HeaderUpdates) *Header {
	data := headerUpdates

	header := &Header{
		QuoationType:                   data.QuoationType,
		Buyer:                          data.Buyer,
		Seller:                         data.Seller,
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
		TotalNetAmount:                 data.TotalNetAmount,
		TotalTaxAmount:                 data.TotalTaxAmount,
		TotalGrossAmount:               data.TotalGrossAmount,
		TransactionCurrency:            data.TransactionCurrency,
		PricingDate:                    data.PricingDate,
		PriceDetnExchangeRate:          data.PriceDetnExchangeRate,
		RequestedDeliveryDate:          data.RequestedDeliveryDate,
		OrderProbabilityInPercent:      data.OrderProbabilityInPercent,
		ExpectedOrderNetAmount:         data.ExpectedOrderNetAmount,
		Incoterms:                      data.Incoterms,
		PaymentTerms:                   data.PaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		InvoiceDocumentDate:            data.InvoiceDocumentDate,
		HeaderText:                     data.HeaderText,
		HeaderBlockStatus:              data.HeaderBlockStatus,
	}

	return header
}

func ConvertToItemCreates(SDC *dpfm_api_input_reader.SDC) *[]Item {
	var item []Item

	for _, data := range SDC.Header.Item {
		item = append(item, Item{
			Quotation:                               data.Quotation,
			QuotationItem:                           data.QuotationItem,
			QuotationItemCategory:                   *data.QuotationItemCategory,
			SupplyChainRelationshipID:               *data.SupplyChainRelationshipID,
			QuotationItemText:                       *data.QuotationItemText,
			QuotationItemTextByBuyer:                *data.QuotationItemTextByBuyer,
			QuotationItemTextBySeller:               *data.QuotationItemTextBySeller,
			Product:                                 *data.Product,
			ProductStandardID:                       *data.ProductStandardID,
			ProductGroup:                            data.ProductGroup,
			BaseUnit:                                *data.BaseUnit,
			PricingDate:                             *data.PricingDate,
			PriceDetnExchangeRate:                   data.PriceDetnExchangeRate,
			RequestedDeliveryDate:                   *data.RequestedDeliveryDate,
			CreationDate:                            *data.CreationDate,
			LastChangeDate:                          *data.LastChangeDate,
			DeliveryUnit:                            *data.DeliveryUnit,
			ServicesRenderingDate:                   data.ServicesRenderingDate,
			QuotationQuantityInBaseUnit:             *data.QuotationQuantityInBaseUnit,
			QuotationQuantityInDeliveryUnit:         *data.QuotationQuantityInDeliveryUnit,
			ItemWeightUnit:                          data.ItemWeightUnit,
			ProductGrossWeight:                      data.ProductGrossWeight,
			ItemGrossWeight:                         data.ItemGrossWeight,
			ProductNetWeight:                        data.ProductNetWeight,
			ItemNetWeight:                           data.ItemNetWeight,
			InternalCapacityQuantity:                data.InternalCapacityQuantity,
			InternalCapacityQuantityUnit:            data.InternalCapacityQuantityUnit,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			Incoterms:                               data.Incoterms,
			TransactionTaxClassification:            *data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   *data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: *data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                *data.DefinedTaxClassification,
			AccountAssignmentGroup:                  *data.AccountAssignmentGroup,
			ProductAccountAssignmentGroup:           *data.ProductAccountAssignmentGroup,
			PaymentTerms:                            *data.PaymentTerms,
			PaymentMethod:                           *data.PaymentMethod,
			Project:                                 data.Project,
			AccountingExchangeRate:                  data.AccountingExchangeRate,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
			CountryOfOrigin:                         data.CountryOfOrigin,
			CountryOfOriginLanguage:                 data.CountryOfOriginLanguage,
			ItemBlockStatus:                         data.ItemBlockStatus,
			IsCancelled:                             data.IsCancelled,
			IsMarkedForDeletion:                     data.IsMarkedForDeletion,
		})
	}

	return &item
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) *[]Item {
	var item []Item

	for _, data := range *itemUpdates {
		item = append(item, Item{
			QuotationItemCategory:                   data.QuotationItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			QuotationItemText:                       data.QuotationItemText,
			QuotationItemTextByBuyer:                data.QuotationItemTextByBuyer,
			QuotationItemTextBySeller:               data.QuotationItemTextBySeller,
			Product:                                 data.Product,
			RequestedDeliveryDate:                   data.RequestedDeliveryDate,
			DeliveryUnit:                            data.DeliveryUnit,
			QuotationQuantityInBaseUnit:             data.QuotationQuantityInBaseUnit,
			QuotationQuantityInDeliveryUnit:         data.QuotationQuantityInDeliveryUnit,
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
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			ProductAccountAssignmentGroup:           data.ProductAccountAssignmentGroup,
			PaymentTerms:                            data.PaymentTerms,
			PaymentMethod:                           data.PaymentMethod,
			Project:                                 data.Project,
			ItemBlockStatus:                         data.ItemBlockStatus,
		})
	}

	return &item
}

func ConvertToItemPricingElementCreates(SDC *dpfm_api_input_reader.SDC) *[]ItemPricingElement {
	var itemPricingElement []ItemPricingElement

	for _, itemData := range SDC.Header.Item {
		for _, data := range itemData.ItemPricingElement {
			itemPricingElement = append(itemPricingElement, ItemPricingElement{
				Quotation:                  data.Quotation,
				QuotationItem:              data.QuotationItem,
				SupplyChainRelationshipID:  data.SupplyChainRelationshipID,
				Buyer:                      data.Buyer,
				Seller:                     data.Seller,
				PricingProcedureCounter:    data.PricingProcedureCounter,
				ConditionRecord:            data.ConditionRecord,
				ConditionSequentialNumber:  data.ConditionSequentialNumber,
				ConditionType:              data.ConditionType,
				PricingDate:                data.PricingDate,
				ConditionRateValue:         data.ConditionRateValue,
				ConditionCurrency:          data.ConditionCurrency,
				ConditionQuantity:          data.ConditionQuantity,
				ConditionQuantityUnit:      data.ConditionQuantityUnit,
				TaxCode:                    data.TaxCode,
				ConditionAmount:            data.ConditionAmount,
				TransactionCurrency:        data.TransactionCurrency,
				ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
			})
		}
	}
	return &itemPricingElement
}

func ConvertToItemPricingElementUpdates(itemPricingElementUpdates *[]dpfm_api_processing_formatter.ItemPricingElementUpdates) *[]ItemPricingElement {
	var itemPricingElement []ItemPricingElement

	for _, data := range *itemPricingElementUpdates {
		itemPricingElement = append(itemPricingElement, ItemPricingElement{
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionRateValue:         data.ConditionRateValue,
			TaxCode:                    data.TaxCode,
			ConditionAmount:            data.ConditionAmount,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}

	return &itemPricingElement
}

func ConvertToPartnerCreates(SDC *dpfm_api_input_reader.SDC) *[]Partner {
	var partner []Partner

	for _, data := range SDC.Header.Partner {
		partner = append(partner, Partner{
			Quotation:               data.Quotation,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}

	return &partner
}

func ConvertToPartnerUpdates(partnerUpdates *[]dpfm_api_processing_formatter.PartnerUpdates) *[]Partner {
	var partner []Partner

	for _, data := range *partnerUpdates {
		partner = append(partner, Partner{
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
		})
	}

	return &partner
}

func ConvertToAddressCreates(SDC *dpfm_api_input_reader.SDC) *[]Address {
	var address []Address

	for _, data := range SDC.Header.Address {
		address = append(address, Address{
			Quotation:   data.Quotation,
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
		})
	}

	return &address
}

func ConvertToAddressUpdates(addressUpdates *[]dpfm_api_processing_formatter.AddressUpdates) *[]Address {
	var address []Address

	for _, data := range *addressUpdates {
		address = append(address, Address{
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}

	return &address
}
