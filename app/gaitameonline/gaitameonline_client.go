package gaitameonline

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GaitameOnlineClient interface {
	GetCurrentExchangeRate() (*ExchangeRate, error)
}

type GaitameOnlineClientImpl struct {
	Url string
}

func (g GaitameOnlineClientImpl) GetCurrentExchangeRate() (*ExchangeRate, error) {
	resp, err := http.Get(g.Url)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		return nil, fmt.Errorf("failed request to gaitameonline(%s): %s", g.Url, err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read all from gaitameonline: %s", err)
	}

	var response ExchangeRate
	if err := json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("failed parse response body json from gaitameonline: %s", err)
	}

	return &response, nil
}
