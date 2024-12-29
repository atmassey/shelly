package shelly_plug

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"net/http"
)

// Status represents the status of the device.
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

func relayCommand(IP string, Command string, Timer string) error {
	if net.ParseIP(IP) == nil {
		return errors.New("invalid IP address")
	}
	req, err := http.NewRequest("GET", "http://"+IP+"/relay/0", nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("turn", Command)
	q.Add("timer", Timer)
	req.URL.RawQuery = q.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to send command, response status: " + resp.Status)
	}
	return nil
}

// RelayOn turns the relay on.
func RelayOn(IP string) error {
	return relayCommand(IP, "on", "0")
}

// RelayOff turns the relay off.
func RelayOff(IP string) error {
	return relayCommand(IP, "off", "0")
}

// RelayToggle toggles the relay.
func RelayToggle(IP string) error {
	return relayCommand(IP, "toggle", "0")
}
