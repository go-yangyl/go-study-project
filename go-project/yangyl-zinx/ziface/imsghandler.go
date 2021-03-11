package ziface

type IMsgHandler interface {
	AddRouter(msgID uint32, router IRouter)
	DoMsgHandler(request IRequest)
	StartWorkerPool()                    //启动worker工作池
	SendMsgToTaskQueue(request IRequest) //将消息交给TaskQueue,由worker进行处理
}
