package skeleton

import (
	"Server/internal/calc"
	"fmt"
)

type Skeleton struct{}

func NewSkeleton() *Skeleton {
	return &Skeleton{}
}

func (s *Skeleton) Add(a, b float64) (string, error) {
	result := calc.Add(a, b)
	return fmt.Sprintf("%f", result), nil
}

func (s *Skeleton) Sub(a, b float64) (string, error) {
	result := calc.Sub(a, b)
	return fmt.Sprintf("%f", result), nil
}

func (s *Skeleton) Mul(a, b float64) (string, error) {
	result := calc.Mul(a, b)
	return fmt.Sprintf("%f", result), nil
}

func (s *Skeleton) Div(a, b float64) (string, error) {
	result, err := calc.Div(a, b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%f", result), nil
}
