package shellyi3

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Provides basic settings information about the Shelly i3 device.
func GetSettings(IP string) (*Settings, error) {
	var settings Settings
	req, err := http.NewRequest("GET", "http://"+IP+"/settings", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get settings, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

// Provides information about the inputs of the Shelly i3 device.
func GetInputStatus(IP string, index int) (*Input, error) {
	var inputs Input
	req, err := http.NewRequest("GET", "http://"+IP+"/settings/inputs/"+fmt.Sprint(index), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get inputs, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&inputs); err != nil {
		return nil, err
	}
	return &inputs, nil
}

// Provides information about a specific input channel of the Shelly i3 device.
func GetInputChannelStatus(IP string, index int) (*InputChannel, error) {
	var input InputChannel
	req, err := http.NewRequest("GET", "http://"+IP+"/input/"+fmt.Sprint(index), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get inputs, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&input); err != nil {
		return nil, err
	}
	return &input, nil
}
