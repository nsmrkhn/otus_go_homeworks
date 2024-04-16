package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	runesArray := []rune(s)
	var resultString string

	for i, value := range runesArray {
		if i == 0 && unicode.IsDigit(value) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(value) && unicode.IsDigit(runesArray[i-1]) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(value) {
			multiplier, err := strconv.Atoi(string(value))
			if err != nil {
				return "", err
			}

			resultString = resultString[:len(resultString)-1]
			if multiplier != 0 {
				for j := 0; j < multiplier; j++ {
					resultString += string(runesArray[i-1])
				}
			}
		} else {
			resultString += string(value)
		}
	}
	return resultString, nil
}
