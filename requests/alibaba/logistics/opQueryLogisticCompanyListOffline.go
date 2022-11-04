package logistics

import (
	"github.com/bububa/go1688"
)

// OpQueryLogisticCompanyListOfflineRequest 物流公司列表-自联物流 API Request
type OpQueryLogisticCompanyListOfflineRequest struct{}

// Name implement RequestData interface
func (r OpQueryLogisticCompanyListOfflineRequest) Name() string {
	return "alibaba.logistics.OpQueryLogisticCompanyList.offline"
}

// Map implement RequestData interface
func (r OpQueryLogisticCompanyListOfflineRequest) Map() map[string]string {
	ret := make(map[string]string)
	return ret
}

// GetLogisticCompanyListOffline 物流公司列表-自联物流
func GetLogisticCompanyListOffline(client *go1688.Client, accessToken string) ([]OpLogisticsCompany, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, &OpQueryLogisticCompanyListOfflineRequest{})
	var resp OpQueryLogisticCompanyListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
