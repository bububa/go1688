package trade

import (
	"github.com/bububa/go1688"
)

// AddressCodeParseRequest 根据地址解析地区码 API Request
type AddressCodeParseRequest struct {
	// AddressInfo 地址信息
	AddressInfo string `json:"addressInfo"`
}

// Name implement RequestData interface
func (r *AddressCodeParseRequest) Name() string {
	return "alibaba.trade.addresscode.parse"
}

// Map implement RequestData interface
func (r *AddressCodeParseRequest) Map() map[string]string {
	return map[string]string{
		"addressInfo": r.AddressInfo,
	}
}

// AddressCodeParseResponse 根据地址解析地区码 API Response
type AddressCodeParseResponse struct {
	go1688.BaseResponse
	// Result 获地址
	Result *ReceiveAddress `json:"result,omitempty"`
}

// ReceiveAddress 收获地址
type ReceiveAddress struct {
	// Address 街道地址，不包括省市编码
	Address string `json:"address,omitempty"`
	// AddressCode 地址区域编码
	AddressCode string `json:"addressCode,omitempty"`
	// AddressCodeText 地址区域编码对应的文本（包括国家，省，城市）
	AddressCodeText string `json:"addressCodeText,omitempty"`
	// AddressID .
	AddressID uint64 `json:"addredssId,omitempty"`
	// BizType 记录收货地址的业务类型
	BizType string `json:"bizType,omitempty"`
	// IsDefault 是否为默认
	IsDefault bool `json:"isDefault,omitempty"`
	// FullName 收货人姓名
	FullName string `json:"fullName,omitempty"`
	// Latest 是否是最后选择的收货地址
	Latest bool `json:"latest,omitempty"`
	// Mobile 手机号
	Mobile string `json:"mobile,omitempty"`
	// Phone 电话
	Phone string `json:"phone,omitempty"`
	// PostCode 邮编
	PostCode string `json:"postCode,omitempty"`
}

// AddressCodeParse 根据地址解析地区码
func AddressCodeParse(client *go1688.Client, req *AddressCodeParseRequest, accessToken string) (*ReceiveAddress, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp AddressCodeParseResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
