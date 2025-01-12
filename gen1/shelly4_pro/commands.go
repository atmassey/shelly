package shelly4_pro

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
)

func GetSettings(IP string) (*Device, error) {
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
		err := resp.Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get settings, response status: " + resp.Status)
	}
	var settings Device
	err = json.NewDecoder(resp.Body).Decode(&settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func GetStatus(IP string) (*RelayStatus, error) {
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
		err := resp.Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to get status, response status: " + resp.Status)
	}
	var status RelayStatus
	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}
