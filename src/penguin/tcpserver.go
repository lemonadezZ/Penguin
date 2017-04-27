package penguin

import (
	"bytes"
	"fmt"
	"net"
	"time"
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
			lasttimeChan := make(chan int64, time.Now().Unix())
			lasttimeChan <- time.Now().Unix()
			go func(conn *net.TCPConn, lasttimeChan chan int64) {
				overtime := <-lasttimeChan
				//判定是否超时
				for {
					//大于10秒发ping 包
					select {
					case el := <-lasttimeChan:
						overtime = el
						// fmt.Println("收到最后更新数据")
						// fmt.Println(el)
					default:
						// now := time.Now().Unix()
						// fmt.Print(overtime)
						// fmt.Print(now)
						// fmt.Println(now-)
						if (time.Now().Unix() - overtime) > 10 {
							conn.Write([]byte(time.Now().Format("2006-01-02 03:04:05 PM")))
							conn.Write([]byte("ping\r\nS:"))
						} else {
							// fmt.Println("没有超时")
							// fmt.Println(overtime)
						}
					}
					time.Sleep(time.Second * 2)
					// if (overtime - time.Now().Unix()) > 10 {
					// 	conn.Write([]byte(time.Now().Format("2006-01-02 03:04:05 PM")))
					// 	conn.Write([]byte("ping\r\nS:"))
					// } else {
					// 	fmt.Println("没有超时")
					// 	fmt.Println(overtime)
					// }
				}
			}(conn, lasttimeChan)
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
					lasttimeChan <- time.Now().Unix()
					conn.Write([]byte(time.Now().Format("2006-01-02 03:04:05 PM")))
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
