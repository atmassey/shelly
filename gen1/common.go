package gen1

import (
	"encoding/json"
	"net/http"
)

// Provides basic information about the Shelly device.
func GetShelly(IP string) (*Shelly, error) {
	var shelly Shelly
	req, err := http.NewRequest("GET", "http://"+IP+"/shelly", nil)
	if err != nil {
		return &Shelly{}, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &Shelly{}, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&shelly); err != nil {
		return &Shelly{}, err
	}
	return &shelly, nil
}

// When requested the device will reboot.
func Reboot(IP string) error {
	req, err := http.NewRequest("GET", "http://"+IP+"/reboot", nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// Provides information about the current firmware version and the ability to trigger and over-the-air update.
func GetOta(IP string) (*Ota, error) {
	var ota Ota
	req, err := http.NewRequest("GET", "http://"+IP+"/ota", nil)
	if err != nil {
		return &Ota{}, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &Ota{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return &Ota{}, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&ota); err != nil {
		return &Ota{}, err
	}
	return &ota, nil
}
