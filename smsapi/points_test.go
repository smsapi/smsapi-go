package smsapi

import (
	"encoding/json"
	"testing"
)

func TestPointsUnmarshalNumber(t *testing.T) {
	var p Points
	if err := json.Unmarshal([]byte(`0.1`), &p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != 0.1 {
		t.Errorf("expected 0.1, got %v", p)
	}
}

func TestPointsUnmarshalString(t *testing.T) {
	var p Points
	if err := json.Unmarshal([]byte(`"0.3000"`), &p); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != 0.3 {
		t.Errorf("expected 0.3, got %v", p)
	}
}

func TestPointsUnmarshalEmptyAndNull(t *testing.T) {
	cases := []string{`null`, `""`}
	for _, c := range cases {
		var p Points = 42
		if err := json.Unmarshal([]byte(c), &p); err != nil {
			t.Errorf("%s: unexpected error: %v", c, err)
		}
		if p != 42 {
			t.Errorf("%s: value was overwritten: %v", c, p)
		}
	}
}

func TestPointsUnmarshalInvalid(t *testing.T) {
	var p Points
	if err := json.Unmarshal([]byte(`"not-a-number"`), &p); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSmsResponsePointsAsString(t *testing.T) {
	payload := []byte(`{"id":"1","points":"0.3000","number":"48100200300"}`)

	var got SmsResponse
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Points != 0.3 {
		t.Errorf("Points: expected 0.3, got %v", got.Points)
	}
	if got.Id != "1" || got.Number != "48100200300" {
		t.Errorf("other fields not preserved: %+v", got)
	}
}

func TestMmsResponsePointsAsString(t *testing.T) {
	payload := []byte(`{"id":"1","points":"0.3000","number":"48100200300"}`)

	var got MmsResponse
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Points != 0.3 {
		t.Errorf("Points: expected 0.3, got %v", got.Points)
	}
}

func TestVmsResponsePointsAsString(t *testing.T) {
	payload := []byte(`{"id":"1","points":"0.3000","number":"48100200300"}`)

	var got VmsResponse
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Points != 0.3 {
		t.Errorf("Points: expected 0.3, got %v", got.Points)
	}
}

func TestProfileDetailsPointsAsString(t *testing.T) {
	payload := []byte(`{"name":"test","points":"12.5000"}`)

	var got ProfileDetailsResponse
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Points != 12.5 {
		t.Errorf("Points: expected 12.5, got %v", got.Points)
	}
}

func TestUserPointsAsString(t *testing.T) {
	payload := []byte(`{"from_account":"100.0000","per_month":"10.5000"}`)

	var got UserPoints
	if err := json.Unmarshal(payload, &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.FromAccount != 100 {
		t.Errorf("FromAccount: expected 100, got %v", got.FromAccount)
	}
	if got.PerMonth != 10.5 {
		t.Errorf("PerMonth: expected 10.5, got %v", got.PerMonth)
	}
}
