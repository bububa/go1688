package trade

import (
	"github.com/bububa/go1688"
)

// RefundReturnGoodsRequest 买家提交退款货信息 API Request
type RefundReturnGoodsRequest struct {
	// RefundID 退款单号，TQ开头
	RefundID string `json:"refundId,omitempty"`
	// LogisticsCompanyNo 物流公司编码，调用alibaba.logistics.OpQueryLogisticCompanyList.offline接口查询
	LogisticsCompanyNo string `json:"logisticsCompanyNo,omitempty"`
	// FreightBill 物流公司运单号，请准确填写，否则卖家有权拒绝退款
	FreightBill string `json:"freightBill,omitempty"`
	// Description 发货说明，内容在2-200个字之间
	Description string `json:"description,omitempty"`
	// Vouchers 凭证图片URLs，必须使用API alibaba.trade.uploadRefundVoucher返回的“图片域名/相对路径”，最多可上传 10 张图片 ；单张大小不超过1M；支持jpg、gif、jpeg、png、和bmp格式。 请上传凭证，以便以后续赔所需（不上传将无法理赔）
	Vouchers []string `json:"vouchers,omitempty"`
}

// Name implement RequestData interface
func (r RefundReturnGoodsRequest) Name() string {
	return "alibaba.trade.refund.returnGoods"
}

// Map implement RequestData interface
func (r RefundReturnGoodsRequest) Map() map[string]string {
	ret := make(map[string]string, 5)
	ret["refundId"] = r.RefundID
	ret["logisticsCompanyNo"] = r.LogisticsCompanyNo
	ret["reightBill"] = r.FreightBill
	ret["description"] = r.Description
	if len(r.Vouchers) > 0 {
		ret["vouchers"] = go1688.JSONMarshal(r.Vouchers)
	}
	return ret
}

// RefundReturnGoodsResponse 买家提交退款货信息 API Response
type RefundReturnGoodsResponse struct {
	go1688.BaseResponse
	Result *RefundReturnGoodsResult `json:"result,omitempty"`
}

// RefundReturnGoodsResult 买家提交退款货信息 API Result
type RefundReturnGoodsResult struct {
	// ErrorCode 错误码
	ErrorCode string `json:"errorCode"`
	// ErrorInfo 错误信息
	ErrorInfo string `json:"errorInfo"`
	// Success 是否提交成功
	Success bool `json:"success"`
}

// IsError check success
func (r RefundReturnGoodsResult) IsError() bool {
	return !r.Success
}

func (r RefundReturnGoodsResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.ErrorCode)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.ErrorInfo)
	return builder.String()
}

// RefundReturnGoods 买家提交退款货信息
func RefundReturnGoods(client *go1688.Client, req *RefundReturnGoodsRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp RefundReturnGoodsResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return err
	}
	if resp.Result.IsError() {
		return resp.Result
	}
	return nil
}
