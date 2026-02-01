package mpm

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rafelychandra/go-qris/crc16"
)

func (p *Payload) Validate() error {
	err := validateStruct(reflect.ValueOf(p).Elem())
	if err != nil {
		return err
	}

	if p.PayloadFormatIndicator != "01" {
		return fmt.Errorf("invalid payload format indicator, expected: 01")
	}

	if err := p.PointOfInitiationMethod.Validate(); err != nil {
		return err
	}

	if p.MerchantAccountInformation != nil {
		if err := p.MerchantAccountInformation.Validate(); err != nil {
			return err
		}
	}

	if p.MerchantAccountInformationDomesticCentralRepository != nil {
		if err := p.MerchantAccountInformationDomesticCentralRepository.Validate(); err != nil {
			return err
		}
	}

	if p.CountryCode == Indonesia && p.TransactionCurrency != IDR {
		return fmt.Errorf("invalid transaction currency for %s, expected: %s", Indonesia, IDR)
	}

	if p.CRC != "" {
		ok, expected := validateCRC(p.QRString)
		if !ok {
			return fmt.Errorf("invalid CRC field %s, expected CRC: %v", p.CRC, expected)
		}
	}

	return nil
}

func (mai *MerchantAccountInformation) Validate() error {
	if mai.GlobalUniqueIdentifier != "" {

	}
	return nil
}

func (maid *MerchantAccountInformationDomesticCentralRepository) Validate() error {
	if maid.GlobalUniqueIdentifier != "" {

	}
	return nil
}

func validateStruct(v reflect.Value) error {
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		fieldType := t.Field(i)

		presence := fieldType.Tag.Get("presence")
		tag := fieldType.Tag.Get("tag")
		lengthTag := fieldType.Tag.Get("length")
		lengthType := fieldType.Tag.Get("type")

		// skip CRC tag
		if tag == TagCRC.String() {
			continue
		}

		switch fieldVal.Kind() {
		case reflect.Ptr:
			if fieldVal.IsNil() {
				if presence == "M" {
					return fmt.Errorf("mandatory field %s (tag:%s) is missing", fieldType.Name, tag)
				}
				continue
			}
			elem := fieldVal.Elem()
			if elem.Kind() == reflect.Struct {
				if err := validateStruct(elem); err != nil {
					return err
				}
			} else if elem.Kind() == reflect.String {
				if err := validateLength(elem.String(), lengthTag, lengthType, fieldType.Name, tag); err != nil {
					return err
				}
			}
		case reflect.Struct:
			if err := validateStruct(fieldVal); err != nil {
				return err
			}
		case reflect.Map:
			if presence == "M" && (fieldVal.IsNil() || fieldVal.Len() == 0) {
				return fmt.Errorf("mandatory map field %s (tag:%s) is missing", fieldType.Name, tag)
			}
		default:
			if presence == "M" && fieldVal.IsZero() {
				return fmt.Errorf("mandatory field %s (tag:%s) is missing", fieldType.Name, tag)
			}
			if fieldVal.Kind() == reflect.String && fieldVal.String() != "" {
				if err := validateLength(fieldVal.String(), lengthTag, lengthType, fieldType.Name, tag); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func validateLength(val, lengthTag, lengthType, fieldName, tag string) error {
	if lengthTag == "" || lengthType == "" {
		return nil
	}

	length := len(val)

	switch lengthType {
	case "permissive":
		return nil
	case "fix":
		exp, _ := strconv.Atoi(lengthTag)
		if length != exp {
			return fmt.Errorf("field %s (tag:%s) must have exact length %d, got %d", fieldName, tag, exp, length)
		}
	case "upto":
		maxLength, _ := strconv.Atoi(lengthTag)
		if length > maxLength {
			return fmt.Errorf("field %s (tag:%s) must have length ≤ %d, got %d", fieldName, tag, maxLength, length)
		}
	case "between":
		parts := strings.Split(lengthTag, "-")
		minLength, _ := strconv.Atoi(parts[0])
		maxLength, _ := strconv.Atoi(parts[1])
		if length < minLength || length > maxLength {
			return fmt.Errorf("field %s (tag:%s) must have length between %d and %d, got %d", fieldName, tag, minLength, maxLength, length)
		}
	}

	return nil
}

func validateCRC(tlv string) (bool, string) {
	if len(tlv) < 8 { // at least must have "6304" + CRC
		return false, ""
	}

	providedCRC := tlv[len(tlv)-4:]

	data := tlv[:len(tlv)-4]

	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	crcValue := crc16.Checksum([]byte(data), table)

	expectedCRC := fmt.Sprintf("%04X", crcValue)

	return strings.EqualFold(providedCRC, expectedCRC), expectedCRC
}
