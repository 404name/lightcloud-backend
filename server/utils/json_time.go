package utils

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type JSONTime time.Time

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = JSONTime(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = JSONTime(now)
	return
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t JSONTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

func (t *JSONTime) Scan(v interface{}) error {
	tTime, _ := time.Parse("2006-01-02 15:04:05 +0800 CST", v.(time.Time).String())
	*t = JSONTime(tTime)
	return nil
}

func (t JSONTime) String() string {
	return time.Time(t).Format(TimeFormat)
}
