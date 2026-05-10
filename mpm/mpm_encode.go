package mpm

import (
	"fmt"
)

func NewQR(req Payload) (*MPM, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	mpm := req.mapping()
	return mpm.generateMPM(), nil
}

func (mpm *MPM) generateMPM() *MPM {
	payload := ""

	if mpm.PayloadFormatIndicator.Valid() {
		payload += mpm.PayloadFormatIndicator.TLV()
	}

	if mpm.PointOfInitiationMethod.Valid() {
		payload += mpm.PointOfInitiationMethod.TLV()
	}

	if mpm.MerchantAccountInformationNonDomestic.Valid() {
		payload += mpm.MerchantAccountInformationNonDomestic.TLV()
	}

	if mpm.MerchantAccountInformation != nil {
		payload += encodeMapTLV(mpm.MerchantAccountInformation)
	}

	if mpm.MerchantAccountInformationReservedDomesticID.Valid() {
		payload += mpm.MerchantAccountInformationReservedDomesticID.TLV()
	}

	if mpm.TransferCOCIAccountInformation != nil {
		payload += encodeMapTLV(mpm.TransferCOCIAccountInformation)
	}

	if mpm.MerchantAccountInformationReDomesticID.Valid() {
		payload += mpm.MerchantAccountInformationReDomesticID.TLV()
	}

	if mpm.MerchantAccountInformationDomesticCentralRepository != nil {
		payload += encodeMapTLV(mpm.MerchantAccountInformationDomesticCentralRepository)
	}

	if mpm.MerchantCategoryCode.Valid() {
		payload += mpm.MerchantCategoryCode.TLV()
	}

	if mpm.TransactionCurrency.Valid() {
		payload += mpm.TransactionCurrency.TLV()
	}

	if mpm.TransactionAmount.Valid() {
		payload += mpm.TransactionAmount.TLV()
	}

	if mpm.TipIndicator.Valid() {
		payload += mpm.TipIndicator.TLV()
	}

	if mpm.TipValueOfFixed.Valid() {
		payload += mpm.TipValueOfFixed.TLV()
	}

	if mpm.TipValueOfPercentage.Valid() {
		payload += mpm.TipValueOfPercentage.TLV()
	}

	if mpm.CountryCode.Valid() {
		payload += mpm.CountryCode.TLV()
	}

	if mpm.MerchantName.Valid() {
		payload += mpm.MerchantName.TLV()
	}

	if mpm.MerchantCity.Valid() {
		payload += mpm.MerchantCity.TLV()
	}

	if mpm.PostalCode.Valid() {
		payload += mpm.PostalCode.TLV()
	}

	if mpm.AdditionalDataFieldTemplate != nil {
		payload += encodeMapTLV(mpm.AdditionalDataFieldTemplate)
	}

	if mpm.MerchantCountryOfOrigin.Valid() {
		payload += mpm.MerchantCountryOfOrigin.TLV()
	}

	if mpm.MerchantInformationLanguageTemplate != nil {
		payload += encodeMapTLV(mpm.MerchantInformationLanguageTemplate)
	}

	mpm.generateCRC(payload)
	mpm.QRString = payload + mpm.CRC.TLV()
	return mpm
}

func encodeMapTLV(m map[Tag][]*TLV) string {
	result := ""
	for tag, tlvs := range m {
		if len(tlvs) == 0 {
			continue
		}
		subPayload := ""
		for _, tlv := range tlvs {
			subPayload += tlv.TLV()
		}

		length := fmt.Sprintf("%02d", len(subPayload))
		result += tag.String() + length + subPayload
	}

	return result
}
