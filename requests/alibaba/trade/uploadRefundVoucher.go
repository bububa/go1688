package trade

import (
	"io"

	"github.com/bububa/go1688"
)

// UploadRefundVoucherRequest 上传退款退货凭证 API Request
type UploadRefundVoucherRequest struct {
	// ImageData 凭证图片数据。小于1M，jpg格式。
	ImageData io.Reader `json:"imageData,omitempty"`
}

// Name implement RequestData interface
func (r UploadRefundVoucherRequest) Name() string {
	return "alibaba.trade.uploadRefundVoucher"
}

// Map implement UploadRequestData interface
func (r UploadRefundVoucherRequest) Map() map[string]string {
	return make(map[string]string)
}

// Files implement UploadRequestData interface
func (r UploadRefundVoucherRequest) Files() map[string]io.Reader {
	return map[string]io.Reader{
		"imageData": r.ImageData,
	}
}

// UploadRefundVoucherResponse 上传退款退货凭证 API Response
type UploadRefundVoucherResponse struct {
	go1688.BaseResponse
	Result *UploadRefundVoucherResult `json:"result,omitempty"`
}

// UploadRefundVoucherResult 返回结果
type UploadRefundVoucherResult struct {
	// Code 错误码
	Code string `json:"code"`
	// Message 错误信息
	Message string `json:"message"`
	// Success 是否提交成功
	Success bool `json:"success"`
	// Result 上传退款退货凭证结果
	Result *OrderRefundUploadVoucherResult `json:"result,omitempty"`
}

// IsError check success
func (r UploadRefundVoucherResult) IsError() bool {
	return !r.Success
}

func (r UploadRefundVoucherResult) Error() string {
	builder := go1688.GetStringsBuilder()
	defer go1688.PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	builder.WriteString(r.Code)
	builder.WriteString(", MSG: ")
	builder.WriteString(r.Message)
	return builder.String()
}

// OrderRefundUploadVoucherResult 上传退款退货凭证结果
type OrderRefundUploadVoucherResult struct {
	// ImageDomain 图片域名
	ImageDomain string `json:"imageDomain,omitempty"`
	// ImageRelativeURL 图片相对路径
	ImageRelativeURL string `json:"imageRelativeURL,omitempty"`
}

// UploadRefundVoucher 上传退款退货凭证
func UploadRefundVoucher(client *go1688.Client, req *UploadRefundVoucherRequest, accessToken string) (*OrderRefundUploadVoucherResult, error) {
	finalRequest := go1688.NewUploadRequest(NAMESPACE, req)
	var resp UploadRefundVoucherResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	if resp.Result.IsError() {
		return nil, resp.Result
	}
	return resp.Result.Result, nil
}
