package go1688

import (
	"fmt"
	"strings"
	"time"
)

const JSONTIME_FORMAT = "20060102150405.000Z0700"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	formated := time.Time(t).Format(JSONTIME_FORMAT)
	return []byte(fmt.Sprintf(`"%s"`, strings.Replace(formated, ".", "", -1))), nil
}

func (t *JsonTime) UnmarshalJSON(s []byte) (err error) {
	if len(s) == 2 {
		t = nil
		return
	}
	var buf []byte
	for idx, c := range s {
		if idx == 15 {
			buf = append(buf, '.')
		}
		buf = append(buf, c)
	}
	parsed, err := time.Parse(`"`+JSONTIME_FORMAT+`"`, string(buf))
	if err != nil {
		return err
	}
	*(*time.Time)(t) = parsed
	return
}

func (t JsonTime) IsZero() bool {
	return time.Time(t).IsZero()
}

func (t JsonTime) Time() time.Time { return time.Time(t) }

func (t JsonTime) String() string { return time.Time(t).String() }
