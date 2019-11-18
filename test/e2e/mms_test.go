package e2e

import (
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"testing"
	"time"
)

var scheduledMmsId string

const (
	smsapiLogoUrl = "https://www.smsapi.pl/public/images/logo.1fa4730ff6f8a8c8755ecdd76acdf960.png"
)

func TestSendMms(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Mms.Send(ctx, phoneNumber, "go-smsapi", smsapiLogoUrl)

	if err != nil {
		log.Fatal(err)
	}
}

func TestScheduleMms(t *testing.T) {
	future := &smsapi.Timestamp{time.Now().Local().AddDate(0, 0, 2)}

	ctx, cancel := createCtx()
	defer cancel()

	r, err := client.Mms.Schedule(ctx, phoneNumber, "go-smsapi", smsapiLogoUrl, future)

	if err != nil {
		log.Fatal(err)
	}

	scheduledMmsId = r.Collection[0].Id
}

func TestGetMms(t *testing.T) {
	if scheduledMmsId == "" {
		t.Skip()
	}

	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Mms.Get(ctx, scheduledMmsId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveScheduledMms(t *testing.T) {
	if scheduledMmsId == "" {
		t.Skip()
	}

	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Mms.RemoveScheduled(ctx, scheduledMmsId)

	if err != nil {
		log.Fatal(err)
	}
}
