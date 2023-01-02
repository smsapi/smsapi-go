// +build unit

package smsapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestGetClicks(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/short_url/clicks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("short_url/clicks.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.ShortUrl.GetClicks(ctx, nil)

	expected := &ClicksCollectionResponse{
		Size: 1,
		Collection: []*ClickResponse{
			createClickResponse(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestCreateReport(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	filters := &ClicksCollectionFilters{
		DateFrom: "2019-01-01",
		DateTo:   "2020-01-01",
		LinkIds:  []string{"1"},
	}

	mux.HandleFunc("/short_url/clicks_reports", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("short_url/clicks_report.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestQueryParam(t, r, "date_from", "2019-01-01")
		assertRequestQueryParam(t, r, "date_to", "2020-01-01")
		assertRequestQueryParam(t, r, "links", "1")
	})

	result, _ := client.ShortUrl.CreateReport(ctx, filters)

	expected := &ClicksReportResponse{
		ReportUrl: "https://api.smsapi.com/short_url/clicks_reports/cb211ef485719e3c6357ceb38",
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetLinks(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/short_url/links", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("short_url/links.json"))

		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.ShortUrl.GetLinks(ctx)

	expected := &LinksCollectionResponse{
		Size: 1,
		Collection: []*LinkResponse{
			createLinkResponse(),
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestGetLink(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/short_url/links/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("short_url/link.json"))
		assertRequestMethod(t, r, "GET")
	})

	result, _ := client.ShortUrl.GetLink(ctx, "1")

	expected := createLinkResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestCreateLink(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	link := &Link{
		Url:         "https://some-url",
		Description: "test",
		Type:        linkTypeUrl,
	}

	mux.HandleFunc("/short_url/links", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("short_url/link.json"))

		assertRequestMethod(t, r, "POST")
		assertRequestUrlencoded(t, r, link)
	})

	result, _ := client.ShortUrl.CreateLink(ctx, "https://some-url", "test", "test")

	expected := createLinkResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func TestUpdateLink(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	link := &Link{
		Description: "demo",
	}

	mux.HandleFunc("/short_url/links/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, readFixture("short_url/link.json"))

		assertRequestMethod(t, r, "PUT")
		assertRequestUrlencoded(t, r, link)
	})

	result, _ := client.ShortUrl.UpdateLink(ctx, "1", "", "demo", "demo")

	expected := createLinkResponse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Given: %+v Expected: %+v", result, expected)
	}
}

func createClickResponse() *ClickResponse {
	return &ClickResponse{
		Name:        "test1",
		Suffix:      "test1",
		ShortUrl:    "https://cut.li/test1",
		PhoneNumber: "100200300",
		DateHit:     "1970-01-01T00:00:00+0100",
		Os:          "Windows",
		Browser:     "Firefox",
		Device:      "PC",
	}
}

func createLinkResponse() *LinkResponse {
	return &LinkResponse{
		Id:          "1",
		Name:        "test",
		Url:         "https://localhost.com",
		ShortUrl:    "http://idz.do/test",
		Filename:    "",
		Type:        "URL",
		Expire:      &Timestamp{time.Date(2037, time.December, 02, 15, 14, 36, 0, time.UTC)},
		Hits:        0,
		HitsUnique:  0,
		Description: "test",
	}
}
