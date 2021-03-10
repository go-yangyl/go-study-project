package ziface

type IMsgHandler interface {
	AddRouter(msgID uint32, router IRouter)
	DoMsgHandler(request IRequest)
}
