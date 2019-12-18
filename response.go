package go1688

import (
	"fmt"
)

type Response interface {
	Error() string
	IsError() bool
}

type BaseResponse struct {
	ErrorCode      string `json:"error_code,omitempty"`
	ErrorMessage   string `json:"error_message,omitempty"`
	RequestId      string `json:"request_id,omitempty"`
	Success        bool   `json:"success,omitempty"`
	InnerErrorCode string `json:"errorCode,omitempty"`
	InnerErrorMsg  string `json:"errorMsg,omitempty"`
}

func (this BaseResponse) Error() string {
	if this.InnerErrorCode != "" {
		return fmt.Sprintf("CODE: %s, MSG: %s", this.InnerErrorCode, this.InnerErrorMsg)
	}
	return fmt.Sprintf("CODE: %s, MSG: %s", this.ErrorCode, this.ErrorMessage)
}

func (this *BaseResponse) IsError() bool {
	return this.ErrorCode != "" || this.InnerErrorCode != ""
}
