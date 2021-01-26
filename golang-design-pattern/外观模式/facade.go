package facade

type Facade struct {
	One IFacadeOne
}

func NewFacade() *Facade {
	return &Facade{}
}

type IFacadeOne interface {
	Get() string
}

type FacadeOne struct {
	Name string
}

func (f *FacadeOne) Get() string {
	return f.Name
}
