package main

import (
	"fmt"
	"sync"
	"time"
)

type MaStruct struct {
	Mutex sync.Mutex
	val   string
}

func block(s int, result chan bool, allDone *sync.WaitGroup, str *MaStruct) {
	// defer wg.Done() // Est appelé à la fin de la fonction

	// 1ere partie
	time.Sleep(time.Second * time.Duration(s))

	if s == 3 {
		//simule un bug
		result <- false
		return
	}

	str.Mutex.Lock()
	str.val = "test"
	str.Mutex.Unlock()

	result <- true

	allDone.Wait()
	// 2ème partie
	time.Sleep(time.Second * time.Duration(s))

	if s == 1 {
		//simule un bug
		result <- false
		return
	}
	result <- true
}

func main() {

	start := time.Now()

	result := make(chan bool)

	var wg sync.WaitGroup
	var allDone sync.WaitGroup

	allDone.Add(1)

	wg.Add(3)

	maStruct := MaStruct{}

	go block(5, result, &allDone, &maStruct)
	go block(1, result, &allDone, &maStruct)
	go block(3, result, &allDone, &maStruct)

	goroutunes := 3
	success := 0

	for goroutunes != 0 {
		r := <-result
		if r {
			success++
		}
		goroutunes--
	}

	fmt.Println("Success: ", success)
	allDone.Done()

	totalSuccess := 0

	for success != 0 {
		r := <-result
		if r {
			totalSuccess++
		}
		success--
	}

	fmt.Println("Total Success: ", totalSuccess)

	fmt.Printf("Durée d'éxecution: %s\n", time.Since(start))

}
