package ziface

type IRequest interface {
	GetConnnect() IConnection // 获取连接信息
	GetData() []byte          // 获取数据
	GetMsgID() uint32
}
