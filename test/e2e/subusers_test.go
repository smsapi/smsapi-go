package e2e

import (
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"testing"
)

var (
	userId string
)

func TestCreateUser(t *testing.T) {
	data := &smsapi.User{
		Credentials: &smsapi.UserCredentials{
			Username:    "go-smsapi",
			Password:    "Go-smsapi-1",
			ApiPassword: "Go-smsapi-1",
		},
		Active:      true,
		Description: "go-smsapi",
		Points: &smsapi.UserPoints{
			FromAccount: 10,
			PerMonth:    10,
		},
	}

	ctx, cancel := createCtx()
	defer cancel()

	user, err := client.Subusers.CreateUser(ctx, data)

	if err != nil {
		log.Fatal(err)
	}

	userId = user.Id
}

func TestGetUser(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Subusers.GetUser(ctx, userId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestGetUsersList(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Subusers.ListUsers(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	data := &smsapi.User{
		Credentials: &smsapi.UserCredentials{
			Password: "Go-smsapi-2",
		},
		Points: &smsapi.UserPoints{
			FromAccount: 5,
			PerMonth:    5,
		},
	}

	ctx, cancel := createCtx()
	defer cancel()

	user, err := client.Subusers.UpdateUser(ctx, userId, data)

	if err != nil {
		log.Fatal(err)
	}

	userId = user.Id
}

func TestRemoveUser(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Subusers.DeleteUser(ctx, userId)

	if err != nil {
		log.Fatal(err)
	}
}
