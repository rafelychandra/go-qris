package go_qris

import (
	"fmt"
	"strconv"
	"strings"

	"go-qris/crc16"
)

type (
	TLV struct {
		Tag    Tag    `json:"tag"`
		Length string `json:"length"`
		Value  string `json:"value"`
	}
)

func (t *TLV) TLV() string {
	if t == nil {
		return ""
	}
	length := fmt.Sprintf("%02d", len(t.Value))
	return t.Tag.String() + length + t.Value
}

func (t *TLV) Valid() bool {
	return t != nil && t.Tag != "" && t.Value != ""
}

func (t *TLV) WithLuhn() (string, error) {
	return LuhnChecksum(t.Value)
}

func generateCRC(value string) string {
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)
	crcValue := crc16.Checksum([]byte(value+TagCRC.String()+"04"), table)
	crcValueString := strconv.FormatUint(uint64(crcValue), 16)
	s := "0000" + strings.ToUpper(crcValueString)
	return fmt.Sprintf("%s", s[len(s)-4:])
}
