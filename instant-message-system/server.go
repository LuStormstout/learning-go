package main

import (
	"fmt"
	"net"
	"sync"
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
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	server.Message <- sendMsg
}

func (server *Server) Handler(conn net.Conn) {
	// ...当前连接的业务
	// fmt.Println("连接建立成功")

	user := NewUser(conn)

	// 用户上线，将用户加入到 OnlineMap 中
	server.mapLock.Lock()
	server.OnlineMap[user.Name] = user
	server.mapLock.Unlock()

	// 广播当前用户上线消息
	server.BroadCast(user, "已上线")

	// 当前 handle 阻塞
	select {}
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
