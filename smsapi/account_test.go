package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAccountDetails(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/details.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Account.Details(ctx)

	expected := &AccountDetailsResponse{
		Points:      0,
		Email:       "test",
		Name:        "test",
		PaymentType: "prepaid",
		PhoneNumber: "100200300",
		Username:    "test",
		UserType:    "native",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestCreateAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	user := &User{
		Credentials: nil,
		Active:      true,
		Description: "test user",
		Points: &UserPoints{
			FromAccount: 100,
			PerMonth:    10,
		},
	}

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/user.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestBody(t, r, new(User), user)
	})

	result, _ := client.Account.CreateUser(ctx, user)

	expected := createUserResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestUpdateAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	user := &User{
		Credentials: nil,
		Active:      false,
		Description: "updated description",
		Points:      nil,
	}

	mux.HandleFunc("/subusers/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/user.json"))

		assertRequestMethod(t, r, "PUT")
		assertRequestBody(t, r, new(User), user)
	})

	result, _ := client.Account.UpdateUser(ctx, "1", user)

	expected := createUserResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/subusers/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/user.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Account.GetUser(ctx, "1")

	expected := createUserResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestDeleteAccountUser(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/subusers/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	client.Account.DeleteUser(ctx, "1")
}

func TestGetUsersList(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/subusers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("account/users_list.json"))
		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Account.ListUsers(ctx)

	expected := &UserCollectionResponse{
		Size: 1,
		Collection: []*UserResponse{
			createUserResponse(),
		},
	}


	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func createUserResponse() *UserResponse {
	return &UserResponse{
		Id:          "1",
		Username:    "test",
		Active:      true,
		Description: "",
		Points: &UserPoints{
			FromAccount: 0,
			PerMonth:    0,
		},
	}
}
