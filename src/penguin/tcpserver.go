package penguin

import (
	"bytes"
	"fmt"
	"net"
)

const hostname string = "0.0.0.0"
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

			var buffer bytes.Buffer
			conn.Write([]byte("hello Welcome to penguin\r\n"))
			conn.Write([]byte("S:"))
			for {
				i, err := conn.Read(data)
				buffer.Write(data[0:i])
				if err != nil {
					fmt.Println("数据读取错误", err.Error())
					break
				}
				//优化回显
				// r := string(data[0:i])
				next := data[i-1 : i][0]
				if next == 10 {
					//clean
					fmt.Print(conn.RemoteAddr().String() + " :")
					fmt.Print(buffer.String())
					fmt.Println(buffer.Bytes())
					conn.Write([]byte("R:"))
					conn.Write(buffer.Bytes())
					buffer.Reset()
					conn.Write([]byte("S:"))
				}
			}
		}()
	}
}
