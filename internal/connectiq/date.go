package connectiq

import "time"

type DateTime struct {
	date time.Time
}

func (d *DateTime) UnmarshalText(text []byte) error {
	t, err := time.ParseInLocation(time.DateTime, string(text), time.UTC)
	if err != nil {
		return err
	}
	*d = DateTime{date: t}

	return nil
}

func (d *DateTime) String() string {
	return d.date.UTC().Format(time.DateTime)
}

func (d *DateTime) Time() time.Time {
	return d.date
}
