package shelly_ht

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

// Returns the status of the device in JSON format
func GetStatus(IP string) (*Status, error) {
	if net.ParseIP(IP) == nil {
		return nil, fmt.Errorf("invalid IP address: %s", IP)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/status", IP), nil)
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

// Returns the settings of the device in JSON format
func GetSettings(IP string) (*Settings, error) {
	if net.ParseIP(IP) == nil {
		return nil, fmt.Errorf("invalid IP address: %s", IP)
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/settings", IP), nil)
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
		return nil, errors.New("failed to get settings, response status: " + resp.Status)
	}
	var settings Settings
	if err := json.NewDecoder(resp.Body).Decode(&settings); err != nil {
		return nil, err
	}
	return &settings, nil
}
