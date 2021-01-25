package simple

var _ Api = (*FactoryOne)(nil)
var _ Api = (*FactoryTwo)(nil)

type Api interface {
	Say() string
}
type FactoryOne struct {
}
type FactoryTwo struct {
}

func (f FactoryOne) Say() string {
	return "hello world"
}
func (f FactoryTwo) Say() string {
	return "hello world"
}

func NewFactory(typ int) Api {
	switch typ {
	case 1:
		return &FactoryOne{}
	case 2:
		return &FactoryTwo{}
	}
	return nil
}
