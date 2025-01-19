package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/GauravC4/pokedexcli/internal/pokecache"
)

func Http_get(address string, respPtr any, cachePtr *pokecache.Cache) error {
	if _, err := url.ParseRequestURI(address); err != nil {
		return fmt.Errorf("invalid url : %v", address)
	}

	var body []byte
	cacheHit := false

	if cachePtr != nil {
		body, cacheHit = cachePtr.Get(address)
	}
	if !cacheHit {
		res, err := http.Get(address)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code %v and body %s", res.StatusCode, body)
		}
		cachePtr.Add(address, body)
	}

	err := json.Unmarshal(body, respPtr)
	if err != nil {
		return err
	}
	return nil
}
