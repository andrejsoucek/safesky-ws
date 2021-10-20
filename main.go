package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type LatLng struct {
	latitude  float64
	longitude float64
}

func getAircrafts() {
	req, err := http.NewRequest(
		"GET",
		"https://public-api.safesky.app/v1/beacons",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	sw := LatLng{47.739323, 11.985945}
	ne := LatLng{51.079371, 22.585201}
	apiKey := os.Getenv("API_KEY")
	q := req.URL.Query()
	q.Add("viewport", fmt.Sprintf("%f,%f,%f,%f", sw.latitude, sw.longitude, ne.latitude, ne.longitude))
	q.Add("altitude_max", "1829")
	req.URL.RawQuery = q.Encode()
	req.Header = http.Header{
		"x-api-Key": []string{apiKey},
		"origin":    []string{"https://live.safesky.app/"},
	}

	fmt.Println(req.URL.String())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func doEvery(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}

func main() {
	doEvery(4000*time.Millisecond, getAircrafts)
}
