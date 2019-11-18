package smsapi

import (
	"context"
	"fmt"
)

const (
	profileApiPath = "/profile"
	usersApiPath   = "/subusers"
)

type AccountApi struct {
	client *Client
}

type AccountDetailsResponse struct {
	Name        string  `json:"name,omitempty"`
	Email       string  `json:"email,omitempty"`
	Username    string  `json:"username,omitempty"`
	PhoneNumber string  `json:"phone_number,omitempty"`
	PaymentType string  `json:"payment_type,omitempty"`
	UserType    string  `json:"user_type,omitempty"`
	Points      float32 `json:"points,omitempty"`
}

func (accountApi *AccountApi) Details(ctx context.Context) (*AccountDetailsResponse, error) {
	var result = new(AccountDetailsResponse)

	err := accountApi.client.Get(ctx, profileApiPath, result)

	return result, err
}

type User struct {
	Credentials *UserCredentials `json:"credentials"`
	Active      bool             `json:"Active"`
	Description string           `json:"Description"`
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

func (accountApi *AccountApi) GetUser(ctx context.Context, id string) (*UserResponse, error) {
	var result = new(UserResponse)

	uri := fmt.Sprintf("%s/%s", usersApiPath, id)

	err := accountApi.client.Get(ctx, uri, result)

	return result, err
}

func (accountApi *AccountApi) CreateUser(ctx context.Context, user *User) (*UserResponse, error) {
	var result = new(UserResponse)

	err := accountApi.client.Post(ctx, usersApiPath, result, user)

	return result, err
}

func (accountApi *AccountApi) UpdateUser(ctx context.Context, id string, user *User) (*UserResponse, error) {
	var result = new(UserResponse)

	uri := fmt.Sprintf("/subusers/%s", id)

	err := accountApi.client.Put(ctx, uri, result, user)

	return result, err
}

func (accountApi *AccountApi) DeleteUser(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", usersApiPath, id)

	err := accountApi.client.Delete(ctx, uri)

	return err
}

func (accountApi *AccountApi) ListUsers(ctx context.Context) (*UserCollectionResponse, error) {
	var result = new(UserCollectionResponse)

	err := accountApi.client.Get(ctx, usersApiPath, result)

	return result, err
}
