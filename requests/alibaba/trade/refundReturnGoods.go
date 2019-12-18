package trade

import (
	"encoding/json"
	"errors"

	"github.com/bububa/go1688"
)

type RefundReturnGoodsRequest struct {
	RefundId           string   `json:"refundId,omitempty"`           // 退款单号，TQ开头
	LogisticsCompanyNo string   `json:"logisticsCompanyNo,omitempty"` // 物流公司编码，调用alibaba.logistics.OpQueryLogisticCompanyList.offline接口查询
	FreightBill        string   `json:"freightBill,omitempty"`        // 物流公司运单号，请准确填写，否则卖家有权拒绝退款
	Description        string   `json:"description,omitempty"`        // 发货说明，内容在2-200个字之间
	Vouchers           []string `json:"vouchers,omitempty"`           // 凭证图片URLs，必须使用API alibaba.trade.uploadRefundVoucher返回的“图片域名/相对路径”，最多可上传 10 张图片 ；单张大小不超过1M；支持jpg、gif、jpeg、png、和bmp格式。 请上传凭证，以便以后续赔所需（不上传将无法理赔）
}

type RefundReturnGoodsRefinedRequest struct {
	RefundId           string `json:"refundId,omitempty"`           // 退款单号，TQ开头
	LogisticsCompanyNo string `json:"logisticsCompanyNo,omitempty"` // 物流公司编码，调用alibaba.logistics.OpQueryLogisticCompanyList.offline接口查询
	FreightBill        string `json:"freightBill,omitempty"`        // 物流公司运单号，请准确填写，否则卖家有权拒绝退款
	Description        string `json:"description,omitempty"`        // 发货说明，内容在2-200个字之间
	Vouchers           string `json:"vouchers,omitempty"`           // 凭证图片URLs，必须使用API alibaba.trade.uploadRefundVoucher返回的“图片域名/相对路径”，最多可上传 10 张图片 ；单张大小不超过1M；支持jpg、gif、jpeg、png、和bmp格式。 请上传凭证，以便以后续赔所需（不上传将无法理赔
}

func (this *RefundReturnGoodsRequest) Refine() *RefundReturnGoodsRefinedRequest {
	var vouchers []byte
	if len(this.Vouchers) > 0 {
		vouchers, _ = json.Marshal(this.Vouchers)
	}

	return &RefundReturnGoodsRefinedRequest{
		RefundId:           this.RefundId,
		LogisticsCompanyNo: this.LogisticsCompanyNo,
		FreightBill:        this.FreightBill,
		Description:        this.Description,
		Vouchers:           string(vouchers),
	}
}

func (this *RefundReturnGoodsRefinedRequest) Name() string {
	return "alibaba.trade.refund.returnGoods"
}

type RefundReturnGoodsResponse struct {
	go1688.BaseResponse
	Result *RefundReturnGoodsResult `json:"result,omitempty"`
}

type RefundReturnGoodsResult struct {
	ErrorCode string `json:"errorCode"` // 错误码
	ErrorInfo string `json:"errorInfo"` // 错误信息
	Success   bool   `json:"success"`   // 是否提交成功
}

func RefundReturnGoods(client *go1688.Client, req *RefundReturnGoodsRequest, accessToken string) error {
	refinedReq := req.Refine()
	finalRequest := go1688.NewRequest(NAMESPACE, refinedReq)
	resp := &RefundReturnGoodsResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp
	}
	if !resp.Result.Success {
		return errors.New(resp.Result.ErrorInfo)
	}
	return nil
}
