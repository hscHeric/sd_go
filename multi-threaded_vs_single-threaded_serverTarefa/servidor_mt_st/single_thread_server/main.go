package main

import (
	"fmt"
	"net"
)

func main() {
	port := "7896"
	listenSocket, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Println("Listen:", err)
	}

	defer listenSocket.Close()

	fmt.Println("Servidor TCP pronto para aceitar conexões na porta", port)

	for {
		clientSocket, err := listenSocket.Accept()
		if err != nil {
			fmt.Println("Accept:", err)
			continue
		}

		func(conn net.Conn) {
			connection := &Connection{
				conn: conn,
			}
			connection.HandleConnection()
		}(clientSocket)
	}
}
