package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, target interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err

	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, target); err != nil {
		return err
	}
	return nil

}
