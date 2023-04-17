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

func block(s int, wg *sync.WaitGroup, allDone *sync.WaitGroup, str *MaStruct) {
	// defer wg.Done() // Est appelé à la fin de la fonction

	// 1ere partie
	start := time.Now()
	time.Sleep(time.Second * time.Duration(s))

	str.Mutex.Lock()
	str.val = "test"
	str.Mutex.Unlock()

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

	maStruct := MaStruct{}

	go block(5, &wg, &allDone, &maStruct)
	go block(1, &wg, &allDone, &maStruct)
	go block(3, &wg, &allDone, &maStruct)

	wg.Wait()
	wg.Add(3)
	allDone.Done()
	wg.Wait()

	fmt.Printf("Durée d'éxecution: %s\n", time.Since(start))

}
