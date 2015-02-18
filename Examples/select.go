package main

import "fmt"

//PICK1 OMIT
func send(messages chan<- string, name string) {
	for i := 0; i < 5; i++ {
		messages <- fmt.Sprintf("%s - %d", name, i)
	}
}

func receive(messages <-chan string) {
	select {
	case message := <-messages:
		fmt.Printf("Hello, %s.\n", message)
	default:
		fmt.Printf("No message.\n")
	}
}

//same main function
//END1 OMIT

func main() {
	messages := make(chan string, 1)
	go send(messages, "Jeff")
	receive(messages)
	for _, char := range "test" {
		fmt.Println(char)
	}
}
