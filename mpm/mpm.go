package mpm

type (
	MPM struct {
		QRString                                            string         `json:"-,omitempty"`
		PayloadFormatIndicator                              *TLV           `json:"payloadFormatIndicator" tag:"00" presence:"M"`
		PointOfInitiationMethod                             *TLV           `json:"pointOfInitiationMethod" tag:"01" presence:"O"`
		MerchantAccountInformationNonDomestic               *TLV           `json:"merchantAccountInformationNonDomestic" tag:"02-25" presence:"O"`
		MerchantAccountInformation                          map[Tag][]*TLV `json:"merchantAccountInformation" tag:"26-35" presence:"C"`
		MerchantAccountInformationReservedDomesticID        *TLV           `json:"merchantAccountInformationReservedDomesticID" tag:"36-39" presence:"O"`
		TransferCOCIAccountInformation                      map[Tag][]*TLV `json:"transferCOCIAccountInformation" tag:"40" presence:"C"`
		MerchantAccountInformationReDomesticID              *TLV           `json:"merchantAccountInformationReDomesticID" tag:"41-50" presence:"O"`
		MerchantAccountInformationDomesticCentralRepository map[Tag][]*TLV `json:"merchantAccountInformationDomesticCentralRepository" tag:"51" presence:"C"`
		MerchantCategoryCode                                *TLV           `json:"merchantCategoryCode" tag:"52" presence:"M"`
		TransactionCurrency                                 *TLV           `json:"transactionCurrency" tag:"53" presence:"M"`
		TransactionAmount                                   *TLV           `json:"transactionAmount" tag:"54" presence:"C"`
		TipIndicator                                        *TLV           `json:"tipIndicator" tag:"55" presence:"O"`
		TipValueOfFixed                                     *TLV           `json:"tipValueOfFixed" tag:"56" presence:"C"`
		TipValueOfPercentage                                *TLV           `json:"tipValueOfPercentage" tag:"57" presence:"C"`
		CountryCode                                         *TLV           `json:"countryCode" tag:"58" presence:"M"`
		MerchantName                                        *TLV           `json:"merchantName" tag:"59" presence:"M"`
		MerchantCity                                        *TLV           `json:"merchantCity" tag:"60" presence:"M"`
		PostalCode                                          *TLV           `json:"postalCode" tag:"61" presence:"C"`
		AdditionalDataFieldTemplate                         map[Tag][]*TLV `json:"additionalDataFieldTemplate" tag:"62" presence:"O"`
		MerchantInformationLanguageTemplate                 map[Tag][]*TLV `json:"merchantInformationLanguageTemplate" tag:"64" presence:"O"`
		MerchantCountryOfOrigin                             *TLV           `json:"merchantCountryOfOrigin" tag:"65" presence:"O"`
		RFUForEMVCo                                         *TLV           `json:"rfuForEmvCo" tag:"66-79" presence:"O"`
		ReservedForDomesticID                               *TLV           `json:"reservedForDomesticID" tag:"80-89" presence:"O"`
		UnreservedTemplate                                  *TLV           `json:"unreservedTemplate" tag:"90-99" presence:"O"`
		CRC                                                 *TLV           `json:"CRC" tag:"63" presence:"M"`
		ProductIndicator                                    *TLV           `json:"productIndicator" tag:"PI" presence:"O"`
		CustomerData                                        *TLV           `json:"customerData" tag:"CD" presence:"O"`
		MerchantCriteria                                    *TLV           `json:"merchantCriteria" tag:"MC" presence:"O"`
	}
)
