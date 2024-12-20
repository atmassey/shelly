package shelly1_1pm

import "net/http"

func relayCommand(IP string, Command string, Timer string) error {
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
	if resp.StatusCode != http.StatusOK {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func RelayOn(IP string) error {
	return relayCommand(IP, "on", "0")
}

func RelayOff(IP string) error {
	return relayCommand(IP, "off", "0")
}

func RelayToggle(IP string) error {
	return relayCommand(IP, "toggle", "0")
}
