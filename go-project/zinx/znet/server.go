package znet

import (
	"fmt"
	"go-study-project/go-project/zinx/ziface"
	"net"
	"time"
)

type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int

	Router map[uint32]ziface.IRouter
}

//开启网络服务
func (s *Server) Start() {
	//开启一个go去做服务端Linster业务
	go func() {
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		var cid uint32 = 0

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, _ := listenner.AcceptTCP()
			dealConn := NewConnection(conn, cid, s.Router)

			cid++

			go dealConn.Start()

		}
	}()
}

func (s *Server) Stop() {

}

type UserRouter struct {
	BaseRouter
}

func (u *UserRouter) Handle(req ziface.IRequest) {
}

func (s *Server) Serve() {
	user := new(UserRouter)
	s.AddRouter(0, user)

	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
		Router:    make(map[uint32]ziface.IRouter, 0),
	}
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := s.Router[msgID]; !ok {
		s.Router[msgID] = router
	}

}
