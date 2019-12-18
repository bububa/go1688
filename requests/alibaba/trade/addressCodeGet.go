package trade

import (
	"github.com/bububa/go1688"
)

type AddressCodeGetRequest struct {
	AreaCode string     `json:"areaCode"` // 地址code码
	Website  APIWebsite `json:"webSite"`  // 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
}

func (this *AddressCodeGetRequest) Name() string {
	return "alibaba.trade.addresscode.get"
}

type AddressCodeGetResponse struct {
	go1688.BaseResponse
	Result *TradeAddressCode `json:"result,omitempty"`
}

type TradeAddressCode struct {
	Code       string   `json:"code,omitempty"`       // 地区编码
	Name       string   `json:"name,omitempty"`       // 地区名称
	ParentCode string   `json:"parentCode,omitempty"` // 父节点编码，可能为空
	Post       string   `json:"post,omitempty"`       // 邮编
	Children   []string `json:"children,omitempty"`   // 下一级编码列表，可能为空
}

func AddressCodeGet(client *go1688.Client, req *AddressCodeGetRequest, accessToken string) (*TradeAddressCode, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &AddressCodeGetResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Result, nil
}
