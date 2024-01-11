package counter

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ON = iota
	OFF
)

type Counter struct {
	status chan int
	handle func(int, int)
}

func (counter *Counter) runDevice(interrupt chan struct{}) {
	data := 0
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	data = rng.Intn(30) + 1

	sun := 0
	sun = rng.Intn(30) + 1

	for {
		select {
		case <-interrupt:
			counter.handle(0, 0)
			return
		default:
			data++
			sun++
			counter.handle(data, sun)
			fmt.Println("Temperature value:", data)
			time.Sleep(1 * time.Second)
		}
	}
}

func (counter *Counter) initDevice() {
	interrupt := make(chan struct{})

	for {
		select {
		case status := <-counter.status:
			if status == ON {
				go counter.runDevice(interrupt)
			}
			if status == OFF {
				interrupt <- struct{}{}
			}
		}
	}
}

func (counter *Counter) TurnOn() {
	counter.status <- ON
}

func (counter *Counter) TurnOff() {
	counter.status <- OFF
}

func NewCounter(h func(x, y int)) *Counter {
	counter := &Counter{
		status: make(chan int),
		handle: h,
	}

	go counter.initDevice()

	return counter
}

func CloseCounter(counter *Counter) {
	close(counter.status)
}
