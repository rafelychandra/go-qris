package mpm

type Payload struct {
	QRString                                            string                                               `json:"-"`
	PayloadFormatIndicator                              string                                               `json:"payloadFormatIndicator" tag:"00" presence:"M" length:"2" type:"fix"`
	PointOfInitiationMethod                             POI                                                  `json:"pointOfInitiationMethod" tag:"01" presence:"O" length:"2" type:"fix"`
	MerchantAccountInformationNonDomestic               string                                               `json:"merchantAccountInformationNonDomestic" tag:"02-25" presence:"O" length:"99" type:"upto"`
	MerchantAccountInformation                          *MerchantAccountInformation                          `json:"merchantAccountInformation" tag:"26-35" presence:"C" length:"99" type:"upto"`
	MerchantAccountInformationReservedDomesticID        string                                               `json:"merchantAccountInformationReservedDomesticID" tag:"36-39" presence:"O" length:"99" type:"upto"`
	TransferCOCIAccountInformation                      *TransferCOCIAccountInformation                      `json:"transferCOCIAccountInformation" tag:"40" presence:"C" length:"99" type:"upto"`
	MerchantAccountInformationReDomesticID              string                                               `json:"merchantAccountInformationReDomesticID" tag:"41-50" presence:"O" length:"99" type:"upto"`
	MerchantAccountInformationDomesticCentralRepository *MerchantAccountInformationDomesticCentralRepository `json:"merchantAccountInformationDomesticCentralRepository" tag:"51" presence:"C" length:"99" type:"upto"`
	MerchantCategoryCode                                string                                               `json:"merchantCategoryCode" tag:"52" presence:"M" length:"4" type:"fix"`
	TransactionCurrency                                 Currency                                             `json:"transactionCurrency" tag:"53" presence:"M" length:"3" type:"fix"`
	TransactionAmount                                   string                                               `json:"transactionAmount" tag:"54" presence:"C" length:"13" type:"upto"`
	TipIndicator                                        string                                               `json:"tipIndicator" tag:"55" presence:"O" length:"2" type:"fix"`
	TipValueOfFixed                                     string                                               `json:"tipValueOfFixed" tag:"56" presence:"C" length:"13" type:"upto"`
	TipValueOfPercentage                                string                                               `json:"tipValueOfPercentage" tag:"57" presence:"C" length:"5" type:"upto"`
	CountryCode                                         CountryCode                                          `json:"countryCode" tag:"58" presence:"M" length:"2" type:"fix"`
	MerchantName                                        string                                               `json:"merchantName" tag:"59" presence:"M" length:"25" type:"upto"`
	MerchantCity                                        string                                               `json:"merchantCity" tag:"60" presence:"M" length:"15" type:"upto"`
	PostalCode                                          string                                               `json:"postalCode" tag:"61" presence:"C" length:"10" type:"upto"`
	AdditionalDataFieldTemplate                         *AdditionalDataFieldTemplate                         `json:"additionalDataFieldTemplate" tag:"62" presence:"O" length:"99" type:"upto"`
	MerchantInformationLanguageTemplate                 *MerchantInformationLanguageTemplate                 `json:"merchantInformationLanguageTemplate" tag:"64" presence:"O" length:"99" type:"upto"`
	MerchantCountryOfOrigin                             string                                               `json:"merchantCountryOfOrigin" tag:"65" presence:"O" length:"2" type:"fix"`
	RFUForEMVCo                                         string                                               `json:"rfuForEmvCo" tag:"66-79" presence:"O" length:"99" type:"upto"`
	ReservedForDomesticID                               string                                               `json:"reservedForDomesticID" tag:"80-89" presence:"O" length:"99" type:"upto"`
	UnreservedTemplate                                  string                                               `json:"unreservedTemplate" tag:"90-99" presence:"O" length:"99" type:"upto"`
	CRC                                                 string                                               `json:"CRC" tag:"63" presence:"M" length:"4" type:"fix"`
}

