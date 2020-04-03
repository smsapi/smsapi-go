package smsapi

import (
	"context"
	"fmt"
)

const blacklistApiPath = "/blacklist/phone_numbers"

type BlacklistApi struct {
	client *Client
}

type BlacklistPhoneNumbersCollectionFilters struct {
	PaginationFilters
	Query string `url:"q,omitempty"`
}

type BlackListPhoneNumber struct {
	Id          string     `json:"id,omitempty"`
	PhoneNumber string     `json:"phone_number,omitempty"`
	ExpireAt    *Date      `json:"expire_at,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
}

type BlacklistPhoneNumberCollection struct {
	CollectionMeta
	Collection []*BlackListPhoneNumber `json:"collection"`
}

type BlacklistPhoneNumbersCollectionIterator struct {
	i *PageIterator
}

func (b *BlacklistPhoneNumbersCollectionIterator) Next() (*BlacklistPhoneNumberCollection, error) {
	bp := new(BlacklistPhoneNumberCollection)

	err := b.i.Next(bp)

	if err != nil {
		return nil, err
	}

	return bp, nil
}

func (blacklistApi *BlacklistApi) GetPhoneNumbers(ctx context.Context, filters *BlacklistPhoneNumbersCollectionFilters) (*BlacklistPhoneNumberCollection, error) {
	var result = new(BlacklistPhoneNumberCollection)

	uri, _ := addQueryParams(blacklistApiPath, filters)

	err := blacklistApi.client.Get(ctx, uri, result)

	return result, err
}

func (blacklistApi *BlacklistApi) GetPageIterator(ctx context.Context, filters *BlacklistPhoneNumbersCollectionFilters) *BlacklistPhoneNumbersCollectionIterator {
	ci := NewPageIterator(blacklistApi.client, ctx, blacklistApiPath, filters)
	bi := &BlacklistPhoneNumbersCollectionIterator{ci}

	return bi
}

func (blacklistApi *BlacklistApi) AddPhoneNumber(ctx context.Context, phoneNumber string, expireAt *Date) (*BlackListPhoneNumber, error) {
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
