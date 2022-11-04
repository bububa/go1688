package product

import (
	"strconv"

	"github.com/bububa/go1688"
)

// CategoryGetRequest 根据类目Id查询类目 API Request
type CategoryGetRequest struct {
	// ID 类目id,必须大于等于0， 如果为0，则查询所有一级类目
	ID uint64 `json:"categoryID,omitempty"`
}

// Name implement RequestData interface
func (r CategoryGetRequest) Name() string {
	return "alibaba.category.get"
}

// Map implement RequestData interface
func (r CategoryGetRequest) Map() map[string]string {
	ret := make(map[string]string)
	ret["categoryID"] = strconv.FormatUint(r.ID, 10)
	return ret
}

// CategoryGetResponse 根据类目Id查询类目 API Response
type CategoryGetResponse struct {
	go1688.BaseResponse
	// Category 类目
	Category *CategoryInfo `json:"categoryInfo,omitempty"`
}

// CategoryInfo 类目
type CategoryInfo struct {
	// ID 类目ID
	ID uint64 `json:"categoryID,omitempty"`
	// Name 类目名称
	Name string `json:"name,omitempty"`
	// IsLeaf 是否叶子类目（只有叶子类目才能发布商品）
	IsLeaf bool `json:"isLeaf,omitempty"`
	// ParentIDs 父类目ID数组,1688只返回一个父id
	ParentIDs []uint64 `json:"parentIDs,omitempty"`
	// MinOrderQuality 最小起订量
	MinOrderQuantity uint64 `json:"minOrderQuantity,omitempty"`
	// ChildCategories 子类目信息
	ChildCategories []ChildCategoryInfo `json:"childCategorys,omitempty"`
	// Features 类目特征信息
	Features []CategoryFeatureInfo `json:"featureInfos,omitempty"`
}

// ChildCategoryInfo 子类目信息
type ChildCategoryInfo struct {
	// ID 子类目ID
	ID uint64 `json:"id,omitempty"`
	// Name 子类目名称
	Name string `json:"name,omitempty"`
}

// CategoryFeatureInfo 类目特征信息
type CategoryFeatureInfo struct {
	// Key 名称
	Key string `json:"key,omitempty"`
	// Value 值
	Value string `json:"value,omitempty"`
	// Status 状态
	Status int `json:"status,omitempty"`
	// Hierarchy 是否继承到子元素上
	Hierarchy bool `json:"hierarchy,omitempty"`
}

// CategoryGet 根据类目Id查询类目
func CategoryGet(client *go1688.Client, req *CategoryGetRequest, accessToken string) (*CategoryInfo, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	var resp CategoryGetResponse
	if err := client.Do(finalRequest, accessToken, &resp); err != nil {
		return nil, err
	}
	return resp.Category, nil
}
