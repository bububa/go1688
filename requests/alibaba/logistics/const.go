package logistics

const (
	NAMESPACE = "com.alibaba.logistics"
)

type APIWebsite = string

const (
	API_ALIBABA APIWebsite = "alibaba"
	API_1688    APIWebsite = "1688"
)

type LogisticsStatus = string

const (
	WAIT_ACCEPT_STATUS LogisticsStatus = "WAITACCEPT" // 未受理
	CANCEL_STATUS      LogisticsStatus = "CANCEL"     // 已撤销
	ACCEPT_STATUS      LogisticsStatus = "ACCEPT"     // 已受理
	TRANSPORT_STATUS   LogisticsStatus = "TRANSPORT"  // 运输中
	NOGET_STATUS       LogisticsStatus = "NOGET"      // 揽件失败
	SIGN_STATUS        LogisticsStatus = "SIGN"       // 已签收
	UNSIGN_STATUS      LogisticsStatus = "UNSIGN"     // 签收异常
)
