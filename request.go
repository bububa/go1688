package go1688

import (
	"fmt"
)

type Request interface {
	Namespace() string
	Name() string
	Version() string
	Path() string
	Params() map[string]string
}

type BaseRequest struct {
	namespace string
	name      string
	version   string
}

func NewBaseRequest(namespace string, name string) BaseRequest {
	return BaseRequest{
		namespace: namespace,
		name:      name,
	}
}

func (this *BaseRequest) SetVersion(version string) {
	this.version = version
}

func (this *BaseRequest) Name() string {
	return this.name
}

func (this *BaseRequest) Namespace() string {
	return this.namespace
}

func (this *BaseRequest) Version() string {
	if this.version == "" {
		return VERSION
	}
	return this.version
}

func (this *BaseRequest) Path() string {
	return fmt.Sprintf("%s/%s/%s", this.Version(), this.Namespace(), this.Name())
}

type RequestData interface {
	Name() string
}

type FinalRequest struct {
	BaseRequest
	data interface{}
}

func NewRequest(namespace string, data RequestData) *FinalRequest {
	return &FinalRequest{
		BaseRequest: NewBaseRequest(namespace, data.Name()),
		data:        data,
	}
}

func (this *FinalRequest) Params() map[string]string {
	return structToMap(this.data)
}
