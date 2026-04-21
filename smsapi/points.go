package smsapi

import (
	"strconv"
)

// Points is an amount of SMSAPI points. It decodes from JSON values
// that may be either a number (0.3) or a numeric string ("0.3000"),
// since SMSAPI is inconsistent across endpoints.
type Points float32

func (p *Points) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		return nil
	}

	s := string(data)

	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
		if s == "" {
			return nil
		}
	}

	v, err := strconv.ParseFloat(s, 32)

	if err != nil {
		return err
	}

	*p = Points(v)

	return nil
}
