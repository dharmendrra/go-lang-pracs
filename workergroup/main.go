package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("practicing routines and channels")
	start := time.Now()
	ch := make(chan int)
	result := make(chan string)
	iteration := 20

	go func() {
		defer close(ch)
		for i := 1; i <= iteration; i++ {
			ch <- i
		}
	}()
	setUpWorkers(ch, result)
	reader(result)
	//time.Sleep(time.Second * 100)
	log.Println(time.Now().Sub(start))
}

func setUpWorkers(ch chan int, result chan string) {
	wg := new(sync.WaitGroup)
	num := 10
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go func(i int) {
			worker(ch, result, i)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
}

func worker(ch chan int, result chan string, routineID int) {
	for c := range ch {
		time.Sleep(time.Millisecond * 10)
		result <- fmt.Sprintf("routine#%d num:%d", routineID, c)
	}
}

func reader(result chan string) {
	for r := range result {
		fmt.Println(r)
	}
}
