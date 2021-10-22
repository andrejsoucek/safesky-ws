package authentication

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Authenticate(c Credentials) (User, error) {
	req, err := createRequest(c.Email, c.Key)
	if err != nil {
		return User{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return User{}, err
		}

		user, err := CreateUserFromJson(body)
		if err != nil {
			return User{}, err
		}

		return user, nil
	}

	return User{}, errors.New(fmt.Sprintf("Error: %d", resp.StatusCode))
}

func createRequest(email string, key string) (*http.Request, error) {
	data := url.Values{}
	data.Add("email", email)
	data.Add("key", key)
	req, err := http.NewRequest(
		"POST",
		"https://pocketpilot.cz/api/v1/login",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
