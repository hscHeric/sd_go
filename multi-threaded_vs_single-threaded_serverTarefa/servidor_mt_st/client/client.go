package main

import (
	"fmt"
	"net"
)

type ClientTCP struct {
	conn net.Conn
}

func NewClientTcp(ipServer string, port string) (*ClientTCP, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", ipServer, port))
	if err != nil {
		return nil, err
	}

	return &ClientTCP{
		conn: conn,
	}, nil
}

func (c *ClientTCP) SendRequest(request string) error {
	_, err := c.conn.Write([]byte(request))
	return err
}

func (c *ClientTCP) GetResponse() (string, error) {
	buffer := make([]byte, 1024)
	n, err := c.conn.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:n]), nil
}

func (c *ClientTCP) Close() {
	c.conn.Close()
}
