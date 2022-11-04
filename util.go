package go1688

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

func structToMap(data interface{}) map[string]string {
	elem := reflect.ValueOf(data).Elem()
	relType := elem.Type()
	totalFields := relType.NumField()
	m := make(map[string]string, totalFields)
	for i := 0; i < totalFields; i++ {
		relField := relType.Field(i)
		kind := relField.Type.Kind()
		field := elem.Field(i)
		var val string
		switch kind {
		case reflect.Bool:
			val = "false"
			if field.Bool() {
				val = "true"
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = strconv.FormatInt(field.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = strconv.FormatUint(field.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			val = strconv.FormatFloat(field.Float(), 'f', -1, 64)
		case reflect.String:
			val = field.String()
		default:
			continue
		}
		if val == "" {
			continue
		}
		name := relField.Name
		if tagName := relField.Tag.Get("json"); tagName != "" {
			if strings.HasSuffix(tagName, ",omitempty") {
				name = strings.TrimSuffix(tagName, ",omitempty")
			} else {
				name = tagName
			}
		}
		m[name] = val
	}
	return m
}

var stringsBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func GetStringsBuilder() *strings.Builder {
	builder := stringsBuilderPool.Get().(*strings.Builder)
	builder.Reset()
	return builder
}

func PutStringsBuilder(builder *strings.Builder) {
	builder.Reset()
	stringsBuilderPool.Put(builder)
}

var bytesBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBytesBuffer() *bytes.Buffer {
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func PutBytesBuffer(buf *bytes.Buffer) {
	bytesBufferPool.Put(buf)
}

// JSONMarshal encode json without html escape
func JSONMarshal(req interface{}) string {
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(req)
	return buf.String()
}
