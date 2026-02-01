package mpm

import (
	"fmt"
	"strconv"
)

func LuhnChecksum(value string) (string, error) {
	sum := 0
	alt := true

	for i := len(value) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(value[i]))
		if err != nil {
			return "", err
		}
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}

	checkDigit := (10 - (sum % 10)) % 10
	return fmt.Sprintf("%s%d", value, checkDigit), nil
}

func LuhnValidate(number string) bool {
	sum := 0
	alt := false

	for i := len(number) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}

	return sum%10 == 0
}
