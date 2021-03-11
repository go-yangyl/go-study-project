package znet

import (
	"go-study-project/go-project/yangyl-zinx/ziface"
)

type MsgHandler struct {
	routers        map[uint32]ziface.IRouter
	WorkerPoolSize uint32                 //业务工作Worker池的数量
	TaskQueue      []chan ziface.IRequest //Worker负责取任务的消息队列
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := m.routers[msgID]; !ok {
		m.routers[msgID] = router
	}
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	if router, ok := m.routers[request.GetMsgID()]; ok {
		router.Handle(request)
	}
}

func (m *MsgHandler) StartWorkerPool() {
	for i := 0; i < 10; i++ {
		m.TaskQueue[i] = make(chan ziface.IRequest, m.WorkerPoolSize)
		for {
			select {
			case request, ok := <-m.TaskQueue[i]:
				if ok {
					m.DoMsgHandler(request)
				}
			}
		}
	}
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		routers:        make(map[uint32]ziface.IRouter),
		WorkerPoolSize: 1024,
		TaskQueue:      make([]chan ziface.IRequest, 10),
	}
}

func (m *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	workID := request.GetMsgID() % m.WorkerPoolSize

	m.TaskQueue[workID] <- request
}
