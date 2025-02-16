package shelly_3em

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetSettings returns the settings of the device.
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

// GetSettingsActions returns the action settings of the device.
func GetSettingActions(IP string) (*SettingActions, error) {
	var settingActions SettingActions
	req, err := http.NewRequest("GET", "http://"+IP+"/settings/actions", nil)
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
		return nil, errors.New("failed to get setting actions, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&settingActions); err != nil {
		return nil, err
	}
	return &settingActions, nil
}

// GetRelays returns the relays of the device.
func GetRelay(IP string, index int) (*Relay, error) {
	var relay Relay
	req, err := http.NewRequest("GET", "http://"+IP+"/relay/"+fmt.Sprint(index), nil)
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
		return nil, errors.New("failed to get relay, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&relay); err != nil {
		return nil, err
	}
	return &relay, nil
}

// GetStatus returns the status of the device.
func GetStatus(IP string) (*Status, error) {
	var status Status
	req, err := http.NewRequest("GET", "http://"+IP+"/status", nil)
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
		return nil, errors.New("failed to get status, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}
	return &status, nil
}

// GetEmeter returns the emeter of the device.
func GetEmeter(IP string, index int) (*Emeter, error) {
	var emeter Emeter
	req, err := http.NewRequest("GET", "http://"+IP+"/emeter/"+fmt.Sprint(index), nil)
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
		return nil, errors.New("failed to get emeter, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&emeter); err != nil {
		return nil, err
	}
	return &emeter, nil
}
