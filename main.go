package main

func main() {
	server := NewApiServer(":8085")

	server.Run()
}