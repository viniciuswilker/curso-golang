package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// WAITGROUP

	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // contador

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go routine 3")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Go routine 3")
		waitGroup.Done() // -1
	}()
	waitGroup.Wait() //esperar a contagem da goroutines chegar em 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
