/**
API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。

facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。


*/
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
