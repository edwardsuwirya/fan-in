package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	chProducer1 := make(chan int)
	chProducer2 := make(chan int)
	chConsumer := make(chan int)

	go producers(chProducer1, "A")
	go producers(chProducer2, "B")
	go singleConsumer(chConsumer)

	fanIn(chConsumer, chProducer1, chProducer2)
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
}

func producers(ch chan<- int, name string) {
	for {
		sleep()

		n := rand.Intn(100)

		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}
func singleConsumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}
func fanIn(chC chan<- int, ch ...<-chan int) {
	//var n int
	var wg sync.WaitGroup
	wg.Add(len(ch))
	for _, c := range ch {

		go func(newCh <-chan int) {
			for {
				chC <- <-newCh
			}
		}(c)
	}
	wg.Wait()
	//select {
	//case n = <-chA:
	//	chC <- n
	//case n = <-chB:
	//	chC <- n
	//}

}
