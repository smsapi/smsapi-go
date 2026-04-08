package smsapi

import (
	"context"
)

type ShipmentApi struct {
	client *Client
}

type ShipmentCountryVolume struct {
	CountryCode string `json:"country_code"`
	Usage       int    `json:"usage"`
}

type ShipmentCountryVolumeCollection struct {
	Size       int                      `json:"size"`
	Collection []*ShipmentCountryVolume `json:"collection"`
}

type ShipmentCountryVolumeFilters struct {
	Year  string `url:"year,omitempty"`
	Month string `url:"month,omitempty"`
}

// ListCountryVolumes lists SMS shipment usage per country for a given year/month.
func (api *ShipmentApi) ListCountryVolumes(ctx context.Context, filters *ShipmentCountryVolumeFilters) (*ShipmentCountryVolumeCollection, error) {
	result := new(ShipmentCountryVolumeCollection)
	uri, _ := addQueryParams("/shipment/country_volumes", filters)
	err := api.client.Get(ctx, uri, result)
	return result, err
}
