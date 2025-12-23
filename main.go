package main

import (
	"fmt"
	"time"
)


func main() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("test")
	}

}

// docker run --rm -it -v "$PWD":/app -w /app golang:1.21-alpine go run main.go
