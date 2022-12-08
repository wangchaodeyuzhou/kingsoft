package network

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read form client failed, err: ", err)
			break
		}

		recvStr := string(buf[:n])
		fmt.Println("client send msg", recvStr)
		conn.Write([]byte(recvStr))
	}
}
