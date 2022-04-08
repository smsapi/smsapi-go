package main

import (
	"context"
	"fmt"
	"github.com/smsapi/smsapi-go/smsapi"
	"log"
	"os"
)

func main() {
	accessToken := os.Getenv("SMSAPI_ACCESS_TOKEN")

	client := smsapi.NewPlClient(accessToken, nil)

	result, err := client.Sms.Send(context.Background(), "+48500500500", "go", "")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sent messages count", result.Count)

	for _, sms := range result.Collection {
		fmt.Println(sms.Id, sms.Status, sms.Points)
	}
}
