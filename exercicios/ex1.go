package main

import (
	"fmt"
	"time"
)


func contador(tipo string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s: %d\n", tipo, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func workerCalcula(workerID string, msg chan int) {
	for res := range msg {
		fmt.Printf("Worker %s recebeu %d\n", workerID, res)
		time.Sleep(time.Second)
	}
}


func main() {
	msg := make(chan int)

	go workerCalcula("João", msg)
	go workerCalcula("Will", msg)
	go workerCalcula("Pedrin", msg)

	for i := 1; i < 20; i++ {
		msg <- i
	}
}
