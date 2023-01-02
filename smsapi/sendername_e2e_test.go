// +build e2e

package smsapi

import (
	"log"
	"testing"
)

var (
	senderName string
)

func TestCreateSenderName(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	result, err := client.Sender.Create(ctx, "custom")

	if err != nil {
		log.Fatal(err)
	}

	senderName = result.Name
}

func TestGetSenderName(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Sender.Get(ctx, senderName)

	if err != nil {
		log.Fatal(err)
	}
}

func TestGetAll(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Sender.GetAll(ctx)

	if err != nil {
		log.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	err := client.Sender.Delete(ctx, senderName)

	if err != nil {
		log.Fatal(err)
	}
}
