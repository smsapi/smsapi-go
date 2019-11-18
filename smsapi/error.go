package smsapi

import (
	"fmt"
)

type ErrorResponse struct {
	Status         int
	Code           int              `json:"error"`
	Message        string           `json:"message"`
	InvalidNumbers []*InvalidNumber `json:"invalid_numbers,omitempty"`
}

type InvalidNumber struct {
	Number          string `json:"number,omitempty"`
	SubmittedNumber string `json:"submitted_number,omitempty"`
	Message         string `json:"message,omitempty"`
}

func (error *ErrorResponse) Error() string {
	return fmt.Sprintf("Status: %d Code: %v Message: %s",
		error.Status,
		error.Code,
		error.Message,
	)
}
