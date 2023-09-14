package e2e

import (
	"log"
	"testing"
)

func TestHlr(t *testing.T) {
	ctx, cancel := createCtx()
	defer cancel()

	_, err := client.Hlr.CheckNumber(ctx, "695418520")

	if err != nil {
		log.Fatal(err)
	}
}
