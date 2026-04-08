package smsapi

import (
	"context"
	"fmt"
)

const (
	usersApiPath = "/subusers"
)

type SubusersApi struct {
	client *Client
}

type User struct {
	Credentials *UserCredentials `json:"credentials"`
	Active      *bool            `json:"active,omitempty"`
	Description string           `json:"description"`
	Points      *UserPoints      `json:"points"`
}

type UserCredentials struct {
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
	ApiPassword string `json:"api_password,omitempty"`
}

type UserPoints struct {
	FromAccount float32 `json:"from_account,omitempty"`
	PerMonth    float32 `json:"per_month,omitempty"`
}

type UserResponse struct {
	Id          string      `json:"id"`
	Username    string      `json:"username"`
	Active      bool        `json:"active"`
	Description string      `json:"description"`
	Points      *UserPoints `json:"points"`
}

type UserCollectionResponse struct {
	Size       int             `json:"size"`
	Collection []*UserResponse `json:"collection"`
}

type UserCollectionFilters struct {
	Query string `url:"q,omitempty"`
}

func (accountApi *SubusersApi) GetUser(ctx context.Context, id string) (*UserResponse, error) {
	var result = new(UserResponse)

	uri := fmt.Sprintf("%s/%s", usersApiPath, id)

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}

func (accountApi *SubusersApi) CreateUser(ctx context.Context, user *User) (*UserResponse, error) {
	var result = new(UserResponse)

	err := accountApi.client.Post(ctx, usersApiPath, result, user)

	return result, err
}

func (accountApi *SubusersApi) UpdateUser(ctx context.Context, id string, user *User) (*UserResponse, error) {
	var result = new(UserResponse)

	uri := fmt.Sprintf("%s/%s", usersApiPath, id)

	err := accountApi.client.Put(ctx, uri, result, user)

	return result, err
}

func (accountApi *SubusersApi) DeleteUser(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", usersApiPath, id)

	err := accountApi.client.Delete(ctx, uri)

	return err
}

func (accountApi *SubusersApi) ListUsers(ctx context.Context, filters *UserCollectionFilters) (*UserCollectionResponse, error) {
	var result = new(UserCollectionResponse)

	uri, _ := addQueryParams(usersApiPath, filters)

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}

type SubuserAccess struct {
	Access    string   `json:"access"`
	Senders   []string `json:"senders,omitempty"`
	Templates []string `json:"templates,omitempty"`
	Numbers   []string `json:"numbers,omitempty"`
}

type SubuserSharesResponse struct {
	Sendernames *SubuserAccess `json:"sendernames,omitempty"`
	Blacklist   *SubuserAccess `json:"blacklist,omitempty"`
	Templates   *SubuserAccess `json:"templates,omitempty"`
}

// GetShares returns the subuser shares overview (sendernames, blacklist, templates).
func (accountApi *SubusersApi) GetShares(ctx context.Context, id string) (*SubuserSharesResponse, error) {
	var result = new(SubuserSharesResponse)

	uri := fmt.Sprintf("%s/%s/shares", usersApiPath, id)

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}

// GetSendernamesAccess returns the subuser's native sendernames access configuration.
func (accountApi *SubusersApi) GetSendernamesAccess(ctx context.Context, id string) (*SubuserAccess, error) {
	var result = new(SubuserAccess)

	uri := fmt.Sprintf("%s/%s/shares/sendernames", usersApiPath, id)

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}

// UpdateSendernamesAccess updates the subuser's native sendernames access configuration.
func (accountApi *SubusersApi) UpdateSendernamesAccess(ctx context.Context, id string, access *SubuserAccess) error {
	uri := fmt.Sprintf("%s/%s/shares/sendernames", usersApiPath, id)

	return accountApi.client.Put(ctx, uri, nil, access)
}

// GetTemplatesAccess returns the subuser's native templates access configuration.
func (accountApi *SubusersApi) GetTemplatesAccess(ctx context.Context, id string) (*SubuserAccess, error) {
	var result = new(SubuserAccess)

	uri := fmt.Sprintf("%s/%s/shares/templates", usersApiPath, id)

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}

// UpdateTemplatesAccess updates the subuser's native templates access configuration.
func (accountApi *SubusersApi) UpdateTemplatesAccess(ctx context.Context, id string, access *SubuserAccess) (*SubuserAccess, error) {
	var result = new(SubuserAccess)

	uri := fmt.Sprintf("%s/%s/shares/templates", usersApiPath, id)

	err := accountApi.client.Put(ctx, uri, result, access)

	return result, err
}
