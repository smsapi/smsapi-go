package e2e

import (
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"testing"
	"time"
)

var scheduledSmsId string

func TestSendSms(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Sms.Send(ctx, phoneNumber, "go-smsapi", "")

	if err != nil {
		log.Fatal(err)
	}
}

func TestScheduleSms(t *testing.T) {
	future := &smsapi.Timestamp{time.Now().Local().AddDate(0, 0, 2)}

	ctx, cancel := createCtx()
	defer cancel()

	r, err := client.Sms.Schedule(ctx, phoneNumber, "go-smsapi", "", future)

	if err != nil {
		log.Fatal(err)
	}

	scheduledSmsId = r.Collection[0].Id
}

func TestGetSms(t *testing.T) {
	if scheduledSmsId == "" {
		t.Skip()
	}

	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Sms.Get(ctx, scheduledSmsId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveScheduleSms(t *testing.T) {
	if scheduledSmsId == "" {
		t.Skip()
	}

	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Sms.RemoveScheduled(ctx, scheduledSmsId)

	if err != nil {
		log.Fatal(err)
	}
}
