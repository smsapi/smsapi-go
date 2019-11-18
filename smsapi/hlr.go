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

func (hlrApi *HlrApi) CheckNumber(ctx context.Context, phonenumber string) (*HlrResponse, error) {
	var result = new(HlrResponse)

	v := struct {
		Number string `url:"number"`
	}{Number: phonenumber}

	uri, _ := addQueryParams("/hlr.do", v)

	err := hlrApi.client.LegacyGet(ctx, uri, result)

	return result, err
}
