package calendar

import (
	"errors"
	"unicode/utf8"
)

// Event struct
type Event struct {
	title string
	Date
}

// Title returns Event title.
func (e *Event) Title() string {
	return e.title
}

// SetTitle method for Event.
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 40 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
