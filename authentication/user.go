package authentication

import (
	"encoding/json"
)

type User struct {
	SessionId string `json:"session"`
	UserId    int    `json:"userId"`
	Premium   bool   `json:"premium"`
}

func CreateUserFromJson(data []byte) (User, error) {
	x := User{}
	err := json.Unmarshal(data, &x)
	if err != nil {
		return x, err
	}

	return x, nil
}
