package gen1

import (
	"encoding/json"
	"net/http"
)

func GetShelly(IP string) (Shelly, error) {
	var shelly Shelly
	req, err := http.NewRequest("GET", "http://"+IP+"/shelly", nil)
	if err != nil {
		return Shelly{}, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Shelly{}, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&shelly); err != nil {
		return Shelly{}, err
	}
	return shelly, nil
}
