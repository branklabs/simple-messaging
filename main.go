package main

import (
	"fmt"
	"kpc/internal"
	"time"
)

func main() {
	fmt.Println("simple-consumer-producer")

	ticker := time.NewTicker(5 * time.Minute)
	go internal.Producer()
	go internal.Consumer()

	select {
	case <-ticker.C:
		fmt.Println("shutting down")
	}
}
