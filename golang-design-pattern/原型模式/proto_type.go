package proto_type

type CloneAble interface {
	Clone() CloneAble
}

type ProtoTypeManager struct {
	protoTypes map[string]CloneAble
}

func NewProtoTypeManager() *ProtoTypeManager {
	return &ProtoTypeManager{
		protoTypes: make(map[string]CloneAble),
	}
}

func (p *ProtoTypeManager) Get(name string) CloneAble {
	return p.protoTypes[name]
}

func (p *ProtoTypeManager) Set(name string, prototype CloneAble) {
	p.protoTypes[name] = prototype
}
