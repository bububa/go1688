package logistics

import (
	"github.com/bububa/go1688"
)

// OpQueryLogisticCompanyListRequest 物流公司列表-所有的物流公司
type OpQueryLogisticCompanyListRequest struct{}

// Name implement RequestData interface
func (r OpQueryLogisticCompanyListRequest) Name() string {
	return "alibaba.logistics.OpQueryLogisticCompanyList"
}

// Map implement RequestData interface
func (r OpQueryLogisticCompanyListRequest) Map() map[string]string {
	ret := make(map[string]string)
	return ret
}

// OpQueryLogisticCompanyListResponse 物流公司列表-所有的物流公司
type OpQueryLogisticCompanyListResponse struct {
	go1688.BaseResponse
	// Result 物流公司列表
	Result []OpLogisticsCompany `json:"result,omitempty"`
}

// OpLogisticsCompany 物流公司
type OpLogisticsCompany struct {
	// ID 订单编号
	ID uint64 `json:"id,omitempty"`
	// Name 物流公司名称
	Name string `json:"companyName,omitempty"`
	// No 物流公司编号
	No string `json:"companyNo,omitempty"`
	// Phone 物流公司服务电话
	Phone string `json:"companyPhone,omitempty"`
	// SupportPrint 是否支持打印
	SupportPrint bool `json:"supportPrint,omitempty"`
	// Spelling 全拼
	Spelling string `json:"spelling,omitempty"`
}

// GetLogisticCompanyList 物流公司列表-所有的物流公司
func GetLogisticCompanyList(client *go1688.Client, accessToken string) ([]OpLogisticsCompany, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, &OpQueryLogisticCompanyListRequest{})
	var resp OpQueryLogisticCompanyListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
