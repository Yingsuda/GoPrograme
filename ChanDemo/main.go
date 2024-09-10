package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
}

type Bus struct {
	realList chan *Value
}

var bus *Bus

func init() {
	bus = new(Bus)
	bus.realList = make(chan *Value, 1000)
}

func GetChan() chan *Value {
	return bus.realList
}

var hisChan = make(chan int, 100)

var once sync.Once

func main() {
	go func() {
		ticker := time.NewTicker(time.Second * 2)

		for {
			select {
			case <-ticker.C:
				once.Do(
					func() {
						fmt.Println("rebuild")
						bus.realList = make(chan *Value, 100)
					},
				)
				fmt.Println("Len:", len(bus.realList))
				bus.realList <- &Value{}
			default:
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	for {
	LOOP:
		select {
		case v := <-GetChan():
			fmt.Println("val:", v)
		case his := <-hisChan:
			fmt.Println("his:", his)
			goto LOOP
			//default:
		}
	}
}
