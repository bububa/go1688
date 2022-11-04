package trade

import (
	"github.com/bububa/go1688"
)

// AddressCodeGetRequest 获取交易地址代码表详情 API Request
type AddressCodeGetRequest struct {
	//  AreaCode 地址code码
	AreaCode string `json:"areaCode,omitempty"`
	// Website 站点信息，指定调用的API是属于国际站（alibaba）还是1688网站（1688）
	Website APIWebsite `json:"webSite,omitempty"`
}

// Name implement RequestData interface
func (r AddressCodeGetRequest) Name() string {
	return "alibaba.trade.addresscode.get"
}

// Map implement RequestData interface
func (r AddressCodeGetRequest) Map() map[string]string {
	return map[string]string{
		"areaCode": r.AreaCode,
		"webSite":  r.Website,
	}
}

// AddressCodeGetResponse 获取交易地址代码表详情 API Response
type AddressCodeGetResponse struct {
	go1688.BaseResponse
	// Result 地区信息
	Result *TradeAddressCode `json:"result,omitempty"`
}

// TradeAddressCode 地区信息
type TradeAddressCode struct {
	// Code 地区编码
	Code string `json:"code,omitempty"`
	// Name 地区名称
	Name string `json:"name,omitempty"`
	// ParentCode 父节点编码，可能为空
	ParentCode string `json:"parentCode,omitempty"`
	// Post 邮编
	Post string `json:"post,omitempty"`
	//  Children 下一级编码列表，可能为空
	Children []string `json:"children,omitempty"`
}

// AddressCodeGet 获取交易地址代码表详情
func AddressCodeGet(client *go1688.Client, req *AddressCodeGetRequest, accessToken string) (*TradeAddressCode, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp AddressCodeGetResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
