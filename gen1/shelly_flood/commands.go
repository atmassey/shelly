package shelly_flood

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
)

// returns the status of the shelly flood device
func GetStatus(IP string) (*Status, error) {
	if net.ParseIP(IP) == nil {
		return nil, errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/status", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get status, response status: " + resp.Status)
	}
	var status Status
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}
	return &status, nil
}
