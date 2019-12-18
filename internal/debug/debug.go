// +build debug

package debug

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/url"
)

func DebugPrintError(err error) {
	log.Println("[VERED_DEBUG] [ERROR]", err)
}

func DebugPrintStringResponse(str string) {
	log.Println("[VERED_DEBUG] [RESPONSE]", str)
}

func DebugPrintGetRequest(url string) {
	log.Println("[VERED_DEBUG] [API] GET", url)
}

func DebugPrintPostMapRequest(url string, data url.Values) {
	const format = "[VERED_DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"
	log.Printf(format, url, data.Encode())
}

func DebugPrintPostJSONRequest(url string, body []byte) {
	const format = "[VERED_DEBUG] [API] JSON POST %s\n" +
		"http request body:\n%s\n"

	buf := bytes.NewBuffer(make([]byte, 0, len(body)+1024))
	if err := json.Indent(buf, body, "", "    "); err == nil {
		body = buf.Bytes()
	}
	log.Printf(format, url, body)
}

func DebugPrintPostMultipartRequest(url string, body []byte) {
	log.Println("[VERED_DEBUG] [API] multipart/form-data POST", url)
}

func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	body2 := body
	buf := bytes.NewBuffer(make([]byte, 0, len(body2)+1024))
	if err := json.Indent(buf, body2, "", "    "); err == nil {
		body2 = buf.Bytes()
	}
	log.Printf("[VERED_DEBUG] [API] http response body:\n%s\n", body2)

	return json.Unmarshal(body, v)
}
