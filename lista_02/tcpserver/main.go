package main

func main() {
	server := NewTCPServer()

	server.Connect("8080")
}
