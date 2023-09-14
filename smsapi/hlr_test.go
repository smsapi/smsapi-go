package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCheckNumberByHlr(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/hlr.do", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("hlr/check_number.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "format", "json")
		assertRequestJsonContains(t, r, "number", "100200300")
	})

	result, _ := client.Hlr.CheckNumber(ctx, "100200300")

	expected := &HlrResponse{
		Id:     "1",
		Status: "OK",
		Number: "48100200300",
		Price:  0.1,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}
