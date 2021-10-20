package safesky

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/andrejsoucek/safesky-ws/geography"
)

func GetAircrafts(target *[][]interface{}) {
	req, err := http.NewRequest(
		"GET",
		"https://public-api.safesky.app/v1/beacons",
		nil,
	)
	if err != nil {
		panic(err)
	}

	sw := geography.LatLng{Lat: 47.739323, Lon: 11.985945}
	ne := geography.LatLng{Lat: 51.079371, Lon: 22.585201}

	apiKey := os.Getenv("API_KEY")
	q := req.URL.Query()
	q.Add("viewport", fmt.Sprintf("%f,%f,%f,%f", sw.Lat, sw.Lon, ne.Lat, ne.Lon))
	q.Add("altitude_max", "1829")
	req.URL.RawQuery = q.Encode()
	req.Header = http.Header{
		"x-api-Key": []string{apiKey},
		"origin":    []string{"https://live.safesky.app/"},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		error := json.Unmarshal([]byte(body), target)
		if err != nil {
			panic(error)
		}
	}
}
