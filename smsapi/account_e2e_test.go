// +build e2e

package smsapi

import (
	"log"
	"testing"
)

var (
	userId string
)

func TestGetUserDetails(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Account.Details(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	data := &User{
		Credentials: &UserCredentials{
			Username:    "go-smsapi",
			Password:    "Go-smsapi-1",
			ApiPassword: "Go-smsapi-1",
		},
		Active:      true,
		Description: "go-smsapi",
		Points: &UserPoints{
			FromAccount: 10,
			PerMonth:    10,
		},
	}

	ctx, cancel := createCtx()
	defer cancel()

	user, err := client.Account.CreateUser(ctx, data)

	if err != nil {
		log.Fatal(err)
	}

	userId = user.Id
}

func TestGetUser(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Account.GetUser(ctx, userId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestGetUsersList(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Account.ListUsers(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	data := &User{
		Credentials: &UserCredentials{
			Password: "Go-smsapi-2",
		},
		Points: &UserPoints{
			FromAccount: 5,
			PerMonth:    5,
		},
	}

	ctx, cancel := createCtx()
	defer cancel()

	user, err := client.Account.UpdateUser(ctx, userId, data)

	if err != nil {
		log.Fatal(err)
	}

	userId = user.Id
}

func TestRemoveUser(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Account.DeleteUser(ctx, userId)

	if err != nil {
		log.Fatal(err)
	}
}
