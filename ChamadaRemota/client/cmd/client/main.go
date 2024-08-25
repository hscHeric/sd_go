package main

import (
	"Client/internal/proxy"
	"Client/internal/tcpclient"
	"Client/internal/user"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const port = "1234"

func main() {
	c, err := tcpclient.NewTCPClient("localhost:" + port)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := c.Close(); err != nil {
			log.Println("Erro ao fechar a conexão:", err)
		}
	}()

	calcProxy := proxy.NewCalcProxy(c)
	userInterface := user.NewUser(calcProxy)

	// Canal para capturar sinais de interrupção
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine para capturar sinais e finalizar o programa
	go func() {
		<-sigs
		fmt.Println("\nFinalizando o programa...")
		userInterface.Close() // Fecha o proxy e outras conexões se necessário
		os.Exit(0)
	}()

	for {
		userInterface.PerformOperation()
	}
}
