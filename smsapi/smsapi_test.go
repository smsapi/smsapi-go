package smsapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	ctx = context.Background()
)

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	handler := http.NewServeMux()
	handler.Handle("/", mux)

	server := httptest.NewServer(handler)

	apiUrl := fmt.Sprintf("%s/", server.URL)

	client = NewPlClient("", nil)
	baseUrl, _ := url.Parse(apiUrl)
	client.BaseUrl = baseUrl

	return client, mux, server.Close
}

func TestJsonRequest(t *testing.T) {
	accessToken := "fake-token"
	client := NewPlClient(accessToken, http.DefaultClient)

	req, _ := client.NewJsonRequest("", "", "")

	expectedHeaders := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", accessToken)},
		"User-Agent":    {Name + "/" + Version},
	}

	if !reflect.DeepEqual(expectedHeaders, req.Header) {
		t.Errorf("Expected: %v Given: %v", expectedHeaders, req.Header)
	}
}

func TestUrlencodedRequest(t *testing.T) {
	accessToken := "fake-token"
	client := NewPlClient(accessToken, http.DefaultClient)

	req, _ := client.NewUrlencodedRequest("", "", "")

	expectedHeaders := http.Header{
		"Content-Type":  {"application/x-www-form-urlencoded"},
		"Authorization": {fmt.Sprintf("Bearer %s", accessToken)},
		"User-Agent":    {Name + "/" + Version},
	}

	if !reflect.DeepEqual(expectedHeaders, req.Header) {
		t.Errorf("Expected: %v Given: %v", expectedHeaders, req.Header)
	}
}

func assertRequestBody(t *testing.T, r *http.Request, given interface{}, expected interface{}) {
	json.NewDecoder(r.Body).Decode(given)

	if !reflect.DeepEqual(expected, given) {
		t.Errorf("Body expected: %v Given: %v", expected, given)
	}
}

func assertRequestUrlencoded(t *testing.T, r *http.Request, expected interface{}) {
	r.ParseMultipartForm(0)

	values, _ := query.Values(expected)

	for k := range values {
		expected := values.Get(k)

		if expected != "" {
			given := r.Form.Get(k)

			if given != expected {
				t.Errorf("Body expected: %v Given: %v", expected, given)
			}
		}
	}
}

func assertRequestQueryParam(t *testing.T, r *http.Request, key, value string) {
	q := r.URL.Query()

	k := q.Get(key)

	if k != value {
		t.Errorf("Querye expected: %s=%s", key, value)
	}
}

func assertRequestMethod(t *testing.T, r *http.Request, method string) {
	if r.Method != method {
		t.Errorf("Request method error: Expected: %s Given: %s", method, r.Method)
	}
}

func assertRequestJsonContains(t *testing.T, r *http.Request, expectedKey, expectedValue string) {
	payload, _ := simplejson.NewFromReader(r.Body)

	given := payload.Get(expectedKey).MustString()

	if given != expectedValue {
		t.Errorf("Invalid JSON: expected: %s=%s given: %s", expectedKey, expectedValue, given)
	}
}

func readFixture(fixtureFile string) string {
	c, err := ioutil.ReadFile(fmt.Sprintf("../testdata/responses/%s", fixtureFile))

	if err != nil {
		log.Fatal(err)
	}

	return string(c)
}
