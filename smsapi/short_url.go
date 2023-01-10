package smsapi

import (
	"context"
	"fmt"
	"net/http"
)

type ShortUrlApi struct {
	client *Client
}

type ClicksCollectionFilters struct {
	DateFrom string   `url:"date_from,omitempty"`
	DateTo   string   `url:"date_to,omitempty"`
	LinkIds  []string `url:"links,omitempty"`
}

type ClickResponse struct {
	Name        string `json:"name"`
	ShortUrl    string `json:"short_url"`
	PhoneNumber string `json:"phone_number"`
	Suffix      string `json:"suffix"`
	DateHit     string `json:"date_hit"`
	Os          string `json:"os"`
	Browser     string `json:"browser"`
	Device      string `json:"device"`
}

type ClicksCollectionResponse struct {
	Size       int              `json:"size"`
	Collection []*ClickResponse `json:"collection"`
}

func (shortUrlApi *ShortUrlApi) GetClicks(ctx context.Context, filters *ClicksCollectionFilters) (*ClicksCollectionResponse, error) {
	uri, _ := addQueryParams("/short_url/clicks", filters)

	var result = new(ClicksCollectionResponse)

	err := shortUrlApi.client.Get(ctx, uri, result)

	return result, err
}

type ClicksReportResponse struct {
	ReportUrl string `json:"link"`
}

func (shortUrlApi *ShortUrlApi) CreateReport(ctx context.Context, filters *ClicksCollectionFilters) (*ClicksReportResponse, error) {
	var result = new(ClicksReportResponse)

	uri, _ := addQueryParams("/short_url/clicks_reports", filters)

	err := shortUrlApi.client.Post(ctx, uri, result, nil)

	return result, err
}

type LinkType string
type ExpireTimeUnit string

const (
	linkTypeUrl = LinkType("URL")
	ExpireTimeSeconds = ExpireTimeUnit("seconds")
	ExpireTimeMinutes = ExpireTimeUnit("minutes")
	ExpireTimeHours = ExpireTimeUnit("hours")
	ExpireTimeDays = ExpireTimeUnit("days")
)

type Link struct {
	Url         string     `json:"url" url:"url,omitempty"`
	Name        string     `json:"name" url:"name,omitempty"`
	ExpireTime  int        `json:"expire_time,omitempty" url:"expire_time,omitempty"`
	ExpireUnit  ExpireTimeUnit `json:"expire_unit,omitempty" url:"expire_unit,omitempty"`
	Description string     `json:"description" url:"description,omitempty"`
	Type        LinkType   `json:"type" url:"type,omitempty"`
}

type LinkResponse struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Url         string     `json:"url"`
	ShortUrl    string     `json:"short_url"`
	Filename    string     `json:"filename"`
	Type        string     `json:"type"`
	Expire      *Timestamp `json:"expire"`
	Hits        int        `json:"hits"`
	HitsUnique  int        `json:"hits_unique"`
	Description string     `json:"description"`
}

type LinksCollectionResponse struct {
	Size       int             `json:"size"`
	Collection []*LinkResponse `json:"collection"`
}

func (shortUrlApi *ShortUrlApi) GetLinks(ctx context.Context) (*LinksCollectionResponse, error) {
	var result = new(LinksCollectionResponse)

	err := shortUrlApi.client.Get(ctx, "/short_url/links", result)

	return result, err
}

func (shortUrlApi *ShortUrlApi) GetLink(ctx context.Context, id string) (*LinkResponse, error) {
	var result = new(LinkResponse)

	uri := fmt.Sprintf("/short_url/links/%s", id)

	err := shortUrlApi.client.Get(ctx, uri, result)

	return result, err
}

func (shortUrlApi *ShortUrlApi) CreateLinkRaw(ctx context.Context, link *Link) (*LinkResponse, error) {
	var result = new(LinkResponse)

	err := shortUrlApi.client.Urlencoded(ctx, http.MethodPost, "/short_url/links", result, link)

	return result, err
}

func (shortUrlApi *ShortUrlApi) CreateLink(ctx context.Context, targetUrl, name, description string) (*LinkResponse, error) {
	link := &Link{
		Name:        name,
		Description: description,
		Url:         targetUrl,
		Type:        linkTypeUrl,
	}

	return shortUrlApi.CreateLinkRaw(ctx, link)
}

func (shortUrlApi *ShortUrlApi) UpdateLinkRaw(ctx context.Context, id string, link *Link) (*LinkResponse, error) {
	uri := fmt.Sprintf("/short_url/links/%s", id)

	var result = new(LinkResponse)

	err := shortUrlApi.client.Urlencoded(ctx, http.MethodPut, uri, result, link)

	return result, err
}

func (shortUrlApi *ShortUrlApi) UpdateLink(ctx context.Context, id, targetUrl, name, description string) (*LinkResponse, error) {
	link := &Link{
		Name:        name,
		Description: description,
		Url:         targetUrl,
		Type:        linkTypeUrl,
	}

	return shortUrlApi.UpdateLinkRaw(ctx, id, link)
}

func (shortUrlApi *ShortUrlApi) DeleteLink(ctx context.Context, id string) error {
	uri := fmt.Sprintf("/short_url/links/%s", id)

	return shortUrlApi.client.Delete(ctx, uri)
}
