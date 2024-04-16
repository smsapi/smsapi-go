package e2e

import (
	"log"
	"testing"
)

func TestGetProfileDetails(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Profile.Details(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
