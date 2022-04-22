package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	modeFlag   int // 当前客户端被选择的模式（模式标志）
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		modeFlag:   999, // 给一个默认值，不然为 0 的话菜单就退出了
	}

	// 连接 server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}
	client.conn = conn

	// 返回对象
	return client
}

func (client *Client) menu() bool {
	var modeFlag int

	fmt.Println("1.广播模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&modeFlag)

	if modeFlag >= 0 && modeFlag <= 3 {
		client.modeFlag = modeFlag
		return true
	} else {
		fmt.Println(">>> 请正确输入菜单列表中的序号 <<<")
		return false
	}
}

func (client *Client) Run() {
	for client.modeFlag != 0 {
		for client.menu() != true {

		}

		// 根据选择的模式处理不同的业务
		switch client.modeFlag {
		case 1:
			// 广播模式
			fmt.Println("广播模式选择...")
			break
		case 2:
			// 私聊模式
			fmt.Println("私聊模式选择...")
			break
		case 3:
			// 更新用户名
			fmt.Println("更新用户名选择...")
			break
		}
	}
}

var serverIp string
var serverPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器 IP 地址（默认是 127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是 8888）")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>> 服务器连接失败...")
		return
	}
	fmt.Println(">>> 服务器连接成功...")

	// 启动客户端业务
	client.Run()
}
