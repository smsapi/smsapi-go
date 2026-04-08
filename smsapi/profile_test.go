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

func TestGetProfilePrices(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/profile/prices", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "type", "eco")
		fmt.Fprint(w, `{"size":1,"collection":[{"price":{"value":"0.08","currency":"PLN"},"country":"PL","network":"Plus","changed_at":"2024-01-01"}]}`)
	})

	result, err := client.Profile.Prices(ctx, "eco")
	if err != nil {
		t.Fatal(err)
	}

	if result.Size != 1 || result.Collection[0].Country != "PL" || result.Collection[0].Price.Value != "0.08" {
		t.Errorf("Unexpected: %+v", result)
	}
}
