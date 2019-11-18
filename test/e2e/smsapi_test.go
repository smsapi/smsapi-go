package e2e

import (
	"context"
	"github.com/smsapi/smsapi-go/smsapi"
	"os"
	"time"
)

var (
	phoneNumber string

	client *smsapi.Client
)

func init() {
	phoneNumber = os.Getenv("PHONE_NUMBER")

	client = smsapi.NewPlClient(os.Getenv("SMSAPI_ACCESS_TOKEN"), nil)
}

func createCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 20*time.Second)
}
