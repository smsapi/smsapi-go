package smsapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMfaCreateCode(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mfa/codes", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
		assertRequestJsonContains(t, r, "phone_number", "48500500500")
		fmt.Fprint(w, `{"id":"abc","code":"123456","phone_number":"48500500500","from":"Info"}`)
	})

	result, err := client.Mfa.CreateCode(ctx, &CreateMfaCode{
		PhoneNumber: "48500500500",
		From:        "Info",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Code != "123456" || result.Id != "abc" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestMfaVerifyCode(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mfa/codes/verifications", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
		assertRequestUrlencoded(t, r, &VerifyMfaCode{
			Code:        "123456",
			PhoneNumber: "48500500500",
		})
		w.WriteHeader(http.StatusNoContent)
	})

	if err := client.Mfa.VerifyCode(ctx, "48500500500", "123456"); err != nil {
		t.Fatal(err)
	}
}
