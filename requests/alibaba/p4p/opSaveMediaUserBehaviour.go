package p4p

import (
	"strconv"

	"github.com/bububa/go1688"
)

// OpSaveMediaUserBehaviourRequest 机构下用户行为日志上报 API Request
type OpSaveMediaUserBehaviourRequest struct {
	// Uuid 代表唯一一条日志记录
	Uuid string `json:"uuid,omitempty"`
	// DeviceIDMd5 设备id求md5(32位小写)(手机号与设备号至少一个)
	DeviceIDMd5 string `json:"deviceIdMd5,omitempty"`
	// PhoneMd5 手机号求md5(32位小写)(手机号与设备号至少一个)
	PhoneMd5 string `json:"phoneMd5,omitempty"`
	// ActionTime 行为时间，13位时间戳精确到毫秒
	ActionTime int64 `json:"actionTime,omitempty"`
	// ActionType 行为类型，1688定义枚举值： CLICK(点击商品详情),COMMENT(评论),STORE(收藏),CART(加购物车),SEARCH(关键词搜索，配合actionDetail),曝光(VIEW)
	ActionType ActionType `json:"actionType,omitempty"`
	// ActionDetail 行为类型补充字段，部分类型配合使用，如搜索内容
	ActionDetail string `json:"actionDetail,omitempty"`
	// FeedType 商品类型,1:1688商品;2 机构商品;
	FeedType FeedType `json:"feedType,omitempty"`
	// FeedID 商品id,1688商品必须传商品id
	FeedID uint64 `json:"feedId,omitempty"`
	// FeeddTitle 商品标题
	FeedTitle string `json:"feedTitle,omitempty"`
	// FeedPrice 售卖价格，单位元，两位小数
	FeedPrice float64 `json:"feedPrice,omitempty"`
	// FeedCategory 商品类目，多级类目英文分号分割，按一级类目;二级类目;叶子类目格式传
	FeedCategory string `json:"feedCategory,omitempty"`
}

// Name implement RequestData interface
func (r OpSaveMediaUserBehaviourRequest) Name() string {
	return "alibaba.cps.op.saveMediaUserBehaviour"
}

// Map implement RequestData interface
func (r OpSaveMediaUserBehaviourRequest) Map() map[string]string {
	ret := make(map[string]string, 11)
	ret["uuid"] = r.Uuid
	if r.DeviceIDMd5 != "" {
		ret["deviceIdMd5"] = r.DeviceIDMd5
	}
	if r.PhoneMd5 != "" {
		ret["phoneMd5"] = r.PhoneMd5
	}
	ret["actionTime"] = strconv.FormatInt(r.ActionTime, 10)
	ret["actionType"] = r.ActionType
	if r.ActionDetail != "" {
		ret["actionDetail"] = r.ActionDetail
	}
	ret["feedType"] = strconv.Itoa(r.FeedType)
	if r.FeedID > 0 {
		ret["feedId"] = strconv.FormatUint(r.FeedID, 10)
	}
	if r.FeedTitle != "" {
		ret["feedTitle"] = r.FeedTitle
	}
	if r.FeedPrice > 1e-15 {
		ret["feedPrice"] = strconv.FormatFloat(r.FeedPrice, 'f', 2, 64)
	}
	if r.FeedCategory != "" {
		ret["feedCategory"] = r.FeedCategory
	}
	return ret
}

// OpSaveMediaUserBehaviour 机构下用户行为日志上报
func OpSaveMediaUserBehaviour(client *go1688.Client, req *OpSaveMediaUserBehaviourRequest, accessToken string) error {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	return client.Do(finalRequest, accessToken, nil)
}
