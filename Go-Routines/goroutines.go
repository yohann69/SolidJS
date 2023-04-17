package main

import (
	"fmt"
	"sync"
	"time"
)

func block(s int, wg *sync.WaitGroup, allDone *sync.WaitGroup) {
	// defer wg.Done() // Est appelé à la fin de la fonction

	// 1ere partie
	start := time.Now()
	time.Sleep(time.Second * time.Duration(s))

	wg.Done()

	allDone.Wait()

	// 2ème partie
	time.Sleep(time.Second * time.Duration(s))
	wg.Done()

	fmt.Printf("Durée d'éxecution: %s\n", time.Since(start))

}

func main() {

	start := time.Now()

	var wg sync.WaitGroup
	var allDone sync.WaitGroup

	allDone.Add(1)

	wg.Add(3)

	go block(5, &wg, &allDone)
	go block(1, &wg, &allDone)
	go block(3, &wg, &allDone)

	wg.Wait()
	wg.Add(3)
	allDone.Done()
	wg.Wait()

	fmt.Printf("Durée d'éxecution: %s\n", time.Since(start))

}
