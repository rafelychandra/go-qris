package go_qris

import (
	"strconv"
)

type Tag string

const (
	TagPayloadFormatIndicator  Tag = "00" // Payload Format Indicator
	TagPointOfInitiationMethod Tag = "01" // Point of Initiation Method

	TagMerchantAccountInformationNonDomesticRangeStart Tag = "02" // Merchant Account Information (Non-Domestic) — Range Start (02–25)
	TagMerchantAccountInformationNonDomesticRangeEnd   Tag = "25" // Merchant Account Information (Non-Domestic) — Range End (02–25)

	TagMerchantAccountInformationRangeStart Tag = "26" // Merchant Account Information — Range Start (26–35)
	TagMerchantAccountInformationRangeEnd   Tag = "35" // Merchant Account Information — Range End (26–35)

	TagMerchantAccountInformationReservedDomesticIDStart Tag = "36" // Merchant Account Information — Reserved for Domestic Use — Range Start (36–39)
	TagMerchantAccountInformationReservedDomesticIDEnd   Tag = "39" // Merchant Account Information — Reserved for Domestic Use — Range End (36–39)

	TagTransferCOCIAccountInformation Tag = "40" // Transfer COCI Account Information

	TagMerchantAccountInformationReDomesticIDStart Tag = "41" // Merchant Account Information — Reserved for Domestic Use — Range Start (41–50)
	TagMerchantAccountInformationReDomesticIDEnd   Tag = "50" // Merchant Account Information — Reserved for Domestic Use — Range End (41–50)

	TagMerchantAccountInformationDomesticCentralRepository Tag = "51" // Merchant Account Information — Domestic Central Repository

	TagMerchantCategoryCode Tag = "52" // Merchant Category Code (MCC)
	TagTransactionCurrency  Tag = "53" // Transaction Currency (ISO 4217)
	TagTransactionAmount    Tag = "54" // Transaction Amount

	TagTipOrConvenienceIndicator       Tag = "55" // Tip or Convenience Indicator
	TagValueOfConvenienceFeeFixed      Tag = "56" // Convenience Fee — Fixed Amount
	TagValueOfConvenienceFeePercentage Tag = "57" // Convenience Fee — Percentage

	TagCountryCode  Tag = "58" // Country Code (ISO 3166-1 alpha-2)
	TagMerchantName Tag = "59" // Merchant Name
	TagMerchantCity Tag = "60" // Merchant City
	TagPostalCode   Tag = "61" // Postal Code

	TagAdditionalDataFieldTemplate Tag = "62" // Additional Data Field Template
	TagCRC                         Tag = "63" // CRC (Cyclic Redundancy Check)

	TagMerchantInformationLanguageTemplate Tag = "64" // Merchant Information — Language Template
	TagMerchantCountryOfOrigin             Tag = "65" // Merchant Country of Origin

	TagRFUForEMVCoRangeStart Tag = "66" // RFU for EMVCo — Range Start (66–79)
	TagRFUForEMVCoRangeEnd   Tag = "79" // RFU for EMVCo — Range End (66–79)

	TagReservedForDomesticIDRangeStart Tag = "80" // Reserved for Domestic Use — Range Start (80–89)
	TagReservedForDomesticIDRangeEnd   Tag = "89" // Reserved for Domestic Use — Range End (80–89)

	TagUnreservedTemplatesRangeStart Tag = "90" // Unreserved Templates — Range Start (90–99)
	TagUnreservedTemplatesRangeEnd   Tag = "99" // Unreserved Templates — Range End (90–99)
)

func (tag Tag) String() string {
	return string(tag)
}

func (tag Tag) Between(start Tag, end Tag) (bool, error) {
	idNum, err := tag.ParseInt()
	if err != nil {
		return false, err
	}

	startNum, err := start.ParseInt()
	if err != nil {
		return false, err
	}

	endNum, err := end.ParseInt()
	if err != nil {
		return false, err
	}

	return idNum >= startNum && idNum <= endNum, nil
}

func (tag Tag) ParseInt() (int64, error) {
	return strconv.ParseInt(tag.String(), 10, 64)
}
