package smsapi

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShipmentListCountryVolumes(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/shipment/country_volumes", func(w http.ResponseWriter, r *http.Request) {
		assertRequestMethod(t, r, "GET")
		assertRequestQueryParam(t, r, "year", "2024")
		assertRequestQueryParam(t, r, "month", "01")
		fmt.Fprint(w, `{"size":1,"collection":[{"country_code":"PL","usage":42}]}`)
	})

	result, err := client.Shipment.ListCountryVolumes(ctx, &ShipmentCountryVolumeFilters{
		Year:  "2024",
		Month: "01",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Size != 1 || result.Collection[0].CountryCode != "PL" || result.Collection[0].Usage != 42 {
		t.Errorf("Unexpected: %+v", result)
	}
}
