package smsapi

import "context"

type HlrResponse struct {
	Status string  `json:"status,omitempty"`
	Number string  `json:"number,omitempty"`
	Id     string  `json:"id,omitempty"`
	Price  float32 `json:"price,omitempty"`
}

type HlrApi struct {
	client *Client
}

type Hlr struct {
	PhoneNumber string `json:"number"`
}

func (hlrApi *HlrApi) CheckNumber(ctx context.Context, phonenumber string) (*HlrResponse, error) {
	var result = new(HlrResponse)

	payload := Hlr{
		PhoneNumber: phonenumber,
	}

	err := hlrApi.client.LegacyPost(ctx, "/hlr.do", result, payload)

	return result, err
}
