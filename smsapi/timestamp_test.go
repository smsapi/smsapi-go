package smsapi

import (
	"encoding/json"
	"testing"
	"time"
)

const (
	utcTimeStr   = `"2019-01-01T00:00:00Z"`
	emptyTimeStr = `"0001-01-01T00:00:00Z"`
	unixStr      = `1546300800`
)

var (
	someTime = time.Date(2019, time.January, 01, 00, 0, 0, 0, time.UTC)
)

func TestMarshalDateTimeMismatch(t *testing.T) {
	timestamp := Timestamp{}

	assertMarshalDateTime(t, utcTimeStr, timestamp, false)
}

func TestMarshalUTCTime(t *testing.T) {
	timestamp := Timestamp{someTime}

	assertMarshalDateTime(t, utcTimeStr, timestamp, true)
}

func TestMarshalEmptyTime(t *testing.T) {
	timestamp := Timestamp{}

	assertMarshalDateTime(t, emptyTimeStr, timestamp, true)
}

func TestUnmarshalUTCTime(t *testing.T) {
	expected := Timestamp{someTime}

	assertUnmarshalDateTime(t, expected, utcTimeStr, true)
}

func TestUnmarshalUnixTimestamp(t *testing.T) {
	expected := Timestamp{someTime}

	assertUnmarshalDateTime(t, expected, unixStr, true)
}

func TestUnmarshalEmptyTime(t *testing.T) {
	expected := Timestamp{}

	assertUnmarshalDateTime(t, expected, emptyTimeStr, true)
}

func TestUnmarshalMismatch(t *testing.T) {
	expected := Timestamp{}

	assertUnmarshalDateTime(t, expected, utcTimeStr, false)
}

func TestUnmarshalInvalid(t *testing.T) {
	var result Timestamp
	invalidTime := `"invalid-date-time"`

	err := json.Unmarshal([]byte(invalidTime), &result)

	if err == nil {
		t.Errorf("Unmarshal should fail for time: %s", invalidTime)
	}
}

func assertUnmarshalDateTime(t *testing.T, expected Timestamp, given string, equal bool) {
	var result Timestamp

	err := json.Unmarshal([]byte(given), &result)

	if err != nil {
		t.Errorf("Can not Unmarshal: %s", given)
	}

	tEquals := expected.Equal(result)

	if tEquals != equal {
		t.Errorf("Unmarshal error Given: %s expected: %s", result, expected)
	}
}

func assertMarshalDateTime(t *testing.T, expected string, given Timestamp, equal bool) {
	result, err := given.MarshalJSON()

	if err != nil {
		t.Errorf("Can not Marshal: %s", given)
	}

	marshalResult := string(result)

	if (marshalResult == expected) != equal {
		t.Errorf("Marshal error Given: %s expected: %s", marshalResult, expected)
	}
}
