package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Http_get(address string, respPtr any) error {
	if _, err := url.ParseRequestURI(address); err != nil {
		return fmt.Errorf("invalid url : %v", address)
	}

	res, err := http.Get(address)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code %v and body %s", res.StatusCode, body)
	}

	err = json.Unmarshal(body, respPtr)
	if err != nil {
		return err
	}
	return nil
}
