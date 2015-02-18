package main

import "fmt"

//PICK1 OMIT
func send(messages chan<- string, name string) {
	for i := 0; i < 5; i++ {
		messages <- fmt.Sprintf("%s - %d", name, i)
	}
	close(messages)
}

func receive(messages <-chan string) {
	for message := range messages {
		fmt.Printf("Hello, %s.\n", message)
	}
}

func main() {
	messages := make(chan string)
	go send(messages, "Jeff")
	receive(messages)
	for _, char := range "test" {
		fmt.Println(char)
	}
}

//END1 OMIT
