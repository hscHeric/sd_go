package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	addr, port := "localhost", "7896"
	request := "add 10 20"
	var (
		erroSend    int
		erroRecv    int
		erroConnect int
		sucesso     int
	)

	rodadas := 100
	resposta := "30"
	erroSend = 0
	erroRecv = 0
	erroConnect = 0
	sucesso = 0

	var wg sync.WaitGroup
	wg.Add(rodadas)

	tempoInicialMT := time.Now()
	for i := 0; i < rodadas; i++ {
		go func() {
			defer wg.Done()
			client, err := NewClientTcp(addr, port)
			if err != nil {
				erroConnect++
				return
			}

			defer client.Close()
			err = client.SendRequest(request)
			if err != nil {
				erroSend++
				return
			}

			response, err := client.GetResponse()
			if string(response) != resposta {
				fmt.Println(response)
				erroRecv++
				return
			}
			if err != nil {
				erroRecv++
				return
			}
			sucesso++
			client.conn.Close()
		}()
	}

	wg.Wait()

	tempoFinalMT := time.Since(tempoInicialMT)

	fmt.Println("Erros de Connect:", erroConnect)
	fmt.Println("Erros de Send:", erroSend)
	fmt.Println("Erros de Recv:", erroRecv)
	fmt.Println("Sucesso:", sucesso)
	taxaDeSucesso := float64(sucesso) / float64(rodadas)
	fmt.Printf("Taxa de Sucesso: %.2f%%\n", taxaDeSucesso*100)
	fmt.Println("Tempo Total:", tempoFinalMT)
}
