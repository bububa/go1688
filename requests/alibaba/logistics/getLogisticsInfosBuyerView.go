package logistics

import (
	"github.com/bububa/go1688"
)

type GetLogisticsInfosBuyerViewRequest struct {
	OrderId uint64     `json:"orderId"`          // 订单ID
	Fields  string     `json:"fields,omitempty"` // 需要返回的字段，目前有:company.name,sender,receiver,sendgood。返回的字段要用英文逗号分隔开
	Website APIWebsite `json:"webSite"`          // 是1688业务还是icbu业务
}

func (this *GetLogisticsInfosBuyerViewRequest) Name() string {
	return "alibaba.trade.getLogisticsInfos.buyerView"
}

type GetLogisticsInfosBuyerViewResponse struct {
	go1688.BaseResponse
	Result []*LogisticsOrder `json:"result,omitempty"`
}

type LogisticsOrder struct {
	Sender               *LogisticsSender    `json:"sender,omitempty"` // 发件人信息
	ServiceFeature       string              `json:"serviceFeature,omitempty"`
	OrderEntryIds        string              `json:"orderEntryIds,omitempty"`        // 订单号列表，无子订单的等于主订单编号，否则为对应子订单列表
	LogisticsBillNo      string              `json:"logisticsBillNo,omitempty"`      // 物流单号，运单号
	LogisticsId          string              `json:"logisticsId,omitempty"`          // 物流信息ID
	Receiver             *LogisticsReceiver  `json:"receiver,omitempty"`             // 收件人信息
	LogisticsCompanyId   string              `json:"logisticsCompanyId,omitempty"`   // 物流公司ID
	LogisticsCompanyName string              `json:"logisticsCompanyName,omitempty"` // 物流公司编码
	Status               LogisticsStatus     `json:"status,omitempty"`               // 物流状态。WAITACCEPT:未受理;CANCEL:已撤销;ACCEPT:已受理;TRANSPORT:运输中;NOGET:揽件失败;SIGN:已签收;UNSIGN:签收异常
	SendGoods            []LogisticsSendGood `json:"sendGoods,omitempty"`            // 商品信息
	GmtSystemSend        string              `json:"gmtSystemSend,omitempty"`
	Remarks              string              `json:"remarks,omitempty"` // 备注
}

type LogisticsSender struct {
	Name         string `json:"senderName,omitempty"`
	Phone        string `json:"senderPhone,omitempty"`
	Mobile       string `json:"senderMobile,omitempty"`
	Encrypt      string `json:"encrypt,omitempty"`
	ProvinceCode string `json:"senderProvinceCode,omitempty"`
	CityCode     string `json:"senderCityCode,omitempty"`
	CountyCode   string `json:"senderCountryCode,omitempty"`
	Address      string `json:"senderAddress,omitempty"`
	Province     string `json:"senderProvince,omitempty"`
	City         string `json:"senderCity,omitempty"`
	County       string `json:"senderCounty,omitempty"`
}

type LogisticsReceiver struct {
	Name         string `json:"receiverName,omitempty"`
	Phone        string `json:"receiverPhone,omitempty"`
	Mobile       string `json:"receiverMobile,omitempty"`
	Encrypt      string `json:"encrypt,omitempty"`
	ProvinceCode string `json:"receiverProvinceCode,omitempty"`
	CityCode     string `json:"receiverCityCode,omitempty"`
	CountyCode   string `json:"receiverCountryCode,omitempty"`
	Address      string `json:"receiverAddress,omitempty"`
	Province     string `json:"receiverProvince,omitempty"`
	City         string `json:"receiverCity,omitempty"`
	County       string `json:"receiverCounty,omitempty"`
}

type LogisticsSendGood struct {
	Name     string `json:"goodName,omitempty"` // 商品名
	Quantity string `json:"quantity,omitempty"` // 商品数量
	Unit     string `json:"unit,omitempty"`     // 商品单位
}

func GetLogisticsInfosBuyerView(client *go1688.Client, req *GetLogisticsInfosBuyerViewRequest, accessToken string) ([]*LogisticsOrder, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &GetLogisticsInfosBuyerViewResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Result, nil
}
