package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListSenders(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms/sendernames", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sendernames/list.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Sender.GetAll(ctx)

	expected := &SenderCollectionResponse{
		Size: 1,
		Collection: []*SenderResponse{
			{
				Name:      "test",
				IsDefault: false,
				Status:    "ACTIVE",
				CreatedAt: "1970-01-01T00:00:00+01:00",
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetSender(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms/sendernames/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sendernames/sender.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.Sender.Get(ctx, "test")

	expected := &SenderResponse{
		Name:      "test",
		IsDefault: false,
		Status:    "ACTIVE",
		CreatedAt: "1970-01-01T00:00:00+01:00",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestCreateSender(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	sender := &Sender{
		Name: "test",
	}

	mux.HandleFunc("/sms/sendernames", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sendernames/sender.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestBody(t, r, new(Sender), sender)
	})

	result, _ := client.Sender.Create(ctx, "test")

	expected := &SenderResponse{
		Name:      "test",
		IsDefault: false,
		Status:    "ACTIVE",
		CreatedAt: "1970-01-01T00:00:00+01:00",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestDeleteSender(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms/sendernames/test", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	err := client.Sender.Delete(ctx, "test")

	if err != nil {
		t.Fatal(err)
	}
}

func TestActivateSender(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms/sendernames/test/commands/activate", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
	})

	err := client.Sender.Activate(ctx, "test")

	if err != nil {
		t.Fatal(err)
	}
}

func TestMakeDefaultSender(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms/sendernames/test/commands/make_default", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
	})

	err := client.Sender.MakeDefault(ctx, "test")

	if err != nil {
		t.Fatal(err)
	}
}
