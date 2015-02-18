package main

import "fmt"

func send(messages chan<- string, name string) {
	messages <- fmt.Sprintf("%s", name)
}

func receive(messages <-chan string) {
	message := <-messages
	fmt.Printf("Hello, %s.\n", message)
}

func main() {
	messages := make(chan string)
	go send(messages, "Jeff")
	receive(messages)
}
