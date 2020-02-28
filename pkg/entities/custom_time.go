package entities

import (
	"fmt"
	"strings"
	"time"
)

const ctLayout = "Monday, January 2, 2006 3:04 PM"

var nilTime = (time.Time{}).UnixNano()

// CustomTime - custom time
type CustomTime struct {
	time.Time
}

// UnmarshalJSON - unmarshalling json
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	_ = time.ANSIC
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

// MarshalJSON - marshalling json
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}
