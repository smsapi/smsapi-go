package smsapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-querystring/query"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Version        = "1.1.2"
	Name           = "smsapi-go"
	BaseUrlPl      = "https://api.smsapi.pl/"
	BaseUrlCom     = "https://api.smsapi.com/"
	DefaultTimeout = 30 * time.Second
)

type ContentType string

const (
	ContentTypeJson            = ContentType("application/json")
	ContentTypeXFormUrlencoded = ContentType("application/x-www-form-urlencoded")
)

type BearerAuth struct {
	AccessToken string
}

func (a *BearerAuth) String() string {
	return fmt.Sprintf("Bearer %s", a.AccessToken)
}

type Client struct {
	httpClient *http.Client

	BaseUrl *url.URL
	Auth    *BearerAuth

	Sms       *SmsApi
	Account   *AccountApi
	Contacts  *ContactsApi
	ShortUrl  *ShortUrlApi
	Hlr       *HlrApi
	Sender    *SenderApi
	Blacklist *BlacklistApi

	Mms *MmsApi
	Vms *VmsApi
}

func NewClient(apiUrl string, accessToken string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
		httpClient.Timeout = DefaultTimeout
	}

	baseUrl, _ := url.Parse(apiUrl)

	auth := &BearerAuth{AccessToken: accessToken}

	c := &Client{
		httpClient: httpClient,
		BaseUrl:    baseUrl,
		Auth:       auth,
	}

	c.Sms = &SmsApi{client: c}
	c.Account = &AccountApi{client: c}
	c.Contacts = &ContactsApi{client: c}
	c.ShortUrl = &ShortUrlApi{client: c}
	c.Hlr = &HlrApi{client: c}
	c.Sender = &SenderApi{client: c}
	c.Blacklist = &BlacklistApi{client: c}

	return c
}

func NewPlClient(accessToken string, httpClient *http.Client) *Client {
	smsapiClient := &*NewClient(BaseUrlPl, accessToken, httpClient)

	smsapiClient.Mms = &MmsApi{client: smsapiClient}
	smsapiClient.Vms = &VmsApi{client: smsapiClient}

	return smsapiClient
}

func NewInternationalClient(accessToken string, httpClient *http.Client) *Client {
	return &*NewClient(BaseUrlCom, accessToken, httpClient)
}

func (client *Client) NewUrlencodedRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.Reader

	if body != nil {
		data, _ := query.Values(body)
		buf = strings.NewReader(data.Encode())
	}

	return client.NewRequest(method, path, buf, ContentTypeXFormUrlencoded)
}

func (client *Client) NewJsonRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)

		encoder := json.NewEncoder(buf)
		err := encoder.Encode(body)

		if err != nil {
			return nil, err
		}
	}

	return client.NewRequest(method, path, buf, ContentTypeJson)
}

func (client *Client) NewRequest(method, path string, buf io.Reader, contentType ContentType) (*http.Request, error) {
	u, err := client.BaseUrl.Parse(path)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", string(contentType))
	req.Header.Set("Authorization", client.Auth.String())
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", Name, Version))

	return req, nil
}

func (client *Client) executeRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)

	resp, err := client.httpClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	responseData, err := client.CheckError(resp)

	if err != nil {
		return err
	}

	if v == nil {
		return nil
	}

	responseDataReader := bytes.NewReader(responseData)

	err = json.NewDecoder(responseDataReader).Decode(v)

	return err
}

var legacyQueryParams = struct {
	Format string `url:"format"`
}{Format: "json"}

func (client *Client) LegacyGet(ctx context.Context, path string, result interface{}) error {
	path, _ = addQueryParams(path, legacyQueryParams)

	req, err := client.NewJsonRequest("GET", path, nil)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, result)
}

func (client *Client) LegacyPost(ctx context.Context, path string, result interface{}, data interface{}) error {
	path, _ = addQueryParams(path, legacyQueryParams)

	req, err := client.NewJsonRequest("POST", path, data)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, result)
}

func (client *Client) Get(ctx context.Context, path string, result interface{}) error {
	req, err := client.NewJsonRequest("GET", path, nil)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, result)
}

func (client *Client) Urlencoded(ctx context.Context, method, path string, result interface{}, data interface{}) error {
	req, err := client.NewUrlencodedRequest(method, path, data)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, result)
}

func (client *Client) Post(ctx context.Context, path string, result interface{}, data interface{}) error {
	req, err := client.NewJsonRequest("POST", path, data)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, result)
}

func (client *Client) Put(ctx context.Context, path string, result interface{}, data interface{}) error {
	req, err := client.NewJsonRequest("PUT", path, data)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, result)
}

func (client *Client) Delete(ctx context.Context, path string) error {
	req, err := client.NewJsonRequest("DELETE", path, nil)

	if err != nil {
		return err
	}

	return client.executeRequest(ctx, req, nil)
}

func (client *Client) CheckError(r *http.Response) ([]byte, error) {
	errorResponse := &ErrorResponse{}

	responseData, err := ioutil.ReadAll(r.Body)

	if err == nil && responseData != nil {
		json.Unmarshal(responseData, errorResponse)
	}

	if httpCode := r.StatusCode; errorResponse.Code == 0 && 200 <= httpCode && httpCode <= 300 {
		return responseData, nil
	}

	errorResponse.Status = r.StatusCode

	return responseData, errorResponse
}
