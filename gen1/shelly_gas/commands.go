package shelly_gas

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
)

// Mute the shelly gas device
func Mute(IP string) (*OK, error) {
	if net.ParseIP(IP) == nil {
		return nil, errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/mute", nil)
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
		return nil, errors.New("failed to mute, response status: " + resp.Status)
	}
	var ok OK
	if err := json.NewDecoder(resp.Body).Decode(&ok); err != nil {
		return nil, err
	}
	return &ok, nil
}

// Unmute the shelly gas device
func Unmute(IP string) (*OK, error) {
	if net.ParseIP(IP) == nil {
		return nil, errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/unmute", nil)
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
		return nil, errors.New("failed to unmute, response status: " + resp.Status)
	}
	var ok OK
	if err := json.NewDecoder(resp.Body).Decode(&ok); err != nil {
		return nil, err
	}
	return &ok, nil
}

// SelfTest the shelly gas device
func SelfTest(IP string) (*OK, error) {
	if net.ParseIP(IP) == nil {
		return nil, errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/self_test", nil)
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
		return nil, errors.New("failed to self test, response status: " + resp.Status)
	}
	var ok OK
	if err := json.NewDecoder(resp.Body).Decode(&ok); err != nil {
		return nil, err
	}
	return &ok, nil
}

// GetStatus returns the status of the device in JSON format
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

// GetSettings returns the settings of the device in JSON format
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
