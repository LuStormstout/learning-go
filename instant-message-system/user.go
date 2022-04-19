package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// NewUser 创建用户 API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// 启动监听当前 User channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

// ListenMessage 监听当前 User channel 的方法，一旦有消息就直接发送给客户端
func (u User) ListenMessage() {
	for {
		msg := <-u.C
		u.conn.Write([]byte(msg + "\n"))
	}
}
