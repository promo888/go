package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan int, 1000)
	for i := 0; i < 500; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		c <- rand.Int()
		fmt.Println(len(c))
		//time.Sleep(time.Millisecond * 50)
	}
}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("обработчик %d получил %d\n", w.id, data)
		//time.Sleep(time.Millisecond * 500)
	}
}
