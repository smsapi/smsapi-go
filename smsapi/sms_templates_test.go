package smsapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSmsTemplatesList(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sms/templates", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"size":1,"collection":[{"id":"1","name":"hello","template":"Hi [%name%]","normalize":true}]}`)
	})

	result, err := client.SmsTemplates.List(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result.Size != 1 || result.Collection[0].Id != "1" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestSmsTemplatesGet(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sms/templates/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":"1","name":"hello","template":"Hi [%name%]","normalize":true}`)
	})

	result, err := client.SmsTemplates.Get(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	if result.Name != "hello" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestSmsTemplatesCreate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	tpl := &SmsTemplate{Name: "hello", Template: "Hi [%name%]", Normalize: true}

	mux.HandleFunc("/sms/templates", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "POST")
		assertRequestUrlencoded(t, r, tpl)
		fmt.Fprint(w, `{"id":"1","name":"hello","template":"Hi [%name%]","normalize":true}`)
	})

	result, err := client.SmsTemplates.Create(ctx, tpl)
	if err != nil {
		t.Fatal(err)
	}
	if result.Id != "1" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestSmsTemplatesUpdate(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	tpl := &SmsTemplate{Name: "hello2", Template: "Hey [%name%]"}

	mux.HandleFunc("/sms/templates/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "PUT")
		assertRequestUrlencoded(t, r, tpl)
		fmt.Fprint(w, `{"id":"1","name":"hello2","template":"Hey [%name%]"}`)
	})

	result, err := client.SmsTemplates.Update(ctx, "1", tpl)
	if err != nil {
		t.Fatal(err)
	}
	if result.Name != "hello2" {
		t.Errorf("Unexpected: %+v", result)
	}
}

func TestSmsTemplatesDelete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sms/templates/1", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "DELETE")
	})

	if err := client.SmsTemplates.Delete(ctx, "1"); err != nil {
		t.Fatal(err)
	}
}

func TestSmsTemplatesListAvailable(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sms/templates/available", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		fmt.Fprint(w, `{"size":1,"collection":[{"name":"hello","template":"Hi","normalize":false}]}`)
	})

	result, err := client.SmsTemplates.ListAvailable(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if result.Size != 1 || result.Collection[0].Name != "hello" {
		t.Errorf("Unexpected: %+v", result)
	}
}
