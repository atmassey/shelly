package shelly_smoke

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
)

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

func GetSettings(IP string) (*Settings, error) {
	if net.ParseIP(IP) == nil {
		return nil, errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/settings", nil)
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

func GetSettingsActions(IP string) (*SettingsActions, error) {
	if net.ParseIP(IP) == nil {
		return nil, errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/settings/actions", nil)
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
		return nil, errors.New("failed to get settings actions, response status: " + resp.Status)
	}
	var settingsActions SettingsActions
	if err := json.NewDecoder(resp.Body).Decode(&settingsActions); err != nil {
		return nil, err
	}
	return &settingsActions, nil
}
