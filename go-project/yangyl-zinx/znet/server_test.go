package znet

import (
	"fmt"
	"go-study-project/go-project/yangyl-zinx/ziface"
	"net"
	"testing"
	"time"
)

func (u *UserRouter) Handle(req ziface.IRequest) {
	fmt.Println(111)
}

type UserRouter struct {
	BaseRouter
}

func TestNewServer(t *testing.T) {
	server := NewServer("yangyl-zinx")

	user := new(UserRouter)

	server.AddRouter(0, user)
	server.AddRouter(1, user)

	go ClientTest()

	server.Serve()

}

/*
	模拟客户端
*/
func ClientTest() {

	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		dp := new(DataPack)
		msg, _ := dp.Pack(NewMsgPackage(0, []byte("123322131232321")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
