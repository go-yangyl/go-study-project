package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	server := NewServer("zinx")

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
		msg, _ := dp.Pack(NewMsgPackage(0, []byte("hello")))
		conn.Write(msg)

		time.Sleep(1 * time.Second)
	}
}
