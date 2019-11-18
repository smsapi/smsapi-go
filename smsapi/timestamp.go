package smsapi

import (
	"strconv"
	"time"
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
	str := string(buf)

	i, err := strconv.ParseInt(str, 10, 64)

	if err == nil {
		t.Time = time.Unix(i, 0)
	} else {
		t.Time, err = time.Parse(`"`+time.RFC3339+`"`, str)
	}

	return err
}

func (t Timestamp) Equal(c Timestamp) bool {
	return t.Time.Equal(c.Time)
}
