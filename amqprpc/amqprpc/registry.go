package amqprpc

type Registry struct {
	name    string
	methods map[string]Method
}

func NewRegistry(name string) *Registry {
	return &Registry{name: name, methods: make(map[string]Method)}
}

func (r *Registry) AddMethod(que string, meth Method) {
	r.methods[que] = meth
}

func (r *Registry) GetMethods() map[string]Method {
	return r.methods
}
