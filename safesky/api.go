package safesky

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/geography"
)

func GetAircrafts() ([]aircraft.Aircraft, error) {
	req, err := createRequest()
	if err != nil {
		return []aircraft.Aircraft{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []aircraft.Aircraft{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []aircraft.Aircraft{}, err
		}

		var target [][]interface{}
		err = json.Unmarshal([]byte(body), &target)
		if err != nil {
			return []aircraft.Aircraft{}, err
		}

		return convertResponse(target), nil
	}

	return []aircraft.Aircraft{}, nil
}

func createRequest() (*http.Request, error) {
	req, err := http.NewRequest(
		"GET",
		"https://public-api.safesky.app/v1/beacons",
		nil,
	)
	if err != nil {
		return &http.Request{}, err
	}

	sw := geography.LatLon{Lat: 47.739323, Lon: 11.985945}
	ne := geography.LatLon{Lat: 51.079371, Lon: 22.585201}

	apiKey := os.Getenv("API_KEY")
	q := req.URL.Query()
	q.Add("viewport", fmt.Sprintf("%f,%f,%f,%f", sw.Lat, sw.Lon, ne.Lat, ne.Lon))
	q.Add("altitude_max", "1829")
	req.URL.RawQuery = q.Encode()
	req.Header = http.Header{
		"x-api-Key": []string{apiKey},
		"origin":    []string{"https://live.safesky.app/"},
	}

	return req, nil
}

func convertResponse(resp [][]interface{}) []aircraft.Aircraft {
	var xs []aircraft.Aircraft
	for _, item := range resp {
		xs = append(xs, aircraft.CreateFromResponse(item))
	}

	return xs
}
