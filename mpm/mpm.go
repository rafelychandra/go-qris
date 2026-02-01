package mpm

type (
	MPM struct {
		QRString                                            string         `json:"-,omitempty"`
		PayloadFormatIndicator                              *TLV           `json:"payloadFormatIndicator,omitempty" tag:"00" presence:"M"`
		PointOfInitiationMethod                             *TLV           `json:"pointOfInitiationMethod,omitempty" tag:"01" presence:"O"`
		MerchantAccountInformationNonDomestic               *TLV           `json:"merchantAccountInformationNonDomestic,omitempty" tag:"02-25" presence:"O"`
		MerchantAccountInformation                          map[Tag][]*TLV `json:"merchantAccountInformation,omitempty" tag:"26-35" presence:"C"`
		MerchantAccountInformationReservedDomesticID        *TLV           `json:"merchantAccountInformationReservedDomesticID,omitempty" tag:"36-39" presence:"O"`
		TransferCOCIAccountInformation                      map[Tag][]*TLV `json:"transferCOCIAccountInformation,omitempty" tag:"40" presence:"C"`
		MerchantAccountInformationReDomesticID              *TLV           `json:"merchantAccountInformationReDomesticID,omitempty" tag:"41-50" presence:"O"`
		MerchantAccountInformationDomesticCentralRepository map[Tag][]*TLV `json:"merchantAccountInformationDomesticCentralRepository,omitempty" tag:"51" presence:"C"`
		MerchantCategoryCode                                *TLV           `json:"merchantCategoryCode,omitempty" tag:"52" presence:"M"`
		TransactionCurrency                                 *TLV           `json:"transactionCurrency,omitempty" tag:"53" presence:"M"`
		TransactionAmount                                   *TLV           `json:"transactionAmount,omitempty" tag:"54" presence:"C"`
		TipIndicator                                        *TLV           `json:"tipIndicator,omitempty" tag:"55" presence:"O"`
		TipValueOfFixed                                     *TLV           `json:"tipValueOfFixed,omitempty" tag:"56" presence:"C"`
		TipValueOfPercentage                                *TLV           `json:"tipValueOfPercentage,omitempty" tag:"57" presence:"C"`
		CountryCode                                         *TLV           `json:"countryCode,omitempty" tag:"58" presence:"M"`
		MerchantName                                        *TLV           `json:"merchantName,omitempty" tag:"59" presence:"M"`
		MerchantCity                                        *TLV           `json:"merchantCity,omitempty" tag:"60" presence:"M"`
		PostalCode                                          *TLV           `json:"postalCode,omitempty" tag:"61" presence:"C"`
		AdditionalDataFieldTemplate                         map[Tag][]*TLV `json:"additionalDataFieldTemplate,omitempty" tag:"62" presence:"O"`
		MerchantInformationLanguageTemplate                 map[Tag][]*TLV `json:"merchantInformationLanguageTemplate,omitempty" tag:"64" presence:"O"`
		MerchantCountryOfOrigin                             *TLV           `json:"merchantCountryOfOrigin,omitempty" tag:"65" presence:"O"`
		RFUForEMVCo                                         *TLV           `json:"rfuForEmvCo,omitempty" tag:"66-79" presence:"O"`
		ReservedForDomesticID                               *TLV           `json:"reservedForDomesticID,omitempty" tag:"80-89" presence:"O"`
		UnreservedTemplate                                  *TLV           `json:"unreservedTemplate,omitempty" tag:"90-99" presence:"O"`
		CRC                                                 *TLV           `json:"CRC,omitempty" tag:"63" presence:"M"`
	}
)
