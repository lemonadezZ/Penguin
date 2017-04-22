package penguin

import (
	"fmt"
	"net"
)

const hostname string = "127.0.0.1"
const port int = 8082

//桐城服务器
func TCPServer() {
	ip := net.ParseIP(hostname)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{ip, port, ""})
	if err != nil {
		fmt.Println("端口监听错误", err.Error())
		return
	}
	RunServer(listen)

}

//执行tcp服务器
func RunServer(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接收请求报错", err.Error())
			continue
		}

		fmt.Println("客户端来自:" + conn.RemoteAddr().String())
		defer conn.Close()
		fmt.Println("执行go协程处理客户连接")
		go func() {
			data := make([]byte, 128)
			for {
				i, err := conn.Read(data)
				fmt.Println("客户端的数据为", string(data[0:i]))
				if err != nil {
					fmt.Println("数据读取错误", err.Error())
					break
				}
				r := string(data[0:i])
				if r == string([]byte("\r\n")) {
					conn.Write([]byte("\r\nfinish\r\n"))
				}
			}
		}()
	}
}
