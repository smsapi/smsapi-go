package smsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestSendMms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mms := &Mms{
		To:      "111222333",
		Subject: "test-subject",
		Message: &SMIL{},
	}

	mux.HandleFunc("/mms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("mms/collection.json"))

		assertRequestQueryParam(t, r, "format", "json")
		assertRequestBody(t, r, new(Mms), mms)
	})

	sendResult, _ := client.Mms.Send(ctx, "111222333", "test-subject", "")

	expected := createMmsCollectionResponse()

	if !reflect.DeepEqual(sendResult, expected) {
		t.Errorf("Given: %+v Expected: %+v", sendResult, expected)
	}
}

func TestScheduledMms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	dt := time.Date(2020, time.January, 01, 00, 0, 0, 0, time.Local)
	future := &Timestamp{dt}

	mux.HandleFunc("/mms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("mms/scheduled.json"))

		assertRequestQueryParam(t, r, "format", "json")

		given := new(Mms)

		json.NewDecoder(r.Body).Decode(given)

		if given.Date.String() != future.String() {
			t.Errorf("Given: %+v Expected: %+v", given.Date.String(), future)
		}
	})

	client.Mms.Schedule(ctx, "111222333", "test-subject", "url-to-image", future)
}

func TestGetMms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/mms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("mms/mms.json"))

		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestQueryParam(t, r, "status", "1")
	})

	result, _ := client.Mms.Get(ctx, "1")

	expected := createMmsResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestSendMmsToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mms := &Mms{
		Group:   "test",
		Subject: "test-subject",
		Message: &SMIL{},
	}

	mux.HandleFunc("/mms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("mms/collection.json"))

		assertRequestQueryParam(t, r, "format", "json")
		assertRequestBody(t, r, new(Mms), mms)
	})

	sendResult, _ := client.Mms.SendToGroup(ctx, "test", "test-subject", "url-to-image")

	expected := createMmsCollectionResponse()

	if !reflect.DeepEqual(sendResult, expected) {
		t.Errorf("Given: %+v Expected: %+v", sendResult, expected)
	}
}

func TestRemoveScheduledMms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/mms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("mms/remove_scheduled.json"))

		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestQueryParam(t, r, "sch_del", "1")
	})

	result, _ := client.Mms.RemoveScheduled(ctx, "1")

	expected := &MmsRemoveResponse{
		Count: 1,
		Collection: []*struct {
			Id string `json:"id,omitempty"`
		}{
			{
				Id: "1",
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestMarshalMms(t *testing.T) {
	smil := NewSMIL()
	smil.AddImage("https://smsapi.pl/logo.png")

	mms := &Mms{
		To:      "111222333",
		Message: smil,
	}

	_, err := json.Marshal(mms)

	if err != nil {
		fmt.Println(err)
	}
}

func createMmsCollectionResponse() *MmsCollectionResponse {
	expected := &MmsCollectionResponse{
		Count: 1,
		Collection: []*MmsResponse{
			createMmsResponse(),
		},
	}

	return expected
}

func createMmsResponse() *MmsResponse {
	return &MmsResponse{
		Id:              "1",
		Points:          "0.1000",
		Number:          "48100200300",
		DateSent:        &Timestamp{time.Unix(1560252588, 0)},
		SubmittedNumber: "100200300",
		Status:          "QUEUE",
		Idx:             "",
		Error:           "",
	}
}
