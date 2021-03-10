package ziface

type IRouter interface{
	Handle(request IRequest)	 //处理conn业务的方法
}
