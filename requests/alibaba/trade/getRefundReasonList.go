package trade

import (
	"strconv"

	"github.com/bububa/go1688"
)

// GetRefundReasonListRequest 查询退款退货原因（用于创建退款退货） API Request
type GetRefundReasonListRequest struct {
	// OrderID 主订单
	OrderID uint64 `json:"orderId,omitempty"`
	// OrderEntryIDs 子订单
	OrderEntryIDs []uint64 `json:"orderEntryIds,omitempty"`
	// GoodsStatus 货物状态, 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived"
	GoodsStatus RefundGoodsStatus `json:"goodsStatus,omitempty"`
}

// Name implement RequestData interface
func (r GetRefundReasonListRequest) Name() string {
	return "alibaba.trade.getRefundReasonList"
}

// Map implement RequestData interface
func (r GetRefundReasonListRequest) Map() map[string]string {
	return map[string]string{
		"orderId":       strconv.FormatUint(r.OrderID, 10),
		"orderEntryIDs": go1688.JSONMarshal(r.OrderEntryIDs),
		"goodsStatus":   r.GoodsStatus,
	}
}

// GetRefundReasonListResponse 查询退款退货原因（用于创建退款退货） API Response
type GetRefundReasonListResponse struct {
	go1688.BaseResponse
	Result *GetRefundReasonListResult `json:"result,omitempty"`
}

// GetRefundReasonListResult 查询退款退货原因（用于创建退款退货） API Result
type GetRefundReasonListResult struct {
	// Code 错误码
	Code string `json:"code"`
	// Message 错误信息
	Message string `json:"message"`
	// Result 结果
	Result struct {
		// Reasons 原因列表
		Reasons []OrderRefundReason `json:"reasons,omitempty"`
	} `json:"result,omitempty"`
	// Success 是否成功
	Success bool `json:"success,omitempty"`
}

// IsError check success
func (r GetRefundReasonListResult) IsError() bool {
	return !r.Success
}

// Error implement error interface
func (r GetRefundReasonListResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.Code)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.Message)
	return builder.String()
}

// OrderRefundReason 退款原因
type OrderRefundReason struct {
	// ID 原因id
	ID uint64 `json:"id,omitempty"`
	// Name 原因
	Name string `json:"name,omitempty"`
	// NeedVoucher 凭证是否必须上传
	NeedVoucher bool `json:"needVoucher,omitempty"`
	// NoRefundCarriage 是否支持退运费
	NoRefundCarriage bool `json:"noRefundCarriage,omitempty"`
	// Tip 提示
	Tip string `json:"tip,omitempty"`
}

// GetRefundReasonList 查询退款退货原因（用于创建退款退货）
func GetRefundReasonList(client *go1688.Client, req *GetRefundReasonListRequest, accessToken string) ([]OrderRefundReason, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp GetRefundReasonListResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	if resp.Result.IsError() {
		return nil, resp.Result
	}
	return resp.Result.Result.Reasons, nil
}
