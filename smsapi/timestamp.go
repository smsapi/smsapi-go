package smsapi

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	DateLayout = "2006-01-02"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Time.String()
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.Time.Format(time.RFC3339) + `"`), nil
}

func (t *Timestamp) UnmarshalJSON(buf []byte) error {
	var timestamp = new(int64)

	err := json.Unmarshal(buf, timestamp)

	if err == nil {
		tm := time.Unix(*timestamp, 0)

		*t = Timestamp{tm}

		return nil
	}

	var dtStr = new(string)

	err = json.Unmarshal(buf, dtStr)

	i, err := strconv.ParseInt(*dtStr, 10, 64)

	if err == nil {
		t.Time = time.Unix(i, 0)
	} else {
		t.Time, err = time.Parse(time.RFC3339, *dtStr)
	}

	return err
}

func (t Timestamp) Equal(c Timestamp) bool {
	return t.Time.Equal(c.Time)
}

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d *Date) UnmarshalJSON(buf []byte) error {
	var dStr = new(string)

	if err := json.Unmarshal(buf, dStr); err != nil {
		return err
	}

	tm, err := time.Parse(DateLayout, *dStr)

	if err == nil {
		d.Month = tm.Month()
		d.Year = tm.Year()
		d.Day = tm.Day()
	}

	return err
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d Date) Equal(c Date) bool {
	return d.Year == c.Year && d.Month == c.Month && d.Day == c.Day
}

func (d *Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", d.Year, d.Month, d.Day)
}

func NewDate(year, month, day int) *Date {
	return &Date{
		Year:  year,
		Month: time.Month(month),
		Day:   day,
	}
}
