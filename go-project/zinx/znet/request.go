package znet

import "go-study-project/go-project/zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data ziface.IMessage
}

func (r *Request) GetConnnect() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data.GetData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.data.GetMsgId()
}

func AddRouter() {

}

func NewRequest(conn ziface.IConnection, data ziface.IMessage) *Request {
	return &Request{
		conn: conn,
		data: data,
	}
}
