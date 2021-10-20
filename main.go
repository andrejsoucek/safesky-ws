package main

import (
	"time"

	"github.com/andrejsoucek/safesky-ws/safesky"
)

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(4000*time.Millisecond, safesky.GetAircrafts)
}
