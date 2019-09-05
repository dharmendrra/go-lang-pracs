// package main

// import (
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"
// )

// func main() {
// 	log.Println("practicing routines and channels")
// 	start := time.Now()
// 	ch := make(chan int)
// 	result := make(chan string)
// 	iteration := 10

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go func(ch chan int, result chan string, wg *sync.WaitGroup) {
// 		for c := range ch {
// 			time.Sleep(time.Second)
// 			result <- fmt.Sprintf("routine#1 %d", c)
// 		}
// 		wg.Done()
// 	}(ch, result, &wg)

// 	wg.Add(1)
// 	go func(ch chan int, result chan string, wg *sync.WaitGroup) {
// 		for c := range ch {
// 			time.Sleep(time.Second)
// 			result <- fmt.Sprintf("routine#2 %d", c)
// 		}
// 		wg.Done()
// 	}(ch, result, &wg)

// 	wg.Add(1)
// 	go func(wg *sync.WaitGroup) {
// 		for i := 1; i <= iteration; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 		wg.Done()
// 	}(&wg)

// 	// wg.Add(1)
// 	// go func(wg *sync.WaitGroup) {
// 	// 	for j := 1; j <= iteration; j++ {
// 	// 		r := <-result
// 	// 		fmt.Println(r)
// 	// 	}
// 	// 	wg.Done()
// 	// }(&wg)

// 	// wg.Add(1)
// 	// go func(wg *sync.WaitGroup) {
// 	// 	r, ok := <-result
// 	// 	fmt.Println(r, ok)
// 	// 	if !ok {
// 	// 		wg.Done()
// 	// 	}
// 	// }(&wg)

// 	go func() {
// 		wg.Wait()
// 		close(result)
// 	}()

// 	for r := range result {
// 		fmt.Println(r)
// 	}
// 	//wg.Wait()
// 	log.Println(time.Now().Sub(start))

// 	// for {
// 	// 	select {
// 	// 	case r, ok := <-result:
// 	// 		if ok {
// 	// 			fmt.Println(r)
// 	// 		} else {
// 	// 			//fmt.Println("nok")
// 	// 		}
// 	// 	case _, ok :
// 	// 		fmt.Println("def")
// 	// 		close(result)
// 	// 		break
// 	// 	}
// 	// }

// 	//fmt.Println("exit")
// }
