package safesky

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andrejsoucek/safesky-ws/aircraft"
	"github.com/andrejsoucek/safesky-ws/config"
)

func GetAircrafts(cfg config.Config) ([]aircraft.Aircraft, error) {
	req, err := createRequest(cfg)
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

	return []aircraft.Aircraft{}, errors.New(fmt.Sprintf("Error: Status code %d", resp.StatusCode))
}

func createRequest(cfg config.Config) (*http.Request, error) {
	req, err := http.NewRequest(
		"GET",
		cfg.SafeSkyApiUrl,
		nil,
	)
	if err != nil {
		return &http.Request{}, err
	}

	q := req.URL.Query()
	q.Add("viewport", fmt.Sprintf(
		"%f,%f,%f,%f",
		cfg.SafeSkyBB.SouthWest.Lat,
		cfg.SafeSkyBB.SouthWest.Lng,
		cfg.SafeSkyBB.NorthEast.Lat,
		cfg.SafeSkyBB.NorthEast.Lng,
	))
	q.Add("altitude_max", cfg.SafeSkyMaxAlt)
	req.URL.RawQuery = q.Encode()
	req.Header = http.Header{
		"x-api-Key": []string{cfg.SafeSkyApiKey},
		"origin":    []string{"https://live.safesky.app/"},
	}
	return req, nil
}

func convertResponse(resp [][]interface{}) []aircraft.Aircraft {
	xs := []aircraft.Aircraft{}
	for _, item := range resp {
		xs = append(xs, aircraft.CreateFromResponse(item))
	}

	return xs
}
