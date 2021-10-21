package main

import (
	"fmt"
	"time"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/safesky"
)

func doEvery(d time.Duration, f func() ([]aircraft.Aircraft, error)) {
	for range time.Tick(d) {
		aircrafts, err := f()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(aircrafts)
	}
}

func main() {
	doEvery(4000*time.Millisecond, safesky.GetAircrafts)
}
