package main

import (
	"fmt"
	"sync"
	"time"
)

var orquestrador sync.WaitGroup

func main() {
	orquestrador.Add(2)
	go tarefa1()
	go tarefa2()
	orquestrador.Wait()
}

func tarefa1() {
	fmt.Println(time.Now(), " - Comencando tarefa 1")
	time.Sleep(3 * time.Second)
	fmt.Println(time.Now(), " - Terminando tarefa 1")
	orquestrador.Done()
}

func tarefa2() {
	fmt.Println(time.Now(), " - Comencando tarefa 2")
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now(), " - Terminando tarefa 2")
	orquestrador.Done()
}
