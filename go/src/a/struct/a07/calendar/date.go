package calendar

import "errors"

// Date struct.
type Date struct {
	year  int
	month int
	day   int
}

// Year returns Date year.
func (d *Date) Year() int {
	return d.year
}

// Month returns Date month.
func (d *Date) Month() int {
	return d.month
}

// Day returns Date day.
func (d *Date) Day() int {
	return d.day
}

// SetYear method for Date.
func (d *Date) SetYear(year int) error {
	if year < 1 || year > 9999 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

// SetMonth method for Date.
func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

// SetDay method for Date.
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
