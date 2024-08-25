package proxy

type Proxy interface {
	Add(a, b float64) (float64, error)
	Sub(a, b float64) (float64, error)
	Mul(a, b float64) (float64, error)
	Div(a, b float64) (float64, error)
	Close() error
}
