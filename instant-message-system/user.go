package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// NewUser 创建用户 API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前 User channel 消息的 goroutine
	go user.ListenMessage()

	return user
}

// Online 用户上线的业务
func (user *User) Online() {
	// 用户上线，将用户加入到 OnlineMap 中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 广播当前用户上线消息
	user.server.BroadCast(user, "已上线")
}

// Offline 用户下线的业务
func (user *User) Offline() {
	// 用户下线，将用户加入到 OnlineMap 中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	// 广播当前用户下线消息
	user.server.BroadCast(user, "下线")
}

// SendMsg 给当前 user 对应的客户端发消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// DoMessage 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	if msg == "/who" {
		// 查询当前有哪些用户在线（用"/who"指令来查询当前都有谁在线）
		user.server.mapLock.Lock()
		for _, onlineUser := range user.server.OnlineMap {
			whoOnlineMsg := "[" + onlineUser.Addr + "]" + onlineUser.Name + ":" + "在线 ...\n"
			user.SendMsg(whoOnlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 8 && msg[:8] == "/rename|" {
		// 消息格式 "/rename|张三"
		newName := strings.Split(msg, "|")[1]

		// 判断 newName 是否已经存在了（要修改的用户名是否已被占用了）
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.SendMsg("当前用户名已经被占用，请重新修改\n")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.Name = newName
			user.SendMsg("用户名修改成功：" + user.Name + "\n")
		}
	} else {
		user.server.BroadCast(user, msg)
	}
}

// ListenMessage 监听当前 User channel 的方法，一旦有消息就直接发送给客户端
func (user User) ListenMessage() {
	for {
		msg := <-user.C
		user.conn.Write([]byte(msg + "\n"))
	}
}
