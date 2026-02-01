package mpm

func Decode(tlv string) (*MPM, error) {
	p := NewParser(tlv)
	mpm := new(MPM)
	mpm.QRString = tlv
	for p.Next() {
		tag := p.Tag()
		value := p.Value()
		switch tag {
		case TagPayloadFormatIndicator:
			mpm.SetPayloadFormatIndicator(value)
		case TagPointOfInitiationMethod:
			mpm.SetPointOfInitiationMethod(POI(value))
		case TagTransferCOCIAccountInformation:
			transferCOCIAccountInformation := NewParser(value)
			for transferCOCIAccountInformation.Next() {
				transferCOCIAccountInformationTag := transferCOCIAccountInformation.Tag()
				transferCOCIAccountInformationValue := transferCOCIAccountInformation.Value()
				mpm.SetTransferCOCIAccountInformation(transferCOCIAccountInformationTag, transferCOCIAccountInformationValue)
			}
		case TagMerchantAccountInformationDomesticCentralRepository:
			merchantAccountInformationDomesticCentralRepository := NewParser(value)
			for merchantAccountInformationDomesticCentralRepository.Next() {
				merchantAccountInformationDomesticCentralRepositoryTag := merchantAccountInformationDomesticCentralRepository.Tag()
				merchantAccountInformationDomesticCentralRepositoryValue := merchantAccountInformationDomesticCentralRepository.Value()
				mpm.SetMerchantAccountInformationDomesticCentralRepository(merchantAccountInformationDomesticCentralRepositoryTag, merchantAccountInformationDomesticCentralRepositoryValue)
			}
		case TagMerchantCategoryCode:
			mpm.SetMerchantCategoryCode(value)
		case TagTransactionCurrency:
			mpm.SetTransactionCurrency(Currency(value))
		case TagTransactionAmount:
			mpm.SetTransactionAmount(value)
		case TagTipOrConvenienceIndicator:
			mpm.SetTipIndicator(value)
		case TagValueOfConvenienceFeeFixed:
			mpm.SetTipValueOfFixed(value)
		case TagValueOfConvenienceFeePercentage:
			mpm.SetTipValueOfPercentage(value)
		case TagCountryCode:
			mpm.SetCountryCode(CountryCode(value))
		case TagMerchantName:
			mpm.SetMerchantName(value)
		case TagMerchantCity:
			mpm.SetMerchantCity(value)
		case TagPostalCode:
			mpm.SetPostalCode(value)
		case TagAdditionalDataFieldTemplate:
			additional := NewParser(value)
			for additional.Next() {
				additionalTag := additional.Tag()
				additionalValue := additional.Value()
				mpm.SetAdditionalDataFieldTemplate(additionalTag, additionalValue)
			}
		case TagCRC:
			mpm.SetCRC(value)
		case TagMerchantCountryOfOrigin:
			mpm.SetMerchantCountryOfOrigin(value)
		default:
			var (
				within bool
				err    error
			)

			within, err = tag.Between(TagMerchantAccountInformationNonDomesticRangeStart, TagMerchantAccountInformationNonDomesticRangeEnd)
			if err != nil {
				return nil, err
			}
			if within {
				mpm.SetMerchantAccountInformationNonDomestic(tag, value)
				continue
			}

			within, err = tag.Between(TagMerchantAccountInformationRangeStart, TagMerchantAccountInformationRangeEnd)
			if err != nil {
				return nil, err
			}
			if within {
				merchantAccountInformation := NewParser(value)
				for merchantAccountInformation.Next() {
					merchantAccountInformationTag := merchantAccountInformation.Tag()
					merchantAccountInformationValue := merchantAccountInformation.Value()
					mpm.SetMerchantAccountInformation(tag, merchantAccountInformationTag, merchantAccountInformationValue)
				}
				continue
			}

			within, err = tag.Between(TagMerchantAccountInformationReservedDomesticIDStart, TagMerchantAccountInformationReservedDomesticIDEnd)
			if err != nil {
				return nil, err
			}
			if within {
				mpm.SetMerchantAccountInformationReservedDomesticID(tag, value)
				continue
			}

			within, err = tag.Between(TagMerchantAccountInformationReDomesticIDStart, TagMerchantAccountInformationReDomesticIDEnd)
			if err != nil {
				return nil, err
			}
			if within {
				mpm.SetMerchantAccountInformationReDomesticID(tag, value)
				continue
			}

			within, err = tag.Between(TagRFUForEMVCoRangeStart, TagRFUForEMVCoRangeEnd)
			if err != nil {
				return nil, err
			}
			if within {
				mpm.SetRFUForEMVCo(tag, value)
				continue
			}

			within, err = tag.Between(TagReservedForDomesticIDRangeStart, TagReservedForDomesticIDRangeEnd)
			if err != nil {
				return nil, err
			}
			if within {
				mpm.SetReservedForDomesticID(tag, value)
				continue
			}

			within, err = tag.Between(TagUnreservedTemplatesRangeStart, TagUnreservedTemplatesRangeEnd)
			if err != nil {
				return nil, err
			}
			if within {
				mpm.SetUnreservedTemplate(tag, value)
				continue
			}
		}
	}
	if err := p.Err(); err != nil {
		return nil, err
	}

	return mpm, nil
}
