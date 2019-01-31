package zoop

import (
	"time"
)

const DateFormat = `"2006-01-02"`

type Date time.Time

func NewDate(t time.Time) Date {
	return Date(time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()))
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(d).Format(DateFormat)), nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(DateFormat, string(data))
	if err == nil {
		*d = Date(t)
	}
	return err
}

func (d Date) Pointer() *Date {
	return &d
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d Date) String() string {
	return d.Time().String()
}
