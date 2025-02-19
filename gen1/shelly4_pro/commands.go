package shelly4_pro

import (
	"encoding/json"
	"errors"
	"fmt"
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

func relayCommand(IP string, index int, command string) error {
	if net.ParseIP(IP) == nil {
		return errors.New("invalid IP address")
	}
	req, err := http.NewRequest("POST", "http://"+IP+"/relay/"+fmt.Sprint(index), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	q := req.URL.Query()
	q.Add("turn", command)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to send command, response status: " + resp.Status)
	}
	return nil
}

// TurnOn turns the relay on.
func TurnOn(IP string, index int) error {
	return relayCommand(IP, index, "on")
}

// TurnOff turns the relay off.
func TurnOff(IP string, index int) error {
	return relayCommand(IP, index, "off")
}

// Toggle toggles the relay.
func Toggle(IP string, index int) error {
	return relayCommand(IP, index, "toggle")
}
