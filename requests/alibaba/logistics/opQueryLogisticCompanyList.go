package logistics

import (
	"github.com/bububa/go1688"
)

type OpQueryLogisticCompanyListRequest struct {
}

func (this *OpQueryLogisticCompanyListRequest) Name() string {
	return "alibaba.logistics.OpQueryLogisticCompanyList"
}

type OpQueryLogisticCompanyListResponse struct {
	go1688.BaseResponse
	Result          []*OpLogisticsCompany `json:"result,omitempty"`
	ExtErrorMessage string                `json:"extErrorMessage,omitempty"`
}

type OpLogisticsCompany struct {
	Id           uint64 `json:"id,omitempty"`           // 订单编号
	Name         string `json:"companyName,omitempty"`  // 物流公司名称
	No           string `json:"companyNo,omitempty"`    // 物流公司编号
	Phone        string `json:"companyPhone,omitempty"` // 物流公司服务电话
	SupportPrint bool   `json:"supportPrint,omitempty"` // 是否支持打印
	Spelling     string `json:"spelling,omitempty"`     // 全拼
}

func GetLogisticCompanyList(client *go1688.Client, accessToken string) ([]*OpLogisticsCompany, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, &OpQueryLogisticCompanyListRequest{})
	resp := &OpQueryLogisticCompanyListResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Result, nil
}
