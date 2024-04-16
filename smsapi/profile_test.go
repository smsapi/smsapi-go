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

	result, _ := client.Profile.Details(ctx)

	expected := &ProfileDetailsResponse{
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
