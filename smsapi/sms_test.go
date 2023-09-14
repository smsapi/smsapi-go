package smsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestSendSms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	sms := &Sms{
		To:      "48100200300",
		Message: "test",
		From:    "",
	}

	mux.HandleFunc("/sms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sms/collection.json"))

		assertRequestBody(t, r, new(Sms), sms)
		assertRequestQueryParam(t, r, "format", "json")
	})

	sendResult, _ := client.Sms.Send(ctx, "48100200300", "test", "")

	expected := createSmsResultCollection()

	if !reflect.DeepEqual(sendResult, expected) {
		t.Errorf("Given: %+v Expected: %+v", sendResult, expected)
	}
}

func TestScheduledSms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	dt := time.Date(2020, time.January, 01, 00, 0, 0, 0, time.Local)
	future := &Timestamp{dt}

	mux.HandleFunc("/sms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sms/scheduled.json"))

		assertRequestQueryParam(t, r, "format", "json")

		given := new(Sms)

		json.NewDecoder(r.Body).Decode(given)

		if given.Date.String() != future.String() {
			t.Errorf("Given: %+v Expected: %+v", given.Date.String(), future)
		}
	})

	client.Sms.Schedule(ctx, "48100200300", "test", "", future)
}

func TestSendSmsError(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sms/invalid_recipient.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "format", "json")
	})

	_, err := client.Sms.Send(ctx, "", "", "")

	expected := &ErrorResponse{
		Message: "No correct phone numbers",
		Code:    13,
		Status:  200,
		InvalidNumbers: []*InvalidNumber{
			{
				Number:          "48100200300",
				SubmittedNumber: "100200300",
				Message:         "Invalid phone number",
			},
		},
	}

	if !reflect.DeepEqual(err, expected) {
		t.Errorf("Given: %s Expected: %s", err, expected)
	}
}

func TestRemoveScheduledSms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	smsId := "1"

	mux.HandleFunc("/sms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sms/remove.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestJsonContains(t, r, "sch_del", smsId)
	})

	result, _ := client.Sms.RemoveScheduled(ctx, smsId)

	expected := &SmsRemoveResult{
		Count: 1,
		Collection: []*struct {
			Id string `json:"id,omitempty"`
		}{
			{
				Id: smsId,
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetSms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/sms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sms/sms.json"))

		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestQueryParam(t, r, "status", "1")
	})

	result, _ := client.Sms.Get(ctx, "1")

	expected := createSmsResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestSendSmsToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	sms := &Sms{
		Group:   "some-group",
		Message: "test",
		From:    "",
	}

	mux.HandleFunc("/sms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("sms/collection.json"))

		assertRequestBody(t, r, new(Sms), sms)
		assertRequestQueryParam(t, r, "format", "json")
	})

	sendResult, _ := client.Sms.SendToGroup(ctx, "some-group", "test", "")

	expected := createSmsResultCollection()

	if !reflect.DeepEqual(sendResult, expected) {
		t.Errorf("Given: %+v Expected: %+v", sendResult, expected)
	}
}

func createSmsResultCollection() *SmsResultCollection {
	return &SmsResultCollection{
		Count: 1,
		Collection: []*SmsResponse{
			createSmsResponse(),
		},
	}
}

func createSmsResponse() *SmsResponse {
	return &SmsResponse{
		Id:              "1",
		Points:          0.1,
		Number:          "48100200300",
		DateSent:        &Timestamp{time.Unix(1560252588, 0)},
		SubmittedNumber: "100200300",
		Status:          "QUEUE",
		Message:         "test",
		Length:          1,
		Parts:           1,
	}
}
