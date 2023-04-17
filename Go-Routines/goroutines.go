package main

import (
	"fmt"
	"time"
)

func block(s int) {
	start := time.Now()
	time.Sleep(time.Second * time.Duration(s))
	fmt.Printf("Durée d'éxecution: %s\n", time.Since(start))
}

func main() {

	start := time.Now()

	go block(5)
	go block(1)
	go block(3)

	fmt.Printf("Durée d'éxecution: %s\n", time.Since(start))

}
