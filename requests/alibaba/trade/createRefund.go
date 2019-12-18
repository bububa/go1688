package trade

import (
	"encoding/json"
	"errors"

	"github.com/bububa/go1688"
)

type CreateRefundRequest struct {
	OrderId             uint64            `json:"orderId,omitempty"`             // 主订单
	OrderEntryIds       []uint64          `json:"orderEntryIds,omitempty"`       // 子订单
	DisputeRequest      DisputeRequest    `json:"disputeRequest,omitempty"`      // 退款/退款退货。只有已收到货，才可以选择退款退货。退款:"refund"; 退款退货:"returnRefund"
	ApplyPayment        uint              `json:"applyPayment,omitempty"`        // 退款金额（单位：分）。不大于实际付款金额；等待卖家发货时，必须为商品的实际付款金额。
	ApplyCarriage       uint              `json:"applyCarriage,omitempty"`       // 退运费金额（单位：分）。
	ApplyReasonId       uint64            `json:"applyReasonId,omitempty"`       // 退款原因id（从API getRefundReasonList获取）
	Description         string            `json:"description,omitempty"`         // 退款申请理由，2-150字
	GoodsStatus         RefundGoodsStatus `json:"goodsStatus,omitempty"`         // 货物状态, 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived"
	Vouchers            []string          `json:"vouchers,omitempty"`            // 凭证图片URLs。1-5张，必须使用API uploadRefundVoucher返回的“图片域名/相对路径”
	OrderEntryCountList []OrderEntryCount `json:"orderEntryCountList,omitempty"` // 子订单退款数量。仅在售中买家已收货（退款退货）时，可指定退货数量；默认，全部退货。
}

type OrderEntryCount struct {
	Id    uint64 `json:"id,omitempty"`    // 子订单id
	Count uint   `json:"count,omitempty"` // 子订单购买商品数量
}

type CreateRefundRefinedRequest struct {
	OrderId             uint64            `json:"orderId,omitempty"`             // 主订单
	OrderEntryIds       string            `json:"orderEntryIds,omitempty"`       // 子订单
	DisputeRequest      DisputeRequest    `json:"disputeRequest,omitempty"`      // 退款/退款退货。只有已收到货，才可以选择退款退货。退款:"refund"; 退款退货:"returnRefund"
	ApplyPayment        uint              `json:"applyPayment,omitempty"`        // 退款金额（单位：分）。不大于实际付款金额；等待卖家发货时，必须为商品的实际付款金额。
	ApplyCarriage       uint              `json:"applyCarriage,omitempty"`       // 退运费金额（单位：分）。
	ApplyReasonId       uint64            `json:"applyReasonId,omitempty"`       // 退款原因id（从API getRefundReasonList获取）
	Description         string            `json:"description,omitempty"`         // 退款申请理由，2-150字
	GoodsStatus         RefundGoodsStatus `json:"goodsStatus,omitempty"`         // 货物状态, 售中等待卖家发货:"refundWaitSellerSend"; 售中等待买家收货:"refundWaitBuyerReceive"; 售中已收货（未确认完成交易）:"refundBuyerReceived" 售后未收货:"aftersaleBuyerNotReceived"; 售后已收到货:"aftersaleBuyerReceived"
	Vouchers            string            `json:"vouchers,omitempty"`            // 凭证图片URLs。1-5张，必须使用API uploadRefundVoucher返回的“图片域名/相对路径”
	OrderEntryCountList string            `json:"orderEntryCountList,omitempty"` // 子订单退款数量。仅在售中买家已收货（退款退货）时，可指定退货数量；默认，全部退货。
}

func (this *CreateRefundRequest) Refine() *CreateRefundRefinedRequest {
	var (
		entryIds       []byte
		vouchers       []byte
		entryCountList []byte
	)
	if len(this.OrderEntryIds) > 0 {
		entryIds, _ = json.Marshal(this.OrderEntryIds)
	}
	if len(this.Vouchers) > 0 {
		vouchers, _ = json.Marshal(this.Vouchers)
	}
	if len(this.OrderEntryCountList) > 0 {
		entryCountList, _ = json.Marshal(this.OrderEntryCountList)
	}

	return &CreateRefundRefinedRequest{
		OrderId:             this.OrderId,
		OrderEntryIds:       string(entryIds),
		DisputeRequest:      this.DisputeRequest,
		ApplyPayment:        this.ApplyPayment,
		ApplyCarriage:       this.ApplyCarriage,
		ApplyReasonId:       this.ApplyReasonId,
		Description:         this.Description,
		GoodsStatus:         this.GoodsStatus,
		Vouchers:            string(vouchers),
		OrderEntryCountList: string(entryCountList),
	}
}

func (this *CreateRefundRefinedRequest) Name() string {
	return "alibaba.trade.createRefund"
}

type CreateRefundResponse struct {
	go1688.BaseResponse
	Result *CreateRefundResult `json:"result,omitempty"`
}

type CreateRefundResult struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
	Result  struct {
		RefundId string `json:"refundId,omitempty"` // 创建成功，退款id
	} `json:"result,omitempty"` // 结果
}

func CreateRefund(client *go1688.Client, req *CreateRefundRequest, accessToken string) (string, error) {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &CreateRefundResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", resp
	}
	if resp.Result.Code != "" {
		return "", errors.New(resp.Result.Message)
	}
	return resp.Result.Result.RefundId, nil
}
