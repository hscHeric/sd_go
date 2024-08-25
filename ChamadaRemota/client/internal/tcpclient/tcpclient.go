package tcpclient

import (
	"errors"
	"net"
)

type TCPClient struct {
	conn net.Conn
}

func NewTCPClient(addr string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TCPClient{conn: conn}, nil
}

func (c *TCPClient) SendRequest(r string) error {
	if c.conn == nil {
		return errors.New("não existe coneção estabelecida")
	}

	data := []byte(r + "\n")
	_, err := c.conn.Write(data)
	return err
}

func (c *TCPClient) GetResponse() (string, error) {
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func (c *TCPClient) Close() error {
	return c.conn.Close()
}
