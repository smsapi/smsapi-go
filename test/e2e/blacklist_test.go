package e2e

import (
	"context"
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"testing"
)

var (
	blacklistPhoneNumberId string
)

func TestAddPhoneNumberToBlacklist(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	result, err := client.Blacklist.AddPhoneNumber(ctx, "656545434", nil)

	if err != nil {
		log.Fatal(err)
	}

	blacklistPhoneNumberId = result.Id
}

func TestGetAllPhoneNumbers(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	result, err := client.Blacklist.GetPhoneNumbers(ctx, &smsapi.BlacklistPhoneNumbersCollectionFilters{})

	if err != nil {
		log.Fatal(err)
	}

	if result.Size != 1 {
		log.Fatal("Invalid collection size")
	}
}

func TestDeletePhoneNumber(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Blacklist.DeletePhoneNumber(ctx, blacklistPhoneNumberId)

	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Blacklist.GetPhoneNumbers(ctx, &smsapi.BlacklistPhoneNumbersCollectionFilters{})

	if err != nil {
		log.Fatal(err)
	}

	if result.Size != 0 {
		log.Fatal("Collection should be empty")
	}
}

func TestDeleteAllPhoneNumbers(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Blacklist.AddPhoneNumber(ctx, "656545434", nil)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Blacklist.DeleteAllPhoneNumbers(ctx)

	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Blacklist.GetPhoneNumbers(ctx, &smsapi.BlacklistPhoneNumbersCollectionFilters{})

	if err != nil {
		log.Fatal(err)
	}

	if result.Size != 0 {
		log.Fatal("Collection should be empty")
	}
}

func addPhoneNumber(ctx context.Context, phoneNumber string, expirationDate *smsapi.Date) *smsapi.BlackListPhoneNumber {
	res, err := client.Blacklist.AddPhoneNumber(ctx, phoneNumber, expirationDate)

	if err != nil {
		log.Fatal(err)
	}

	return res
}
