package go1688

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bububa/go1688/crypto"
	"github.com/bububa/go1688/internal/debug"
)

const (
	GATEWAY  = "http://gw.open.1688.com/openapi"
	PROTOCOL = "param2"
	VERSION  = "1"
)

// Client api client
type Client struct {
	clt    *http.Client
	appKey string
	secret []byte
}

// NewClient create a new Client instance
func NewClient(appKey string, secret string, httpClient *http.Client) *Client {
	client := httpClient
	if httpClient == nil {
		client = http.DefaultClient
	}
	return &Client{
		clt:    client,
		appKey: appKey,
		secret: []byte(secret),
	}
}

// Do execute api request
func (c *Client) Do(req Request, accessToken string, resp Response) error {
	reqPath := c.requestPath(req)
	reqParams := req.Params()
	reqParams["_aop_timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	if accessToken != "" {
		reqParams["access_token"] = accessToken
	}
	reqParams["_aop_signature"] = c.Sign(reqPath, reqParams)
	values := url.Values{}
	for k, v := range reqParams {
		values.Add(k, v)
	}
	requestUri := c.requestUri(reqPath)
	debug.DebugPrintPostMapRequest(requestUri, values)
	response, err := c.clt.Post(requestUri, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if resp == nil {
		resp = new(BaseResponse)
	}
	err = debug.DecodeJSONHttpResponse(response.Body, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

// Do execute api request
func (c *Client) Upload(req UploadRequest, accessToken string, resp Response) error {
	reqPath := c.requestPath(req)
	reqParams := req.Params()
	reqParams["_aop_timestamp"] = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	if accessToken != "" {
		reqParams["access_token"] = accessToken
	}
	reqParams["_aop_signature"] = c.Sign(reqPath, reqParams)
	buf := GetBytesBuffer()
	defer PutBytesBuffer(buf)
	mw := multipart.NewWriter(buf)
	for k, v := range reqParams {
		if fw, err := mw.CreateFormField(k); err != nil {
			return err
		} else if _, err := io.Copy(fw, strings.NewReader(v)); err != nil {
			return err
		}
	}
	for k, r := range req.Files() {
		if fw, err := mw.CreateFormFile(k, "file"); err != nil {
			return err
		} else if _, err := io.Copy(fw, r); err != nil {
			return err
		}
	}
	mw.Close()
	requestUri := c.requestUri(reqPath)
	httpReq, err := http.NewRequest("POST", requestUri, buf)
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", mw.FormDataContentType())
	response, err := c.clt.Do(httpReq)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if resp == nil {
		resp = new(BaseResponse)
	}
	err = debug.DecodeJSONHttpResponse(response.Body, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

func (c *Client) requestUri(reqPath string) string {
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	builder.WriteString(GATEWAY)
	builder.WriteString("/")
	builder.WriteString(reqPath)
	return builder.String()
}

func (c *Client) requestPath(req Request) string {
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	builder.WriteString(PROTOCOL)
	builder.WriteString("/")
	builder.WriteString(req.Path())
	builder.WriteString("/")
	builder.WriteString(c.appKey)
	return builder.String()
}

func (c *Client) combine(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	for _, k := range keys {
		builder.WriteString(k)
		builder.WriteString(params[k])
	}
	return builder.String()
}

// Sign sign api request
func (c *Client) Sign(path string, params map[string]string) string {
	builder := GetBytesBuffer()
	defer PutBytesBuffer(builder)
	builder.WriteString(path)
	builder.WriteString(c.combine(params))
	return crypto.HmacSha1(c.secret, builder.Bytes())
}
