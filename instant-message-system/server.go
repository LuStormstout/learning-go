package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的 channel
	Message chan string
}

// NewServer 创建一个 server 的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// ListenMessage 监听 Message 广播 channel 的 goroutine，一旦有消息就发送给全部在线的 User
func (server *Server) ListenMessage() {
	for true {
		msg := <-server.Message

		// 将 msg 发送给全部的在线 User
		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.mapLock.Unlock()
	}
}

// BroadCast 广播消息的方法
func (server *Server) BroadCast(user *User, msg string) {
	sendMsg := "[广播消息]：[" + user.Addr + "]" + user.Name + ":" + msg

	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	// ...当前连接的业务
	// fmt.Println("连接建立成功")

	user := NewUser(conn, server)

	// 用户上线
	user.Online()

	// 监听当前用户是否处于活跃状态
	isActive := make(chan bool)

	// 等待接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			// 如果读到的字节数为 0 则表示该客户端正常关闭，广播该用户下线
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline() // 用户下线
				return
			}

			// 读取失败
			if err != nil && err != io.EOF {
				fmt.Println("conn read error:", err)
				return
			}

			// 提取用户的消息（去除"\n"）
			msg := string(buf[:n-1])

			// 用户针对 msg 进行消息处理
			user.DoMessage(msg)

			// 用户的任何消息，代表当前用户是一个活跃的
			isActive <- true
		}
	}()

	// 当前 handle 阻塞
	for {
		select {
		case <-isActive:
			// 说明当前用户活跃，应该重置定时器
			// 不做任何事情，为了激活 select，更新下面的定时器
		case <-time.After(time.Second * 1800): // 为了测试的话可以将这个时间修改的短一点儿，目前是 30 分钟后就会被踢下线
			// 已经超时
			// 将当前的 User 强制关闭（超时不活跃踢下线）
			user.SendMsg("[系统消息]：由于你长时间处于不活跃状态，你已被踢下线。")

			// 销毁该用户的资源（关闭该用户的 channel）
			close(user.C)
			// 关闭连接
			conn.Close()
			// 退出当前 Handler
			return // runtime.Goexit()
		}
	}
}

// Start 启动服务器的接口
func (server *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("new.Listen error:", err)
		return
	}

	// close listen socket
	defer listener.Close()

	// 启动监听 Message 的 goroutine
	go server.ListenMessage()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}

		// do handle
		go server.Handler(conn)
	}
}
