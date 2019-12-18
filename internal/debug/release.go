// +build !debug

package debug

import (
	"encoding/json"
	"io"
	"net/url"
)

func DebugPrintError(err error) {}

func DebugPrintStringResponse(str string) {}

func DebugPrintGetRequest(url string) {}

func DebugPrintPostMapRequest(url string, data url.Values) {}

func DebugPrintPostJSONRequest(url string, body []byte) {}

func DebugPrintPostMultipartRequest(url string, body []byte) {}

func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
