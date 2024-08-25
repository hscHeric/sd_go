package main

import (
	"Server/internal/dispatcher"
	"Server/internal/tcpserver"
	"log"
	"net"
	"strconv"
	"strings"
)

const port = "1234"

func main() {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
	defer listener.Close()

	dispatch := dispatcher.NewDispatcher()
	log.Println("Servidor ouvindo na porta :" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleConnection(conn, dispatch)
	}
}

func handleConnection(conn net.Conn, dispatch *dispatcher.Dispatcher) {
	defer conn.Close()

	for {
		// Lê a requisição do cliente
		request, err := tcpserver.GetRequest(conn)
		if err != nil {
			log.Println("Erro ao ler requisição:", err)
			return
		}
		if request == "" {
			continue // Ignore mensagens vazias
		}

		// Processa a requisição
		parts := strings.Fields(request)
		if len(parts) != 3 {
			tcpserver.SendResponse(conn, "Erro: formato de requisição inválido\n")
			continue
		}

		operation := parts[0]
		a, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			tcpserver.SendResponse(conn, "Erro: parâmetro inválido\n")
			continue
		}
		b, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			tcpserver.SendResponse(conn, "Erro: parâmetro inválido\n")
			continue
		}

		// Invoca o método apropriado
		response, err := dispatch.Invoke(operation, a, b)
		if err != nil {
			tcpserver.SendResponse(conn, "Erro: "+err.Error()+"\n")
			continue
		}

		// Envia a resposta
		tcpserver.SendResponse(conn, response+"\n")
	}
}