type (
	MerchantAccountInformation struct {
		GlobalUniqueIdentifier string `json:"globalUniqueIdentifier" tag:"00" presence:"M" length:"32" type:"upto"`
		MPAN                   string `json:"mpan" tag:"01" presence:"M" length:"15-18" type:"between"`
		MID                    string `json:"mid" tag:"02" presence:"M" length:"15" type:"upto"`
		Kriteria               string `json:"kriteria" tag:"03" presence:"M" length:"3" type:"fix"`
		BankIdentifierCode     string `json:"bankIdentifierCode" tag:"04" presence:"O" length:"8-11" type:"between"`
	}
	TransferCOCIAccountInformation struct {
		ReverseDomain      string `json:"reverseDomain" tag:"00" presence:"M" length:"32" type:"upto"`
		CustomerPAN        string `json:"customerPAN" tag:"01" presence:"M" length:"16-19" type:"between"`
		BeneficiaryID      string `json:"beneficiaryID" tag:"02" presence:"M" length:"15" type:"upto"`
		Kriteria           string `json:"kriteria" tag:"03" presence:"C"`
		BankIdentifierCode string `json:"bankIdentifierCode" tag:"04" presence:"O" length:"8-11" type:"between"`
	}
	MerchantAccountInformationDomesticCentralRepository struct {
		GlobalUniqueIdentifier string `json:"globalUniqueIdentifier" tag:"00" presence:"M" length:"32" type:"upto"`
		NMID                   string `json:"nmid" tag:"02" presence:"M" length:"15" type:"upto"`
		PrintedVersion         string `json:"printedVersion" tag:"04" presence:"C" length:"15" type:"upto"`
	}
	AdditionalDataFieldTemplate struct {
		BillNumber                    string           `json:"billNumber" tag:"01" presence:"O" length:"25" type:"upto"`
		MobileNumber                  string           `json:"mobileNumber" tag:"02" presence:"O" length:"25" type:"upto"`
		StoreLabel                    string           `json:"storeLabel" tag:"03" presence:"O" length:"25" type:"upto"`
		LoyaltyNumber                 string           `json:"loyaltyNumber" tag:"04" presence:"O" length:"25" type:"upto"`
		ReferenceLabel                string           `json:"referenceLabel" tag:"05" presence:"O" length:"25" type:"upto"`
		CustomerLabel                 string           `json:"customerLabel" tag:"06" presence:"O" length:"25" type:"upto"`
		TerminalLabel                 string           `json:"terminalLabel" tag:"07" presence:"O" length:"25" type:"permissive"`
		PurposeOfTransaction          string           `json:"purposeOfTransaction" tag:"08" presence:"O" length:"25" type:"upto"`
		AdditionalCustomerDataRequest string           `json:"additionalCustomerDataRequest" tag:"09" presence:"O" length:"3" type:"upto"`
		MerchantTaxID                 string           `json:"merchantTaxID" tag:"10" presence:"O" length:"20" type:"upto"`
		MerchantChannel               string           `json:"merchantChannel" tag:"11" presence:"O" length:"3" type:"upto"`
		RFUForEMVCo                   string           `json:"rfuForEmvCo" tag:"12-49" presence:"O"`
		PaymentSystemSpecific         string           `json:"paymentSystemSpecific" tag:"50-98" presence:"O"`
		ProprietaryData               *ProprietaryData `json:"proprietaryData" tag:"99" presence:"O" length:"81" type:"upto"`
	}
	ProprietaryData struct {
		GlobalUniqueIdentifier string `json:"globalUniqueIdentifier" tag:"00" presence:"M" length:"32" type:"upto"`
		ProprietaryData        string `json:"proprietaryData" tag:"01" presence:"M"  length:"81" type:"upto"`
	}
	MerchantInformationLanguageTemplate struct {
		LanguagePreference string `json:"languagePreference" tag:"00" presence:"M" length:"2" type:"fix"`
		MerchantName       string `json:"merchantName" tag:"01" presence:"M" length:"25" type:"upto"`
		MerchantCity       string `json:"merchantCity" tag:"02" presence:"O" length:"15" type:"upto"`
		RFUForEMVCo        string `json:"rfuForEmvCo" tag:"03-99" presence:"O" length:"99" type:"upto"`
	}
)

