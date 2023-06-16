package utils

import (
	"fmt"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

const layout = "2006-01-02"

func (c *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) // remove quotes
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

func (c Date) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}
