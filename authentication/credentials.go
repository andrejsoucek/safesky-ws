package authentication

import (
	"encoding/json"
)

type Credentials struct {
	Email string `json:"email"`
	Key   string `json:"key"`
}

func CreateCredentialsFromJson(data string) (Credentials, error) {
	x := Credentials{}
	err := json.Unmarshal([]byte(data), &x)
	if err != nil {
		return x, err
	}

	return x, nil
}
