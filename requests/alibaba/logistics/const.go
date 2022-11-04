package logistics

const (
	NAMESPACE = "com.alibaba.logistics"
)

// APIWebsite 是1688业务还是icbu业务
type APIWebsite = string

const (
	// API_ALIBABA alibaba
	API_ALIBABA APIWebsite = "alibaba"
	// API_1688 1688
	API_1688 APIWebsite = "1688"
)

// LogisticsStatus 物流状态
type LogisticsStatus = string

const (
	// WAIT_ACCEPT_STATUS 未受理
	WAIT_ACCEPT_STATUS LogisticsStatus = "WAITACCEPT"
	// CANCEL_STATUS 已撤销
	CANCEL_STATUS LogisticsStatus = "CANCEL"
	// ACCEPT_STATUS 已受理
	ACCEPT_STATUS LogisticsStatus = "ACCEPT"
	// TRANSPORT_STATUS 运输中
	TRANSPORT_STATUS LogisticsStatus = "TRANSPORT"
	// NOGET_STATUS 揽件失败
	NOGET_STATUS LogisticsStatus = "NOGET"
	// SIGN_STATUS 已签收
	SIGN_STATUS LogisticsStatus = "SIGN"
	// UNSIGN_STATUS 签收异常
	UNSIGN_STATUS LogisticsStatus = "UNSIGN"
)
