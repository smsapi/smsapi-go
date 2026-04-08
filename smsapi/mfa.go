package smsapi

import (
	"context"
	"net/http"
)

type MfaApi struct {
	client *Client
}

type CreateMfaCode struct {
	PhoneNumber string `json:"phone_number"`
	From        string `json:"from,omitempty"`
	Content     string `json:"content,omitempty"`
	Fast        *bool  `json:"fast,omitempty"`
}

type MfaCode struct {
	Id          string `json:"id"`
	Code        string `json:"code"`
	PhoneNumber string `json:"phone_number"`
	From        string `json:"from,omitempty"`
}

type VerifyMfaCode struct {
	Code        string `url:"code"`
	PhoneNumber string `url:"phone_number"`
}

// CreateCode generates a new MFA code and sends it to the given phone number.
func (api *MfaApi) CreateCode(ctx context.Context, req *CreateMfaCode) (*MfaCode, error) {
	result := new(MfaCode)
	err := api.client.Post(ctx, "/mfa/codes", result, req)
	return result, err
}

// VerifyCode verifies the MFA code for the given phone number.
func (api *MfaApi) VerifyCode(ctx context.Context, phoneNumber, code string) error {
	body := &VerifyMfaCode{Code: code, PhoneNumber: phoneNumber}
	return api.client.Urlencoded(ctx, http.MethodPost, "/mfa/codes/verifications", nil, body)
}
