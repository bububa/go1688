package trade

import (
	"github.com/bububa/go1688"
)

// RefundOpQueryOrderRefundRequest 查询退款单详情-根据退款单ID（买家视角）
type RefundOpQueryOrderRefundRequest struct {
	// RefundID 退款单业务主键 TQ+ID
	RefundID string `json:"refundId,omitempty"`
	// NeedTimeOutInfo 需要退款单的超时信息
	NeedTimeOutInfo bool `json:"needTimeOutInfo,omitempty"`
	// NeedOrderRefundOperation 需要退款单伴随的所有退款操作信息
	NeedOrderRefundOperation bool `json:"needOrderRefundOperation,omitempty"`
}

// Name implement RequestData interface
func (r RefundOpQueryOrderRefundRequest) Name() string {
	return "alibaba.trade.refund.OpQueryOrderRefund-1"
}

// Map implement RequestData interface
func (r RefundOpQueryOrderRefundRequest) Map() map[string]string {
	ret := make(map[string]string, 3)
	ret["refundId"] = r.RefundID
	if r.NeedTimeOutInfo {
		ret["needTimeOutInfo"] = "true"
	}
	if r.NeedOrderRefundOperation {
		ret["needOrderRefundOperation"] = "true"
	}
	return ret
}

// RefundOpQueryOrderRefundResponse 查询退款单详情-根据退款单ID（买家视角）
type RefundOpQueryOrderRefundResponse struct {
	go1688.BaseResponse
	// Result 查询结果
	Result *RefundOpQueryOrderRefundResult `json:"result,omitempty"`
}

type RefundOpQueryOrderRefundResult struct {
	// List 退款模型
	List []OpOrderRefundModel `json:"opOrderRefundModelDetail,omitempty"`
}

// RefundOpQueryOrderRefund 查询退款单详情-根据退款单ID（买家视角）
func RefundOpQueryOrderRefund(client *go1688.Client, req *RefundOpQueryOrderRefundRequest, accessToken string) ([]OpOrderRefundModel, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp RefundOpQueryOrderRefundResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Result.List, nil
}
