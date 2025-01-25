package gen1

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Provides basic information about the Shelly device.
func GetShelly(IP string) (*Shelly, error) {
	var shelly Shelly
	req, err := http.NewRequest("GET", "http://"+IP+"/shelly", nil)
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
		return nil, errors.New("failed to get shelly, response status: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&shelly); err != nil {
		return nil, err
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
		return errors.New("failed to reboot, response status: " + resp.Status)
	}
	defer resp.Body.Close()
	return nil
}

// Provides information about the current firmware version and the ability to trigger and over-the-air update.
func GetOta(IP string) (*Ota, error) {
	var ota Ota
	req, err := http.NewRequest("GET", "http://"+IP+"/ota", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get ota, response status: " + resp.Status)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&ota); err != nil {
		return nil, err
	}
	return &ota, nil
}

// Checks for a new firmware version.
func CheckOta(IP string) (*Ota_Check, error) {
	var otaCheck Ota_Check
	req, err := http.NewRequest("GET", "http://"+IP+"/ota/check", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to check ota, response status: " + resp.Status)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&otaCheck); err != nil {
		return nil, err
	}
	return &otaCheck, nil
}

// Reset STA cache: only applicable on devices which support STA caching
// (battery operated devices Shelly Door/Window 1/2, Shelly H&T, Shelly Smoke, Shelly Flood, Shelly Button1)
func StaCacheReset(IP string) error {
	req, err := http.NewRequest("GET", "http://"+IP+"/sta_cache_reset", nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to reset sta cache, response status: " + resp.Status)
	}
	defer resp.Body.Close()
	return nil
}

// Only available when in AP mode. Starts a WiFi scan and provides information about found networks.
func GetWifiScan(IP string) (*WifiScan, error) {
	var wifiScan WifiScan
	req, err := http.NewRequest("GET", "http://"+IP+"/wifiscan", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get wifi scan, response status: " + resp.Status)
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&wifiScan); err != nil {
		return nil, err
	}
	return &wifiScan, nil
}
