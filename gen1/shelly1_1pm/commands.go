package shelly1_1pm

import (
	"errors"
	"log"
	"net"
	"net/http"
)

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
