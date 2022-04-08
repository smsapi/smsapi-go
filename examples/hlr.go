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

	result, err := client.Hlr.CheckNumber(context.Background(), "+48695418520")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Id, result.Status, result.Number, result.Price)
}
