package factory_method

type IFactory interface {
	Get() string
}

type InitFactory interface {
	Create() IFactory
}

type FactoryOne struct{}

func (FactoryOne) Create() IFactory {
	return &FactoryOne{}
}
func (f FactoryOne) Get() string {
	return "hello world"
}

type FactoryTwo struct{}

func (FactoryTwo) Create() IFactory {
	return &FactoryOne{}
}
func (f FactoryTwo) Get() string {
	return "hello world"
}