func mapMerchantAccountInformation(tlvs map[Tag][]*TLV) *MerchantAccountInformation {
	if tlvs == nil {
		return nil
	}
	info := &MerchantAccountInformation{}
	for _, arr := range tlvs {
		for _, t := range arr {
			switch t.Tag {
			case "00":
				info.GlobalUniqueIdentifier = t.GetValue()
			case "01":
				info.MPAN = t.GetValue()
			case "02":
				info.MID = t.GetValue()
			case "03":
				info.Kriteria = t.GetValue()
			case "04":
				info.BankIdentifierCode = t.GetValue()
			}
		}
	}
	return info
}

func mapTransferCOCIAccountInformation(tlvs map[Tag][]*TLV) *TransferCOCIAccountInformation {
	if tlvs == nil {
		return nil
	}
	info := &TransferCOCIAccountInformation{}
	for _, arr := range tlvs {
		for _, t := range arr {
			switch t.Tag {
			case "00":
				info.ReverseDomain = t.GetValue()
			case "01":
				info.CustomerPAN = t.GetValue()
			case "02":
				info.BeneficiaryID = t.GetValue()
			case "03":
				info.Kriteria = t.GetValue()
			case "04":
				info.BankIdentifierCode = t.GetValue()
			}
		}
	}
	return info
}

func mapMerchantAccountInformationDomesticCentralRepository(tlvs map[Tag][]*TLV) *MerchantAccountInformationDomesticCentralRepository {
	if tlvs == nil {
		return nil
	}
	info := &MerchantAccountInformationDomesticCentralRepository{}
	for _, arr := range tlvs {
		for _, t := range arr {
			switch t.Tag {
			case "00":
				info.GlobalUniqueIdentifier = t.GetValue()
			case "02":
				info.NMID = t.GetValue()
			case "04":
				info.PrintedVersion = t.GetValue()
			}
		}
	}
	return info
}

func mapProprietaryData(tlvs map[Tag][]*TLV) *ProprietaryData {
	if tlvs == nil {
		return nil
	}
	info := &ProprietaryData{}
	for _, arr := range tlvs {
		for _, t := range arr {
			switch t.Tag {
			case "00":
				info.GlobalUniqueIdentifier = t.GetValue()
			case "01":
				info.ProprietaryData = t.GetValue()
			}
		}
	}
	return info
}

func mapAdditionalDataFieldTemplate(tlvs map[Tag][]*TLV) *AdditionalDataFieldTemplate {
	if tlvs == nil {
		return nil
	}
	info := &AdditionalDataFieldTemplate{}
	for _, arr := range tlvs {
		for _, t := range arr {
			switch t.Tag {
			case "01":
				info.BillNumber = t.GetValue()
			case "02":
				info.MobileNumber = t.GetValue()
			case "03":
				info.StoreLabel = t.GetValue()
			case "04":
				info.LoyaltyNumber = t.GetValue()
			case "05":
				info.ReferenceLabel = t.GetValue()
			case "06":
				info.CustomerLabel = t.GetValue()
			case "07":
				info.TerminalLabel = t.GetValue()
			case "08":
				info.PurposeOfTransaction = t.GetValue()
			case "09":
				info.AdditionalCustomerDataRequest = t.GetValue()
			case "10":
				info.MerchantTaxID = t.GetValue()
			case "11":
				info.MerchantChannel = t.GetValue()
			case "99":
				children := parseTLVs(t.Value)
				childMap := make(map[Tag][]*TLV)
				for _, c := range children {
					childMap[c.Tag] = append(childMap[c.Tag], c)
				}
				info.ProprietaryData = mapProprietaryData(childMap)
			}
		}
	}
	return info
}

