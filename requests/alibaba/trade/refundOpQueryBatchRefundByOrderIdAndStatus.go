package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// RefundOpQueryBatchRefundByOrderIdAndStatusRequest 查询退款单详情-根据订单ID（买家视角）
type RefundOpQueryBatchRefundByOrderIdAndStatusRequest struct {
	// OrderID 订单id
	OrderID uint64 `json:"orderId,omitempty"`
	// QueryType 1：活动；3:退款成功（只支持退款中和退款成功）
	QueryType int `json:"queryType,omitempty"`
}

// Name implement RequestData interface
func (r RefundOpQueryBatchRefundByOrderIdAndStatusRequest) Name() string {
	return "alibaba.trade.refund.OpQueryBatchRefundByOrderIdAndStatus-1"
}

// Map implement RequestData interface
func (r RefundOpQueryBatchRefundByOrderIdAndStatusRequest) Map() map[string]string {
	ret := make(map[string]string, 2)
	if r.OrderID > 0 {
		ret["orderId"] = strconv.FormatUint(r.OrderID, 10)
	}
	if r.QueryType != 0 {
		ret["queryType"] = strconv.Itoa(r.QueryType)
	}
	return ret
}

// RefundOpQueryBatchRefundByOrderIdAndStatus 查询退款单详情-根据订单ID（买家视角）
func RefundOpQueryBatchRefundByOrderIdAndStatus(client *go1688.Client, req *RefundOpQueryBatchRefundByOrderIdAndStatusRequest, accessToken string) ([]OpOrderRefundModel, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp RefundBuyerQueryOrderRefundListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result.List, nil
}
