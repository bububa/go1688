package go1688

import (
	"fmt"
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

type Client struct {
	http.Client
	appKey string
	secret []byte
}

func NewClient(appKey string, secret string, httpClient *http.Client) *Client {
	client := httpClient
	if httpClient == nil {
		client = http.DefaultClient
	}
	return &Client{
		Client: *client,
		appKey: appKey,
		secret: []byte(secret),
	}
}

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
	requestUri := fmt.Sprintf("%s/%s", GATEWAY, reqPath)
	debug.DebugPrintPostMapRequest(requestUri, values)
	response, err := c.Post(requestUri, "application/x-www-form-urlencoded; charset=UTF-8", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	err = debug.DecodeJSONHttpResponse(response.Body, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}

func (c *Client) requestPath(req Request) string {
	return fmt.Sprintf("%s/%s/%s", PROTOCOL, req.Path(), c.appKey)
}

func (c *Client) combine(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var stringToBeSigned string
	for _, k := range keys {
		stringToBeSigned += k + params[k]
	}
	return stringToBeSigned
}

func (c *Client) Sign(path string, params map[string]string) string {
	raw := fmt.Sprintf("%s%s", path, c.combine(params))
	return crypto.HmacSha1(c.secret, raw)
}
