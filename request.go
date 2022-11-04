package go1688

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

func (r *BaseRequest) SetVersion(version string) {
	r.version = version
}

func (r *BaseRequest) Name() string {
	return r.name
}

func (r *BaseRequest) Namespace() string {
	return r.namespace
}

func (r *BaseRequest) Version() string {
	if r.version == "" {
		return VERSION
	}
	return r.version
}

func (r *BaseRequest) Path() string {
	builder := GetStringsBuilder()
	defer PutStringsBuilder(builder)
	builder.WriteString(r.Version())
	builder.WriteString("/")
	builder.WriteString(r.Namespace())
	builder.WriteString("/")
	builder.WriteString(r.Name())
	return builder.String()
}

type RequestData interface {
	Name() string
	Map() map[string]string
}

type FinalRequest struct {
	BaseRequest
	data RequestData
}

func NewRequest(namespace string, data RequestData) *FinalRequest {
	return &FinalRequest{
		BaseRequest: NewBaseRequest(namespace, data.Name()),
		data:        data,
	}
}

func (r *FinalRequest) Params() map[string]string {
	return r.data.Map()
}
