package go1688

import (
	"strings"
	"time"
)

const JSONTIME_FORMAT = "20060102150405.000Z0700"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	builder := GetBytesBuffer()
	defer PutBytesBuffer(builder)
	builder.WriteString(`"`)
	builder.WriteString(t.Format())
	builder.WriteString(`"`)
	return builder.Bytes(), nil
}

func (t JsonTime) Format() string {
	formated := time.Time(t).Format(JSONTIME_FORMAT)
	return strings.Replace(formated, ".", "", -1)
}

func (t *JsonTime) UnmarshalJSON(s []byte) (err error) {
	if len(s) == 2 {
		t = nil
		return
	}
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	for idx, c := range s {
		if idx == 15 {
			buf.WriteByte('.')
		}
		buf.WriteByte(c)
	}
	parsed, err := time.Parse(`"`+JSONTIME_FORMAT+`"`, buf.String())
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

// Bool support string or bool for umarshal json
type Bool bool

func (b *Bool) UnmarshalJSON(s []byte) (err error) {
	var val bool
	str := string(s)
	if str == `"true"` || str == "true" {
		val = true
	}
	*b = Bool(val)
	return
}

func (b Bool) Bool() bool {
	return bool(b)
}
