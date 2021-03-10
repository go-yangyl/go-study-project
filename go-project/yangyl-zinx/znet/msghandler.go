package znet

import (
	"go-study-project/go-project/yangyl-zinx/ziface"
)

type MsgHandler struct {
	routers map[uint32]ziface.IRouter
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

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		routers: make(map[uint32]ziface.IRouter),
	}
}
