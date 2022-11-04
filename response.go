package go1688

type Response interface {
	Error() string
	IsError() bool
}

type BaseResponse struct {
	ErrorCode       string `json:"error_code,omitempty"`
	ErrorMessage    string `json:"error_message,omitempty"`
	RequestId       string `json:"request_id,omitempty"`
	Success         bool   `json:"success,omitempty"`
	InnerErrorCode  string `json:"errorCode,omitempty"`
	InnerErrorMsg   string `json:"errorMsg,omitempty"`
	ExtErrorMessage string `json:"extErrorMessage,omitempty"`
}

func (r BaseResponse) Error() string {
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	builder.WriteString("CODE: ")
	if r.InnerErrorCode != "" {
		builder.WriteString(r.InnerErrorCode)
		builder.WriteString(", MSG: ")
		builder.WriteString(r.InnerErrorMsg)
	} else {
		builder.WriteString(r.ErrorCode)
		builder.WriteString(", MSG: ")
		builder.WriteString(r.ErrorMessage)
	}
	if r.ExtErrorMessage != "" {
		builder.WriteString(", EXT: ")
		builder.WriteString(r.ExtErrorMessage)
	}
	return builder.String()
}

func (r BaseResponse) IsError() bool {
	return r.ErrorCode != "" || r.InnerErrorCode != ""
}
