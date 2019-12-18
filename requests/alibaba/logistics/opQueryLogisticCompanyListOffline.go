package logistics

import (
	"github.com/bububa/go1688"
)

type OpQueryLogisticCompanyListOfflineRequest struct {
}

func (this *OpQueryLogisticCompanyListOfflineRequest) Name() string {
	return "alibaba.logistics.OpQueryLogisticCompanyList.offline"
}

func GetLogisticCompanyListOffline(client *go1688.Client, accessToken string) ([]*OpLogisticsCompany, error) {
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
