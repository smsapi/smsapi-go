package e2e

import (
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"testing"
	"time"
)

var scheduledVmsId string

func TestSendVms(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Vms.Send(ctx, phoneNumber, "go-smsapi", "")

	if err != nil {
		log.Fatal(err)
	}
}

func TestScheduleVms(t *testing.T) {
	future := &smsapi.Timestamp{time.Now().Local().AddDate(0, 0, 2)}

	ctx, cancel := createCtx()
	defer cancel()

	r, err := client.Vms.Schedule(ctx, phoneNumber, "go-smsapi", "", future)

	if err != nil {
		log.Fatal(err)
	}

	scheduledVmsId = r.Collection[0].Id
}

func TestGetVms(t *testing.T) {
	if scheduledVmsId == "" {
		t.Skip()
	}

	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Vms.Get(ctx, scheduledVmsId)

	if err != nil {
		log.Fatal(err)
	}
}

func TestRemoveScheduledVms(t *testing.T) {
	if scheduledVmsId == "" {
		t.Skip()
	}

	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Vms.RemoveScheduled(ctx, scheduledVmsId)

	if err != nil {
		log.Fatal(err)
	}
}
