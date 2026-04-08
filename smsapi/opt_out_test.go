package smsapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOptOutList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/opt_outs", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"size":1,"collection":[{"id":"1","phoneNumber":48500500500,"date":"2024-01-01T00:00:00+00:00"}]}`)
	})

	result, err := client.OptOut.List(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Size != 1 || result.Collection[0].Id != "1" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestOptOutDelete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/opt_outs/abc", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	if err := client.OptOut.Delete(ctx, "abc"); err != nil {
		t.Fatal(err)
	}
}

func TestOptOutGetSettings(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/opt_outs/settings", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"brand":"Acme"}`)
	})

	result, err := client.OptOut.GetSettings(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result.Brand != "Acme" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestOptOutUpdateSettings(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/opt_outs/settings", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
		assertRequestJsonContains(t, r, "brand", "Acme")
		fmt.Fprint(w, `{"brand":"Acme"}`)
	})

	result, err := client.OptOut.UpdateSettings(ctx, &OptOutSettings{Brand: "Acme"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Brand != "Acme" {
		t.Errorf("Unexpected: %+v", result)
	}
}
