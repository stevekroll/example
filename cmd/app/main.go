package main

import (
	"example/internal/services/myservice"
	"example/internal/workers/myworker"
	"fmt"
)

func main() {
	myWorker := myworker.New("MyWorker", nil)
	myService := myservice.New("MyService", myWorker)
	fmt.Printf("\nSERVICE HOST | %s\n", myService.Host())
	fmt.Printf("WORKER HOST  | %s\n\n", myService.WorkerHost())
}
