package product

import (
	"github.com/bububa/go1688"
)

type CategoryGetRequest struct {
	Id uint64 `json:"categoryID"` // 类目id,必须大于等于0， 如果为0，则查询所有一级类目
}

func (this *CategoryGetRequest) Name() string {
	return "alibaba.category.get"
}

type CategoryGetResponse struct {
	go1688.BaseResponse
	Category *CategoryInfo `json:"categoryInfo,omitempty"`
}

type CategoryInfo struct {
	Id               uint64                `json:"categoryID,omitempty"`       // 类目ID
	Name             string                `json:"name,omitempty"`             // 类目名称
	IsLeaf           bool                  `json:"isLeaf,omitempty"`           // 是否叶子类目（只有叶子类目才能发布商品）
	ParentIds        []uint64              `json:"parentIDs,omitempty"`        // 父类目ID数组,1688只返回一个父id
	MinOrderQuantity uint64                `json:"minOrderQuantity,omitempty"` // 最小起订量
	ChildCategories  []ChildCategoryInfo   `json:"childCategorys,omitempty"`   // 子类目信息
	Features         []CategoryFeatureInfo `json:"featureInfos,omitempty"`     // 类目特征信息
}

type ChildCategoryInfo struct {
	Id   uint64 `json:"id,omitempty"`   // 子类目ID
	Name string `json:"name,omitempty"` // 子类目名称
}

type CategoryFeatureInfo struct {
	Key       string `json:"key,omitempty"` // 名称
	Value     string `json:"value,omitempty"`
	Status    int    `json:"status,omitempty"`
	Hierarchy bool   `json:"hierarchy,omitempty"` //是否继承到子元素上
}

func CategoryGet(client *go1688.Client, req *CategoryGetRequest, accessToken string) (*CategoryInfo, error) {
	finalRequest := go1688.NewRequest(NAMESPACE, req)
	resp := &CategoryGetResponse{}
	err := client.Do(finalRequest, accessToken, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}
	return resp.Category, nil
}
