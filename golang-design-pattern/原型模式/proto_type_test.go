package proto_type

import (
	"testing"
)

type Typ struct {
	Name string
}

func (t Typ) Clone() CloneAble {
	return Typ{
		Name: t.Name,
	}
}

func TestNewProtoTypeManager(t *testing.T) {
	proto := NewProtoTypeManager()
	typ := &Typ{Name: "hello world"}
	proto.Set("name", typ)
	ITyp := proto.Get("name")

	if res, ok := ITyp.(*Typ); ok {
		t.Log(res.Name, "111")
	}

}
