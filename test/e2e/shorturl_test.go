package e2e

import (
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"testing"
)

var (
	shortUrlId string
)

func TestCreateShortUrl(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	result, err := client.ShortUrl.CreateLink(ctx, "https://smsapi.pl", "go-smsapi", "go-smsapi")

	if err != nil {
		log.Fatal(err)
	}

	shortUrlId = result.Id
}

func TestUpdateShortUrl(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.ShortUrl.UpdateLink(ctx, shortUrlId, "https://smsapi.pl", "go-smsapi-pl", "go-smsapi-pl")

	if err != nil {
		log.Fatal(err)
	}
}

func TestGetShortUrl(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.ShortUrl.GetLink(ctx, shortUrlId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestGetAllShortUrls(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.ShortUrl.GetLinks(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func TestDeleteShortUrl(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.ShortUrl.DeleteLink(ctx, shortUrlId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateShortUrlWithExpire(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

    expireTime := 10
    expireUnit := smsapi.ExpireTimeDays

    link := &smsapi.Link{
		Name:        "go-smsapi-expire",
		Description: "go-smsapi-expire",
		Url:         "https://smsapi.pl",
		Type:        smsapi.LinkType("URL"),
		ExpireTime:  expireTime,
		ExpireUnit:  expireUnit,
	}

	result, err := client.ShortUrl.CreateLinkRaw(ctx, link)

	if err != nil {
		log.Fatal(err)
	}

	client.ShortUrl.DeleteLink(ctx, result.Id)
}
