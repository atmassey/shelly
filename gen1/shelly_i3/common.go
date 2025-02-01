package shellyi3

import (
	"encoding/json"
	"errors"
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
