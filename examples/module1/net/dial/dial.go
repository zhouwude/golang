// make a connection with www.example.org:
package main

import (
	"fmt"
	"net"
	"os"
)

// 先使用 TCP 协议连接远程 80 端口，然后使用 UDP 协议连接，最后使用 TCP 协议连接 IPv6 类型的地址：
func main() {
	conn, err := net.Dial("tcp", "192.0.32.10:80") // tcp ipv4
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "192.0.32.10:80") // udp
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") // tcp ipv6
	checkConnection(conn, err)
}
func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
	/*Connection is made with &{{0xc000124080}}
	Connection is made with &{{0xc000124180}}
	error dial tcp [2620:0:2d0:200::10]:80: connect: no route to host connecting!exit status 1
	*/
}
