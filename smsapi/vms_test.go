package smsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestSendVms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	vms := &Vms{
		To:  "111222333",
		Tts: "demo",
	}

	mux.HandleFunc("/vms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("vms/collection.json"))

		assertRequestQueryParam(t, r, "format", "json")
		assertRequestBody(t, r, new(Vms), vms)
	})

	sendResult, _ := client.Vms.Send(ctx, "111222333", "demo", "")

	expected := createExpectedVmsCollectionResponse()

	if !reflect.DeepEqual(sendResult, expected) {
		t.Errorf("Given: %+v Expected: %+v", sendResult, expected)
	}
}

func TestSendVmsToGroup(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	vms := &Vms{
		Group: "test",
		Tts:   "demo",
	}

	mux.HandleFunc("/vms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("vms/collection.json"))

		assertRequestQueryParam(t, r, "format", "json")
		assertRequestBody(t, r, new(Vms), vms)
	})

	sendResult, _ := client.Vms.SendToGroup(ctx, "test", "demo", "")

	expected := createExpectedVmsCollectionResponse()

	if !reflect.DeepEqual(sendResult, expected) {
		t.Errorf("Given: %+v Expected: %+v", sendResult, expected)
	}
}

func TestRemoveScheduledVms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	vmsId := "1"

	mux.HandleFunc("/vms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("vms/remove.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestJsonContains(t, r, "sch_del", vmsId)
	})

	result, _ := client.Vms.RemoveScheduled(ctx, vmsId)

	expected := &VmsRemoveResponse{
		Count: 1,
		Collection: []*struct {
			Id string `json:"id,omitempty"`
		}{
			{
				Id: vmsId,
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestScheduledVms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	dt := time.Date(2020, time.January, 01, 00, 0, 0, 0, time.Local)
	future := &Timestamp{dt}

	mux.HandleFunc("/vms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("vms/scheduled.json"))

		assertRequestQueryParam(t, r, "format", "json")

		given := new(Vms)

		json.NewDecoder(r.Body).Decode(given)

		if given.Date.String() != future.String() {
			t.Errorf("Given: %+v Expected: %+v", given.Date.String(), future)
		}
	})

	client.Vms.Schedule(ctx, "111222333", "vms-text", "", future)
}

func TestGetVms(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/vms.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("vms/collection.json"))

		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestQueryParam(t, r, "status", "1")
	})

	result, _ := client.Vms.Get(ctx, "1")

	expected := createExpectedVmsCollectionResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func createExpectedVmsCollectionResponse() *VmsCollectionResponse {
	expected := &VmsCollectionResponse{
		Count: 1,
		Collection: []*VmsResponse{
			{
				Id:              "1",
				Points:          0.1000,
				Number:          "48100200300",
				DateSent:        &Timestamp{time.Unix(1560252588, 0)},
				SubmittedNumber: "100200300",
				Status:          "QUEUE",
				Idx:             "",
				Error:           "",
			},
		},
	}

	return expected
}

func createVmsResponse() *VmsResponse {
	return &VmsResponse{
		Id:              "1",
		Points:          0.1000,
		Number:          "48100200300",
		DateSent:        &Timestamp{time.Unix(1560252588, 0)},
		SubmittedNumber: "100200300",
		Status:          "QUEUE",
		Idx:             "",
		Error:           "",
	}
}
