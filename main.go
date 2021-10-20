package main

import (
	"time"

	"github.com/andrejsoucek/safesky-ws/safesky"
)

func doEvery(d time.Duration, f func(*[][]interface{}), aircrafts *[][]interface{}) {
	for range time.Tick(d) {
		f(aircrafts)
	}
}

func main() {
	var aircrafts [][]interface{}
	doEvery(4000*time.Millisecond, safesky.GetAircrafts, &aircrafts)
}
