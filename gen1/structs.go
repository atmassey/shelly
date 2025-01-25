package gen1

type Shelly struct {
	Type         string `json:"type"`
	Mac          string `json:"mac"`
	Auth         bool   `json:"auth"`
	Fw           string `json:"fw"`
	LongId       int    `json:"longid"`
	Discoverable bool   `json:"discoverable"`
}

type Ota struct {
	Status      string `json:"status"`
	HasUpdate   bool   `json:"has_update"`
	New_Version string `json:"new_version"`
	Old_Version string `json:"old_version"`
}

type Ota_Check struct {
	Status string `json:"status"`
}

// WifiScan represents the structure of the wifi scan JSON data
type WifiScan struct {
	WifiScan string   `json:"wifiscan"`
	Results  []Result `json:"results"`
}

// Result represents the structure of each result in the wifi scan
type Result struct {
	SSID    string `json:"ssid"`
	Auth    int    `json:"auth"`
	Channel int    `json:"channel"`
	BSSID   string `json:"bssid"`
	RSSI    int    `json:"rssi"`
}
