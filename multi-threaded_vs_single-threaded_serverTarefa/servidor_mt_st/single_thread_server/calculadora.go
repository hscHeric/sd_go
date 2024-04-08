package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

type Calculadora struct{}

var calculadoraInstance *Calculadora

func getInstance() *Calculadora {
	if calculadoraInstance == nil {
		once.Do(func() {
			calculadoraInstance = &Calculadora{}
		})
	}
	return calculadoraInstance
}

func (Calculadora) Calculate(operation string, num1, num2 float64) (float64, error) {
	time.Sleep(100 * time.Millisecond)

	switch operation {
	case "add":
		return num1 + num2, nil
	case "sub":
		return num1 - num2, nil
	case "mult":
		return num1 * num2, nil
	case "div":
		if num2 == 0 {
			return 0, fmt.Errorf("divisão por zero não é permitida")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("operação não suportada: %s", operation)
	}
}

func (Calculadora) Add(op1, op2 float64) float64 {
	return op1 + op2
}

func (Calculadora) Sub(op1, op2 float64) float64 {
	return op1 - op2
}

func (Calculadora) Mult(op1, op2 float64) float64 {
	return op1 * op2
}

func (Calculadora) Div(op1, op2 float64) (float64, error) {
	if op2 == 0 {
		return 0, fmt.Errorf("divisão por zero")
	}
	return op1 / op2, nil
}
