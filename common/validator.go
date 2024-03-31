package common

import (
	"net/mail"
	"regexp"
	"time"
)

func IsEmptyString(s string) bool {
	return s == ""
}

func IsEmptyPointerString(s *string) bool {
	return s == nil || *s == ""
}

func IsPhone(s string) bool {
	if len(s) != 10 && len(s) != 11 {
		return false
	}

	pattern := `^[0-9]+$`

	regex := regexp.MustCompile(pattern)

	return regex.MatchString(s)
}

func IsEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

func IsImage(s *string, defaultValue string) bool {
	if s == nil {
		s = &defaultValue
		return true
	}
	if *s == "" {
		*s = defaultValue
		return true
	}

	return true
}

func IsNotNilId(id *string) bool {
	if id == nil || len(*id) == 0 || len(*id) > MaxLengthIdCanGenerate {
		return false
	}
	return true
}

func IsFeatureCode(id *string) bool {
	if id == nil || len(*id) == 0 || len(*id) > MaxLengthOfFeatureCode {
		return false
	}
	return true
}

func IsId(id *string) bool {
	if id == nil || len(*id) == 0 {
		return true
	}
	if len(*id) > MaxLengthIdCanGenerate {
		return false
	}
	return true
}

func IsDateString(s string) bool {
	pattern := `^\d{2}/\d{2}/\d{4}$`
	matched, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}

	if !matched {
		return false
	}

	_, err = time.Parse("02/01/2006", s)
	if err != nil {
		return false
	}

	return true
}

func IsNotNegativeNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v >= 0
	case int8:
		return v >= 0
	case int16:
		return v >= 0
	case int64:
		return v >= 0
	case float32:
		return v >= 0
	case float64:
		return v >= 0
	default:
		return false
	}
}

func IsNegativeNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v < 0
	case int8:
		return v < 0
	case int16:
		return v < 0
	case int64:
		return v < 0
	case float32:
		return v < 0
	case float64:
		return v < 0
	default:
		return false
	}
}

func IsNotPositiveNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v <= 0
	case int8:
		return v <= 0
	case int16:
		return v <= 0
	case int64:
		return v <= 0
	case float32:
		return v <= 0
	case float64:
		return v <= 0
	default:
		return false
	}
}

func IsPassword(pass *string) bool {
	return pass != nil && len(*pass) >= 6
}

func IsPositiveNumber(number interface{}) bool {
	switch v := number.(type) {
	case int:
		return v > 0
	case int8:
		return v > 0
	case int16:
		return v > 0
	case int64:
		return v > 0
	case float32:
		return v > 0
	case float64:
		return v > 0
	default:
		return false
	}
}
