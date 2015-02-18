package main

import (
	"fmt"
	"math/rand"
	"time"
)

func wait(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
	fmt.Printf("Hello World %d!\n", i)
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Deterministic randominess begone!
	for i := 0; i < 5; i++ {
		wait(i)
	}
	time.Sleep(time.Millisecond * 300)
}
