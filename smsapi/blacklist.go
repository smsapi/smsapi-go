package smsapi

import (
	"context"
	"fmt"
)

const blacklistApiPath = "/blacklist/phone_numbers"

type BlacklistApi struct {
	client *Client
}

type BlacklistPhoneNumbersListFilters struct {
	Query  string `url:"q,omitempty"`
	Offset int    `url:"offset,omitempty"`
	Limit  int    `url:"limit,omitempty"`
}

type BlackListPhoneNumber struct {
	Id          string    `json:"id,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	ExpireAt    *Timestamp `json:"expire_at,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
}

type BlacklistPhoneNumberCollection struct {
	Size       int                     `json:"size"`
	Collection []*BlackListPhoneNumber `json:"collection"`
}

func (blacklistApi *BlacklistApi) GetAllPhoneNumbers(ctx context.Context, filters *BlacklistPhoneNumbersListFilters) (*BlacklistPhoneNumberCollection, error) {
	var result = new(BlacklistPhoneNumberCollection)

	uri, _ := addQueryParams(blacklistApiPath, filters)

	err := blacklistApi.client.Get(ctx, uri, result)

	return result, err
}

func (blacklistApi *BlacklistApi) AddPhoneNumber(ctx context.Context, phoneNumber string, expireAt *Timestamp) (*BlackListPhoneNumber, error) {
	var result = new(BlackListPhoneNumber)

	blackListPhoneNumber := BlackListPhoneNumber{
		PhoneNumber: phoneNumber,
		ExpireAt:    expireAt,
	}

	err := blacklistApi.client.Post(ctx, blacklistApiPath, result, blackListPhoneNumber)

	return result, err
}

func (blacklistApi *BlacklistApi) DeleteAllPhoneNumbers(ctx context.Context) error {
	err := blacklistApi.client.Delete(ctx, blacklistApiPath)

	return err
}

func (blacklistApi *BlacklistApi) DeletePhoneNumber(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", blacklistApiPath, id)

	err := blacklistApi.client.Delete(ctx, uri)

	return err
}
