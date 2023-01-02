// +build e2e

package smsapi

import (
	"context"
	"os"
	"time"
)

var (
	phoneNumber string

	client *Client
)

func init() {
	phoneNumber = os.Getenv("PHONE_NUMBER")

	client = NewPlClient(os.Getenv("SMSAPI_ACCESS_TOKEN"), nil)
}

func createCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 20*time.Second)
}
