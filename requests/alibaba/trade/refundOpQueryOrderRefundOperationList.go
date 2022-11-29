package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// RefundOpQueryOrderRefundOperationListRequest 退款单操作记录列表（买家视角）
type RefundOpQueryOrderRefundOperationListRequest struct {
	// RefundID 退款单业务主键 TQ+ID
	RefundID string `json:"refundId,omitempty"`
	// PageNo 当前页号
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 页大小
	PageSize int `json:"pageSize,omitempty"`
}

// Name implement RequestData interface
func (r RefundOpQueryOrderRefundOperationListRequest) Name() string {
	return "alibaba.trade.refund.OpQueryOrderRefundOperationList-1"
}

// Map implement RequestData interface
func (r RefundOpQueryOrderRefundOperationListRequest) Map() map[string]string {
	ret := make(map[string]string, 3)
	ret["refundId"] = r.RefundID
	ret["pageNo"] = strconv.Itoa(r.PageNo)
	ret["pageSize"] = strconv.Itoa(r.PageSize)
	return ret
}

// RefundOpQueryOrderRefundOperationListResponse 退款单操作记录列表（买家视角）
type RefundOpQueryOrderRefundOperationListResponse struct {
	go1688.BaseResponse
	// Result 查询结果
	Result *RefundOpQueryOrderRefundOperationListResult `json:"result,omitempty"`
}

type RefundOpQueryOrderRefundOperationListResult struct {
	// List 退款模型
	List []OpOrderRefundOperationModal `json:"opOrderRefundOperationModels,omitempty"`
}

// RefundOpQueryOrderRefundOperationList 退款单操作记录列表（买家视角）
func RefundOpQueryOrderRefundOperationList(client *go1688.Client, req *RefundOpQueryOrderRefundRequest, accessToken string) ([]OpOrderRefundOperationModal, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp RefundOpQueryOrderRefundOperationListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result.List, nil
}
