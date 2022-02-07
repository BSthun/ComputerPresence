package utils

import (
	"reflect"
	"time"
)

func execute(p interface{}, ticker *time.Ticker, stop chan bool) {
	for {
		select {
		case <-stop:
			println("STOP INTERVAL")
			return
		case <-ticker.C:
			reflect.ValueOf(p).Call([]reflect.Value{})
		}
	}
}

func SetInterval(p interface{}, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	stop := make(chan bool)
	go execute(p, ticker, stop)

	return stop
}
