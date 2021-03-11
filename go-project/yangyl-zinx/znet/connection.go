package znet

import (
	"fmt"
	"go-study-project/go-project/yangyl-zinx/ziface"
	"io"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//当前连接的关闭状态
	isClosed bool

	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool

	msgHandler ziface.IMsgHandler
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		ExitBuffChan: make(chan bool, 1),
		msgHandler:   msgHandler,
	}
}

func (c *Connection) StartReader() {
	defer c.Stop()
	for {
		dp := new(DataPack)

		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error ", err)
			c.ExitBuffChan <- true
			return
		}

		//拆包，得到msgid 和 datalen 放在msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error ", err)
			c.ExitBuffChan <- true
			return
		}

		//根据 dataLen 读取 data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				c.ExitBuffChan <- true
				return
			}
		}

		msg.SetData(data)
		req := NewRequest(c, msg)

		c.msgHandler.SendMsgToTaskQueue(req)

	}
}

func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			fmt.Println("ExitBuffChan 退出")
			return
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("Stop 退出")
	if c.isClosed {
		return
	}

	c.Conn.Close()

	c.ExitBuffChan <- true

	c.isClosed = true

	close(c.ExitBuffChan)
}
