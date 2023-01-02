// +build unit

package smsapi

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

type AnyCollection struct {
	CollectionMeta
}

var collection = &AnyCollection{
	CollectionMeta{
		Size: 100,
	},
}

func TestGetFirstPage(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	iterator := NewPageIterator(client, ctx, "/uri", nil)

	mux.HandleFunc("/uri", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{}")

		expected := url.Values{}
		expected.Add("offset", "0")
		expected.Add("limit", "100")

		assertPaginationFilters(t, expected, r.URL.Query())
	})

	iterator.Next(collection)
}

func TestGetNextPage(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	iterator := NewPageIterator(client, ctx, "/uri", nil)

	skipAssert := true

	mux.HandleFunc("/uri", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{}")

		expected := url.Values{}
		expected.Add("offset", "100")
		expected.Add("limit", "100")

		if !skipAssert {
			assertPaginationFilters(t, expected, r.URL.Query())
		}

		skipAssert = false
	})

	iterator.Next(collection)
	iterator.Next(collection)
}

func TestNoMoreResults(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	iterator := NewPageIterator(client, ctx, "/uri", nil)

	mux.HandleFunc("/uri", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{}")
	})

	iterator.Next(collection)
	iterator.Next(collection)

	err := iterator.Next(collection)

	if err != NoMoreResults {
		t.Errorf("Expected NoMoreResults given: %v", err)
	}
}

func assertPaginationFilters(t *testing.T, expected url.Values, given url.Values) {
	if !reflect.DeepEqual(expected, given) {
		t.Errorf("Query string not equals. Expected: %v Given :%v", expected, given)
	}
}