func parseTLVs(data string) []*TLV {
	var result []*TLV
	i := 0
	for i < len(data) {
		if i+4 > len(data) {
			break
		}
		tag := data[i : i+2]
		length := int((data[i+2]-'0')*10 + (data[i+3] - '0'))
		if i+4+length > len(data) {
			break
		}
		value := data[i+4 : i+4+length]
		result = append(result, &TLV{Tag: Tag(tag), Value: value})
		i += 4 + length
	}
	return result
}

func mapMerchantInformationLanguageTemplate(tlvs map[Tag][]*TLV) *MerchantInformationLanguageTemplate {
	if tlvs == nil {
		return nil
	}
	info := &MerchantInformationLanguageTemplate{}
	for _, arr := range tlvs {
		for _, t := range arr {
			switch t.Tag {
			case "00":
				info.LanguagePreference = t.GetValue()
			case "01":
				info.MerchantName = t.GetValue()
			case "02":
				info.MerchantCity = t.GetValue()
			}
		}
	}
	return info
}

func (mpm *MPM) MapMPMToPayload() *Payload {
	if mpm == nil {
		return nil
	}
	return &Payload{
		QRString:                                            mpm.QRString,
		PayloadFormatIndicator:                              mpm.PayloadFormatIndicator.GetValue(),
		PointOfInitiationMethod:                             POI(mpm.PointOfInitiationMethod.GetValue()),
		MerchantAccountInformationNonDomestic:               mpm.MerchantAccountInformationNonDomestic.GetValue(),
		MerchantAccountInformation:                          mapMerchantAccountInformation(mpm.MerchantAccountInformation),
		MerchantAccountInformationReservedDomesticID:        mpm.MerchantAccountInformationReservedDomesticID.GetValue(),
		TransferCOCIAccountInformation:                      mapTransferCOCIAccountInformation(mpm.TransferCOCIAccountInformation),
		MerchantAccountInformationReDomesticID:              mpm.MerchantAccountInformationReDomesticID.GetValue(),
		MerchantAccountInformationDomesticCentralRepository: mapMerchantAccountInformationDomesticCentralRepository(mpm.MerchantAccountInformationDomesticCentralRepository),
		MerchantCategoryCode:                                mpm.MerchantCategoryCode.GetValue(),
		TransactionCurrency:                                 Currency(mpm.TransactionCurrency.GetValue()),
		TransactionAmount:                                   mpm.TransactionAmount.GetValue(),
		TipIndicator:                                        mpm.TipIndicator.GetValue(),
		TipValueOfFixed:                                     mpm.TipValueOfFixed.GetValue(),
		TipValueOfPercentage:                                mpm.TipValueOfPercentage.GetValue(),
		CountryCode:                                         CountryCode(mpm.CountryCode.GetValue()),
		MerchantName:                                        mpm.MerchantName.GetValue(),
		MerchantCity:                                        mpm.MerchantCity.GetValue(),
		PostalCode:                                          mpm.PostalCode.GetValue(),
		AdditionalDataFieldTemplate:                         mapAdditionalDataFieldTemplate(mpm.AdditionalDataFieldTemplate),
		MerchantInformationLanguageTemplate:                 mapMerchantInformationLanguageTemplate(mpm.MerchantInformationLanguageTemplate),
		MerchantCountryOfOrigin:                             mpm.MerchantCountryOfOrigin.GetValue(),
		RFUForEMVCo:                                         mpm.RFUForEMVCo.GetValue(),
		ReservedForDomesticID:                               mpm.ReservedForDomesticID.GetValue(),
		UnreservedTemplate:                                  mpm.UnreservedTemplate.GetValue(),
		CRC:                                                 mpm.CRC.GetValue(),
	}
}

