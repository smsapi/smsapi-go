package smsapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/url"
	"strconv"
)

const DefaultPageSize = 100

var NoMoreResults = errors.New("no more results")

type PaginationFilters struct {
	Offset uint `url:"offset,omitempty"`
	Limit  uint `url:"limit,omitempty"`
}

type PageIterator struct {
	Client  *Client
	Context context.Context
	Filters url.Values
	Uri     string
	Size    uint
	Limit   uint
	Offset  uint
}

func (i *PageIterator) Next(result Collection) error {
	if i.noMoreResults() {
		return NoMoreResults
	}

	err := i.nextPage(result)

	if err != nil {
		return err
	}

	i.Offset = i.Offset + i.Limit
	i.Size = result.GetSize()

	return nil
}

func (i *PageIterator) nextPage(result Collection) error {
	i.Filters.Set("offset", fmt.Sprint(i.Offset))
	i.Filters.Set("limit", fmt.Sprint(i.Limit))

	u, err := url.Parse(i.Uri)

	if err != nil {
		return err
	}

	u.RawQuery = i.Filters.Encode()

	err = i.Client.Get(i.Context, u.String(), result)

	return err
}

func (i *PageIterator) noMoreResults() bool {
	return i.Size > 0 && i.Offset > i.Size
}

func NewPageIterator(c *Client, ctx context.Context, uri string, v interface{}) *PageIterator {
	var offset, limit uint64

	filters, err := query.Values(v)

	if err != nil {
		filters = url.Values{}
	}

	offsetStr := filters.Get("offset")
	offset, err = strconv.ParseUint(offsetStr, 10, 32)

	if err != nil {
		offset = 0
	}

	limitStr := filters.Get("limit")
	limit, err = strconv.ParseUint(limitStr, 10, 32)

	if err != nil || limit == 0 {
		limit = DefaultPageSize
	}

	return &PageIterator{
		Client:  c,
		Context: ctx,
		Uri:     uri,
		Filters: filters,
		Offset:  uint(offset),
		Limit:   uint(limit),
	}
}
