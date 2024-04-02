package main

import (
	"fmt"
	"net"
)

type TCPServer struct {
	listener net.Listener
}

func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (t *TCPServer) Connect(port string) error {
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return err
	}

	t.listener = listener

	defer t.listener.Close()

	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("Error ao tentar aceitar conexão:", err)
			continue
		}
		go t.handleConnection(conn)
	}
}

func (t *TCPServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		fmt.Print(string(buf[:n]))
	}
}
