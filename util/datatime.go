package util

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime time.Time

const ctLayout = "2006-01-02"

// UnmarshalJSON Parses the json string in the custom format
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	nt, err := time.Parse(ctLayout, s)
	if err == nil {
		*ct = CustomTime(nt)
	} else {
		if !nt.IsZero() {
			*ct = CustomTime(nt)
		}
	}

	return
}

// MarshalJSON writes a quoted string in the custom format
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.String()), nil
}

// String returns the time in the custom format
func (ct *CustomTime) String() string {
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(ctLayout))
}

// String returns the time in the custom format
func (ct *CustomTime) Format(layout string) string {
	if ct.IsZero() {
		return ""
	}
	t := time.Time(*ct)
	return fmt.Sprintf("%q", t.Format(layout))
}

// String returns the time in the custom format
func (ct *CustomTime) IsZero() bool {
	t := time.Time(*ct)
	if t.IsZero() {
		return true
	} else {
		return false
	}
}
