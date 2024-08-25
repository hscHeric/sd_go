package dispatcher

import (
	"Server/internal/skeleton"
	"fmt"
)

type Dispatcher struct {
	skeleton *skeleton.Skeleton
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		skeleton: skeleton.NewSkeleton(),
	}
}

func (d *Dispatcher) Invoke(operation string, a, b float64) (string, error) {
	switch operation {
	case "add":
		return d.skeleton.Add(a, b)
	case "sub":
		return d.skeleton.Sub(a, b)
	case "mul":
		return d.skeleton.Mul(a, b)
	case "div":
		return d.skeleton.Div(a, b)
	default:
		return "", fmt.Errorf("operação invalida")
	}
}
