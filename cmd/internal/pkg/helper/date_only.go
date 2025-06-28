package helper

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type DateOnly struct {
	time.Time
}

const timeFormat = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(timeFormat, s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Format(timeFormat))), nil
}

func (DateOnly) GormDataType() string {
	return "date"
}

func (d DateOnly) Value() (driver.Value, error) {
	return d.Format(timeFormat), nil
}

func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
	case []byte:
		t, err := time.Parse(timeFormat, string(v))
		if err != nil {
			return err
		}
		d.Time = t
	case string:
		t, err := time.Parse(timeFormat, v)
		if err != nil {
			return err
		}
		d.Time = t
	default:
		return fmt.Errorf("cannot parse date: %v", value)
	}
	return nil
}
