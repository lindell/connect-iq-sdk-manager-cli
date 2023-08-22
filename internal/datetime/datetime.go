package datetime

import "time"

type DateTime struct {
	date time.Time
}

func Now() DateTime {
	return DateTime{
		date: time.Now().In(time.UTC),
	}
}

func Parse(str string) (DateTime, error) {
	t, err := time.ParseInLocation(time.DateTime, str, time.UTC)
	if err != nil {
		return DateTime{}, err
	}
	return DateTime{date: t}, nil
}

func (d *DateTime) UnmarshalText(text []byte) error {
	t, err := time.ParseInLocation(time.DateTime, string(text), time.UTC)
	if err != nil {
		return err
	}
	*d = DateTime{date: t}

	return nil
}

func (d DateTime) String() string {
	return d.date.UTC().Format(time.DateTime)
}

func (d DateTime) Time() time.Time {
	return d.date
}
