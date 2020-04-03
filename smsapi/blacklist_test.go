package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestAddPhoneNumber(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/blacklist/phone_numbers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("blacklist/phonenumber.json"))

		assertRequestMethod(t, r, "POST")
	})

	result, _ := client.Blacklist.AddPhoneNumber(ctx, "any-phonenumber", nil)

	expected := createPhoneNumberResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestDeleteAllPhoneNumbers(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/blacklist/phone_numbers", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Blacklist.DeleteAllPhoneNumbers(ctx)

	if err != nil {
		t.Error(err)
	}
}

func TestDeletePhoneNumber(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/blacklist/phone_numbers/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Blacklist.DeletePhoneNumber(ctx, "1")

	if err != nil {
		t.Error(err)
	}
}

func TestGetPhoneNumbers(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/blacklist/phone_numbers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("blacklist/collection.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Blacklist.GetPhoneNumbers(ctx, &BlacklistPhoneNumbersCollectionFilters{})

	expected := &BlacklistPhoneNumberCollection{
		CollectionMeta: CollectionMeta{
			Size: 1,
		},
		Collection: []*BlackListPhoneNumber{
			createPhoneNumberResponse(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestBlackListPageIterator(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/blacklist/phone_numbers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("blacklist/collection.json"))

		assertRequestMethod(t, r, "GET")
	})

	iterator := client.Blacklist.GetPageIterator(ctx, nil)

	page, _ := iterator.Next()

	expectedPage := createExpectedBlacklistCollection()

	if !reflect.DeepEqual(page, expectedPage) {
		t.Errorf("Given: %+v Expected: %+v", page, expectedPage)
	}
}

func createExpectedBlacklistCollection() *BlacklistPhoneNumberCollection {
	return &BlacklistPhoneNumberCollection{
		CollectionMeta: CollectionMeta{
			Size: 1,
		},
		Collection: []*BlackListPhoneNumber{
			&BlackListPhoneNumber{
				Id:          "1",
				PhoneNumber: "654543431",
				ExpireAt:    &Date{2060, time.January, 1},
				CreatedAt:   &Timestamp{time.Date(2020, time.March, 18, 13, 0, 0, 0, time.UTC)},
			},
		},
	}
}

func createPhoneNumberResponse() *BlackListPhoneNumber {
	return &BlackListPhoneNumber{
		Id:          "1",
		PhoneNumber: "654543431",
		ExpireAt:    &Date{2060, time.January, 1},
		CreatedAt:   &Timestamp{time.Date(2020, time.March, 18, 13, 0, 0, 0, time.UTC)},
	}
}
