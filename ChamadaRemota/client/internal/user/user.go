package user

import (
	"Client/internal/proxy"
	"fmt"
)

type User struct {
	Proxy proxy.Proxy
}

func NewUser(proxy proxy.Proxy) *User {
	return &User{Proxy: proxy}
}

func (u *User) PerformOperation() {
	fmt.Println("Escolha a operação que deseja realizar (add, sub, mul, div): ")
	var op string
	fmt.Scanln(&op)

	fmt.Println("Coloque dois números (exemplo: 3.2 1.4): ")
	var a, b float64

	fmt.Scanln(&a, &b)
	var result float64
	var err error
	switch op {
	case "add":
		result, err = u.Proxy.Add(a, b)
	case "sub":
		result, err = u.Proxy.Sub(a, b)
	case "mul":
		result, err = u.Proxy.Mul(a, b)
	case "div":
		result, err = u.Proxy.Div(a, b)
	default:
		fmt.Println("Invalid operation")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
}

func (u *User) Close() {
	u.Proxy.Close()
}
