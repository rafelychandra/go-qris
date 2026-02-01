package mpm

import (
	"fmt"
	"unicode/utf8"
)

func l(v string) string {
	return fmt.Sprintf("%02d", utf8.RuneCountInString(v))
}

func (mpm *MPM) SetPayloadFormatIndicator(v string) {
	mpm.PayloadFormatIndicator = &TLV{
		Tag:    TagPayloadFormatIndicator,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetPointOfInitiationMethod(v POI) {
	mpm.PointOfInitiationMethod = &TLV{
		Tag:    TagPointOfInitiationMethod,
		Length: l(string(v)),
		Value:  string(v),
	}
}

func (mpm *MPM) SetMerchantAccountInformationNonDomestic(tag Tag, v string) {
	mpm.MerchantAccountInformationNonDomestic = &TLV{
		Tag:    tag,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetMerchantAccountInformation(parentTag Tag, nestedTag Tag, v string) {
	if mpm.MerchantAccountInformation == nil {
		mpm.MerchantAccountInformation = make(map[Tag][]*TLV)
	}

	tlv := &TLV{
		Tag:    nestedTag,
		Length: l(v),
		Value:  v,
	}

	mpm.MerchantAccountInformation[parentTag] = append(mpm.MerchantAccountInformation[parentTag], tlv)
}

func (mpm *MPM) SetMerchantAccountInformationReservedDomesticID(tag Tag, v string) {
	mpm.MerchantAccountInformationReservedDomesticID = &TLV{
		Tag:    tag,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetTransferCOCIAccountInformation(nestedTag Tag, v string) {
	if mpm.TransferCOCIAccountInformation == nil {
		mpm.TransferCOCIAccountInformation = make(map[Tag][]*TLV)
	}

	tlv := &TLV{
		Tag:    nestedTag,
		Length: l(v),
		Value:  v,
	}

	mpm.TransferCOCIAccountInformation[TagTransferCOCIAccountInformation] = append(mpm.TransferCOCIAccountInformation[TagTransferCOCIAccountInformation], tlv)
}

func (mpm *MPM) SetMerchantAccountInformationReDomesticID(tag Tag, v string) {
	mpm.MerchantAccountInformationReDomesticID = &TLV{
		Tag:    tag,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetMerchantAccountInformationDomesticCentralRepository(nestedTag Tag, v string) {
	if mpm.MerchantAccountInformationDomesticCentralRepository == nil {
		mpm.MerchantAccountInformationDomesticCentralRepository = make(map[Tag][]*TLV)
	}

	tlv := &TLV{
		Tag:    nestedTag,
		Length: l(v),
		Value:  v,
	}

	mpm.MerchantAccountInformationDomesticCentralRepository[TagMerchantAccountInformationDomesticCentralRepository] = append(mpm.MerchantAccountInformationDomesticCentralRepository[TagMerchantAccountInformationDomesticCentralRepository], tlv)
}

func (mpm *MPM) SetMerchantCategoryCode(v string) {
	mpm.MerchantCategoryCode = &TLV{
		Tag:    TagMerchantCategoryCode,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetTransactionCurrency(v Currency) {
	mpm.TransactionCurrency = &TLV{
		Tag:    TagTransactionCurrency,
		Length: l(string(v)),
		Value:  string(v),
	}
}

func (mpm *MPM) SetTransactionAmount(v string) {
	mpm.TransactionAmount = &TLV{
		Tag:    TagTransactionAmount,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetTipIndicator(v string) {
	mpm.TipIndicator = &TLV{
		Tag:    TagTipOrConvenienceIndicator,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetTipValueOfFixed(v string) {
	mpm.TipValueOfFixed = &TLV{
		Tag:    TagValueOfConvenienceFeeFixed,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetTipValueOfPercentage(v string) {
	mpm.TipValueOfPercentage = &TLV{
		Tag:    TagValueOfConvenienceFeePercentage,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetCountryCode(v CountryCode) {
	mpm.CountryCode = &TLV{
		Tag:    TagCountryCode,
		Length: l(string(v)),
		Value:  string(v),
	}
}

func (mpm *MPM) SetMerchantName(v string) {
	mpm.MerchantName = &TLV{
		Tag:    TagMerchantName,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetMerchantCity(v string) {
	mpm.MerchantCity = &TLV{
		Tag:    TagMerchantCity,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetPostalCode(v string) {
	mpm.PostalCode = &TLV{
		Tag:    TagPostalCode,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetAdditionalDataFieldTemplate(nestedTag Tag, v string) {
	if mpm.AdditionalDataFieldTemplate == nil {
		mpm.AdditionalDataFieldTemplate = make(map[Tag][]*TLV)
	}

	tlv := &TLV{
		Tag:    nestedTag,
		Length: l(v),
		Value:  v,
	}

	mpm.AdditionalDataFieldTemplate[TagAdditionalDataFieldTemplate] = append(mpm.AdditionalDataFieldTemplate[TagAdditionalDataFieldTemplate], tlv)
}

func (mpm *MPM) SetMerchantInformationLanguageTemplate(nestedTag Tag, v string) {
	if mpm.MerchantInformationLanguageTemplate == nil {
		mpm.MerchantInformationLanguageTemplate = make(map[Tag][]*TLV)
	}

	tlv := &TLV{
		Tag:    nestedTag,
		Length: l(v),
		Value:  v,
	}

	mpm.MerchantInformationLanguageTemplate[TagMerchantInformationLanguageTemplate] = append(mpm.MerchantInformationLanguageTemplate[TagMerchantInformationLanguageTemplate], tlv)
}

func (mpm *MPM) SetMerchantCountryOfOrigin(v string) {
	mpm.MerchantCountryOfOrigin = &TLV{
		Tag:    TagMerchantCountryOfOrigin,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetRFUForEMVCo(tag Tag, v string) {
	mpm.RFUForEMVCo = &TLV{
		Tag:    tag,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetReservedForDomesticID(tag Tag, v string) {
	mpm.ReservedForDomesticID = &TLV{
		Tag:    tag,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) SetUnreservedTemplate(tag Tag, v string) {
	mpm.UnreservedTemplate = &TLV{
		Tag:    tag,
		Length: l(v),
		Value:  v,
	}
}

func (mpm *MPM) generateCRC(v string) {
	crc := generateCRC(v)
	mpm.CRC = &TLV{
		Tag:    TagCRC,
		Length: l(crc),
		Value:  crc,
	}
}

func (mpm *MPM) SetCRC(v string) {
	mpm.CRC = &TLV{
		Tag:    TagCRC,
		Length: l(v),
		Value:  v,
	}
}
