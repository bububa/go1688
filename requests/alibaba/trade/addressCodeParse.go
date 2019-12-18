package trade

import (
	"github.com/bububa/go1688"
)

type AddressCodeParseRequest struct {
	AddressInfo string `json:"addressInfo"` // 地址信息
}

func (this *AddressCodeParseRequest) Name() string {
	return "alibaba.trade.addresscode.parse"
}

type AddressCodeParseResponse struct {
	go1688.BaseResponse
	Result *ReceiveAddress `json:"result,omitempty"`
}

type ReceiveAddress struct {
	Address         string `json:"address,omitempty"`         // 街道地址，不包括省市编码
	AddressCode     string `json:"addressCode,omitempty"`     // 地址区域编码
	AddressCodeText string `json:"addressCodeText,omitempty"` // 地址区域编码对应的文本（包括国家，省，城市）
	AddressId       uint64 `json:"addredssId,omitempty"`
	BizType         string `json:"bizType,omitempty"`   // 记录收货地址的业务类型
	IsDefault       bool   `json:"isDefault,omitempty"` // 是否为默认
	FullName        string `json:"fullName,omitempty"`  // 收货人姓名
	Latest          bool   `json:"latest,omitempty"`    // 是否是最后选择的收货地址
	Mobile          string `json:"mobile,omitempty"`    // 手机号
	Phone           string `json:"phone,omitempty"`     // 电话
	PostCode        string `json:"postCode,omitempty"`  // 邮编
}

type LogisticsStep struct {
	AcceptTime string `json:"acceptTime,omitempty"` // 物流跟踪单该步骤的时间
	Remark     string `json:"remark,omitempty"`     // 备注，如：“在浙江浦江县公司进行下级地点扫描，即将发往：广东深圳公司”
}

func AddressCodeParse(client *go1688.Client, req *AddressCodeParseRequest, accessToken string) (*ReceiveAddress, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &AddressCodeParseResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Result, nil
}