func (p *Payload) mapping() *MPM {
	mpm := MPM{}
	if p.PayloadFormatIndicator != "" {
		mpm.SetPayloadFormatIndicator(p.PayloadFormatIndicator)
	}
	if p.PointOfInitiationMethod != "" {
		mpm.SetPointOfInitiationMethod(p.PointOfInitiationMethod)
	}
	if p.MerchantAccountInformationNonDomestic != "" {
		mpm.SetMerchantAccountInformationNonDomestic("02", p.MerchantAccountInformationNonDomestic)
	}
	if p.MerchantAccountInformation != nil {
		mai := p.MerchantAccountInformation
		if mai.GlobalUniqueIdentifier != "" {
			mpm.SetMerchantAccountInformation("26", "00", mai.GlobalUniqueIdentifier)
		}
		if mai.MPAN != "" {
			mpm.SetMerchantAccountInformation("26", "01", mai.MPAN)
		}
		if mai.MID != "" {
			mpm.SetMerchantAccountInformation("26", "02", mai.MID)
		}
		if mai.Kriteria != "" {
			mpm.SetMerchantAccountInformation("26", "03", mai.Kriteria)
		}
		if mai.BankIdentifierCode != "" {
			mpm.SetMerchantAccountInformation("26", "04", mai.BankIdentifierCode)
		}
	}
	if p.MerchantAccountInformationReservedDomesticID != "" {
		mpm.SetMerchantAccountInformationReservedDomesticID("36", p.MerchantAccountInformationReservedDomesticID)
	}
	if p.TransferCOCIAccountInformation != nil {
		tcoci := p.TransferCOCIAccountInformation
		if tcoci.ReverseDomain != "" {
			mpm.SetTransferCOCIAccountInformation("00", tcoci.ReverseDomain)
		}
		if tcoci.CustomerPAN != "" {
			mpm.SetTransferCOCIAccountInformation("01", tcoci.CustomerPAN)
		}
		if tcoci.BeneficiaryID != "" {
			mpm.SetTransferCOCIAccountInformation("02", tcoci.BeneficiaryID)
		}
		if tcoci.Kriteria != "" {
			mpm.SetTransferCOCIAccountInformation("03", tcoci.Kriteria)
		}
		if tcoci.BankIdentifierCode != "" {
			mpm.SetTransferCOCIAccountInformation("04", tcoci.BankIdentifierCode)
		}
	}
	if p.MerchantAccountInformationReDomesticID != "" {
		mpm.SetMerchantAccountInformationReDomesticID("41", p.MerchantAccountInformationReDomesticID)
	}
	if p.MerchantAccountInformationDomesticCentralRepository != nil {
		mai := p.MerchantAccountInformationDomesticCentralRepository
		if mai.GlobalUniqueIdentifier != "" {
			mpm.SetMerchantAccountInformationDomesticCentralRepository("00", mai.GlobalUniqueIdentifier)
		}
		if mai.NMID != "" {
			mpm.SetMerchantAccountInformationDomesticCentralRepository("02", mai.NMID)
		}
		if mai.PrintedVersion != "" {
			mpm.SetMerchantAccountInformationDomesticCentralRepository("04", mai.PrintedVersion)
		}
	}
	if p.MerchantCategoryCode != "" {
		mpm.SetMerchantCategoryCode(p.MerchantCategoryCode)
	}
	if p.TransactionCurrency != "" {
		mpm.SetTransactionCurrency(p.TransactionCurrency)
	}
	if p.TransactionAmount != "" {
		mpm.SetTransactionAmount(p.TransactionAmount)
	}
	if p.TipIndicator != "" {
		mpm.SetTipIndicator(p.TipIndicator)
	}
	if p.TipValueOfFixed != "" {
		mpm.SetTipValueOfFixed(p.TipValueOfFixed)
	}
	if p.TipValueOfPercentage != "" {
		mpm.SetTipValueOfPercentage(p.TipValueOfPercentage)
	}
	if p.CountryCode != "" {
		mpm.SetCountryCode(p.CountryCode)
	}
	if p.MerchantName != "" {
		mpm.SetMerchantName(p.MerchantName)
	}
	if p.MerchantCity != "" {
		mpm.SetMerchantCity(p.MerchantCity)
	}
	if p.PostalCode != "" {
		mpm.SetPostalCode(p.PostalCode)
	}
	if p.AdditionalDataFieldTemplate != nil {
		adt := p.AdditionalDataFieldTemplate
		if adt.BillNumber != "" {
			mpm.SetAdditionalDataFieldTemplate("01", adt.BillNumber)
		}
		if adt.MobileNumber != "" {
			mpm.SetAdditionalDataFieldTemplate("02", adt.MobileNumber)
		}
		if adt.StoreLabel != "" {
			mpm.SetAdditionalDataFieldTemplate("03", adt.StoreLabel)
		}
		if adt.LoyaltyNumber != "" {
			mpm.SetAdditionalDataFieldTemplate("04", adt.LoyaltyNumber)
		}
		if adt.ReferenceLabel != "" {
			mpm.SetAdditionalDataFieldTemplate("05", adt.ReferenceLabel)
		}
		if adt.CustomerLabel != "" {
			mpm.SetAdditionalDataFieldTemplate("06", adt.CustomerLabel)
		}
		if adt.TerminalLabel != "" {
			mpm.SetAdditionalDataFieldTemplate("07", adt.TerminalLabel)
		}
		if adt.PurposeOfTransaction != "" {
			mpm.SetAdditionalDataFieldTemplate("08", adt.PurposeOfTransaction)
		}
		if adt.AdditionalCustomerDataRequest != "" {
			mpm.SetAdditionalDataFieldTemplate("09", adt.AdditionalCustomerDataRequest)
		}
		if adt.MerchantTaxID != "" {
			mpm.SetAdditionalDataFieldTemplate("10", adt.MerchantTaxID)
		}
		if adt.MerchantChannel != "" {
			mpm.SetAdditionalDataFieldTemplate("11", adt.MerchantChannel)
		}
		if adt.RFUForEMVCo != "" {
			mpm.SetAdditionalDataFieldTemplate("12", adt.RFUForEMVCo)
		}
		if adt.PaymentSystemSpecific != "" {
			mpm.SetAdditionalDataFieldTemplate("50", adt.PaymentSystemSpecific)
		}
	}
	if p.MerchantInformationLanguageTemplate != nil {
		mi := p.MerchantInformationLanguageTemplate
		if mi.LanguagePreference != "" {
			mpm.SetMerchantInformationLanguageTemplate("00", mi.LanguagePreference)
		}
		if mi.MerchantName != "" {
			mpm.SetMerchantInformationLanguageTemplate("01", mi.MerchantName)
		}
		if mi.MerchantCity != "" {
			mpm.SetMerchantInformationLanguageTemplate("02", mi.MerchantCity)
		}
		if mi.RFUForEMVCo != "" {
			mpm.SetMerchantInformationLanguageTemplate("03", mi.RFUForEMVCo)
		}
	}
	if p.MerchantCountryOfOrigin != "" {
		mpm.SetMerchantCountryOfOrigin(p.MerchantCountryOfOrigin)
	}
	if p.RFUForEMVCo != "" {
		mpm.SetRFUForEMVCo("66", p.RFUForEMVCo)
	}
	if p.ReservedForDomesticID != "" {
		mpm.SetReservedForDomesticID("80", p.ReservedForDomesticID)
	}
	if p.UnreservedTemplate != "" {
		mpm.SetUnreservedTemplate("99", p.UnreservedTemplate)
	}

	return &mpm
}
