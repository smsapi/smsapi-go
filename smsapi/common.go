package smsapi

import (
	"github.com/google/go-querystring/query"
	"net/url"
)

func addQueryParams(uri string, parameters interface{}) (string, error) {
	u, err := url.Parse(uri)

	if err != nil {
		return uri, err
	}

	q, err := query.Values(parameters)

	if err != nil {
		return uri, err
	}

	currentQuery := u.Query()

	for k := range currentQuery {
		q.Add(k, currentQuery.Get(k))
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

