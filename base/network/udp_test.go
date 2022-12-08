package network

import (
	"fmt"
	"net"
	"testing"
)

func TestUdpServer(t *testing.T) {
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 5000,
	})
	if err != nil {
		fmt.Println("listen failed err", err)
		return
	}

	defer socket.Close()

	for {
		var data [1024]byte
		n, addr, err := socket.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("recv udp server err", err)
			continue
		}

		fmt.Println(string(data[:n]), addr)

		_, err = socket.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("witer to udp failed, err", err)
			continue
		}
	}
}

func TestUdpClient(t *testing.T) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 5000,
	})

	if err != nil {
		fmt.Println("connect server fialed, err", err)
		return
	}

	defer socket.Close()

	sendData := []byte("Hello World")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println("send data failed", err)
		return
	}

	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("recv data failed, err", err)
		return
	}

	fmt.Println(string(data[:n]), remoteAddr)
}
