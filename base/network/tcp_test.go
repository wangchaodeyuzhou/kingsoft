package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"testing"
)

func TestTcpNetWork(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}

	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err", err)
			continue
		}
		go process(conn)
	}

}

func TestTcpClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer conn.Close()

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		buf := [1024]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func Encode(msg string) ([]byte, error) {
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}

func ProcessTcp(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := Decode(reader)
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("decode msg fialed", err)
			return
		}

		fmt.Println(string(msg))
	}
}

func TestTcpServerStickPkg(t *testing.T) {
	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("listen fialed err", err)
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept fialed err", err)
			return
		}

		go ProcessTcp(conn)
	}
}

func TestTcpStickPkgClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":5000")
	if err != nil {
		fmt.Println("dail fialed err", err)
		return
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "Hello world, err"
		data, err := Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err", err)
			return
		}

		conn.Write(data)
	}
}
