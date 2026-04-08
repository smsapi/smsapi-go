package smsapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCallbacksList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"size":1,"collection":[{"id":"1","url":"http://example.com","type":"sms_dlr","active":true}]}`)
	})

	result, err := client.Callbacks.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result.Size != 1 || len(result.Collection) != 1 || result.Collection[0].Id != "1" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestCallbacksGet(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"1","url":"http://example.com","type":"sms_dlr","active":true}`)
	})

	result, err := client.Callbacks.Get(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if result.Url != "http://example.com" || result.Type != "sms_dlr" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestCallbacksCreate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
		assertRequestJsonContains(t, r, "url", "http://example.com")
		fmt.Fprint(w, `{"id":"1","url":"http://example.com","type":"sms_dlr","active":true}`)
	})

	cb := &Callback{Url: "http://example.com", Type: "sms_dlr"}
	result, err := client.Callbacks.Create(ctx, cb)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != "1" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestCallbacksUpdate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
		assertRequestJsonContains(t, r, "url", "http://new.example.com")
		fmt.Fprint(w, `{"id":"1","url":"http://new.example.com","type":"sms_dlr","active":true}`)
	})

	result, err := client.Callbacks.Update(ctx, "1", "http://new.example.com")
	if err != nil {
		t.Fatal(err)
	}
	if result.Url != "http://new.example.com" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestCallbacksDelete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	if err := client.Callbacks.Delete(ctx, "1"); err != nil {
		t.Fatal(err)
	}
}

func TestCallbacksActivate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks/1/commands/activate", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
	})

	if err := client.Callbacks.Activate(ctx, "1"); err != nil {
		t.Fatal(err)
	}
}

func TestCallbacksDeactivate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks/1/commands/deactivate", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
	})

	if err := client.Callbacks.Deactivate(ctx, "1"); err != nil {
		t.Fatal(err)
	}
}

func TestCallbacksTest(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/callbacks/1/commands/test", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"test_result":true,"connection_failed":false,"invalid_encoding":false,"http_response_status":200,"http_response_body":"ok"}`)
	})

	result, err := client.Callbacks.Test(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if !result.TestResult || result.HttpResponseStatus == nil || *result.HttpResponseStatus != 200 {
		t.Errorf("Unexpected: %+v", result)
	}
}
