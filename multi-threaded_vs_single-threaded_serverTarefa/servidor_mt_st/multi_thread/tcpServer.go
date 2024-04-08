package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Connection struct {
	conn net.Conn
}

func (c *Connection) getRequest() (string, error) {
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

func (c *Connection) SendResponse(response string) error {
	_, err := c.conn.Write([]byte(response))
	return err
}

func handleCalcString(request string) (string, float64, float64, error) {
	parts := strings.Fields(request)
	if len(parts) != 3 {
		return "", 0, 0, fmt.Errorf("formato de entrada inválido. Deve ser 'operacao numero1 numero2'")
	}

	operation := parts[0]
	num1, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("erro ao converter o primeiro número: %v", err)
	}

	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("erro ao converter o segundo número: %v", err)
	}

	return operation, num1, num2, nil
}

func (c *Connection) HandleConnection() {
	defer c.conn.Close()
	for {
		request, err := c.getRequest()
		if err != nil {
			fmt.Println("Erro ao receber requisição:", err)
			return
		}

		operation, num1, num2, err := handleCalcString(request)
		if err != nil {
			c.SendResponse("Erro ao converter os números: " + err.Error() + "\n")
			continue
		}

		calc := getInstance()
		result, err := calc.Calculate(operation, num1, num2)
		if err != nil {
			c.SendResponse("Erro ao calcular: " + err.Error() + "\n")
			continue
		}

		response := strconv.FormatFloat(result, 'f', -1, 64)
		err = c.SendResponse(response)
		if err != nil {
			fmt.Println("Erro ao enviar resposta:", err)
			return
		}
	}
}
