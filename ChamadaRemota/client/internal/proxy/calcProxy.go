package proxy

import (
	"Client/internal/tcpclient"
	"fmt"
)

type CalcProxy struct {
	client *tcpclient.TCPClient
}

func NewCalcProxy(client *tcpclient.TCPClient) *CalcProxy {
	return &CalcProxy{client: client}
}

func (p *CalcProxy) Add(a, b float64) (float64, error) {
	request := fmt.Sprintf("add %f %f", a, b)
	err := p.client.SendRequest(request)
	if err != nil {
		return 0, err
	}

	response, err := p.client.GetResponse()
	if err != nil {
		return 0, err
	}

	var result float64
	fmt.Sscanf(response, "%f", &result)

	return result, nil
}

func (p *CalcProxy) Sub(a, b float64) (float64, error) {
	request := fmt.Sprintf("sub %f %f", a, b)
	err := p.client.SendRequest(request)
	if err != nil {
		return 0, err
	}
	response, err := p.client.GetResponse()
	if err != nil {
		return 0, err
	}
	var result float64
	fmt.Sscanf(response, "%f", &result)
	return result, nil
}

func (p *CalcProxy) Mul(a, b float64) (float64, error) {
	request := fmt.Sprintf("mul %f %f", a, b)
	err := p.client.SendRequest(request)
	if err != nil {
		return 0, err
	}
	response, err := p.client.GetResponse()
	if err != nil {
		return 0, err
	}
	var result float64
	fmt.Sscanf(response, "%f", &result)
	return result, nil
}

func (p *CalcProxy) Div(a, b float64) (float64, error) {
	request := fmt.Sprintf("div %f %f", a, b)
	err := p.client.SendRequest(request)
	if err != nil {
		return 0, err
	}
	response, err := p.client.GetResponse()
	if err != nil {
		return 0, err
	}
	var result float64
	fmt.Sscanf(response, "%f", &result)
	return result, nil
}

func (p *CalcProxy) Close() error {
	return p.client.Close()
}
