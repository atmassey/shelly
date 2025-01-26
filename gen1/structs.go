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

type Device struct {
	Type     string `json:"type"`
	Mac      string `json:"mac"`
	Hostname string `json:"hostname"`
}

type WifiAp struct {
	Enabled bool   `json:"enabled"`
	Ssid    string `json:"ssid"`
	Key     string `json:"key"`
}

type WifiSta struct {
	Enabled    bool    `json:"enabled"`
	Ssid       string  `json:"ssid"`
	Ipv4Method string  `json:"ipv4_method"`
	Ip         *string `json:"ip"`
	Gw         *string `json:"gw"`
	Mask       *string `json:"mask"`
	Dns        *string `json:"dns"`
}

type ApRoaming struct {
	Enabled   bool `json:"enabled"`
	Threshold int  `json:"threshold"`
}

type Mqtt struct {
	Enable              bool   `json:"enable"`
	Server              string `json:"server"`
	User                string `json:"user"`
	Id                  string `json:"id"`
	ReconnectTimeoutMax int    `json:"reconnect_timeout_max"`
	ReconnectTimeoutMin int    `json:"reconnect_timeout_min"`
	CleanSession        bool   `json:"clean_session"`
	KeepAlive           int    `json:"keep_alive"`
	MaxQos              int    `json:"max_qos"`
	Retain              bool   `json:"retain"`
	UpdatePeriod        int    `json:"update_period"`
}

type CoIot struct {
	Enabled      bool   `json:"enabled"`
	UpdatePeriod int    `json:"update_period"`
	Peer         string `json:"peer"`
}

type Sntp struct {
	Server  string `json:"server"`
	Enabled bool   `json:"enabled"`
}

type Login struct {
	Enabled     bool   `json:"enabled"`
	Unprotected bool   `json:"unprotected"`
	Username    string `json:"username"`
}

type BuildInfo struct {
	BuildId        string `json:"build_id"`
	BuildTimestamp string `json:"build_timestamp"`
	BuildVersion   string `json:"build_version"`
}

type Cloud struct {
	Enabled   bool `json:"enabled"`
	Connected bool `json:"connected"`
}

type Settings struct {
	Device                    Device    `json:"device"`
	WifiAp                    WifiAp    `json:"wifi_ap"`
	WifiSta                   WifiSta   `json:"wifi_sta"`
	WifiSta1                  WifiSta   `json:"wifi_sta1"`
	ApRoaming                 ApRoaming `json:"ap_roaming"`
	Mqtt                      Mqtt      `json:"mqtt"`
	CoIot                     CoIot     `json:"coiot"`
	Sntp                      Sntp      `json:"sntp"`
	Login                     Login     `json:"login"`
	PinCode                   string    `json:"pin_code"`
	Name                      string    `json:"name"`
	Fw                        string    `json:"fw"`
	Discoverable              bool      `json:"discoverable"`
	BuildInfo                 BuildInfo `json:"build_info"`
	Cloud                     Cloud     `json:"cloud"`
	Timezone                  string    `json:"timezone"`
	Lat                       float64   `json:"lat"`
	Lng                       float64   `json:"lng"`
	TzAutodetect              bool      `json:"tzautodetect"`
	TzUtcOffset               int       `json:"tz_utc_offset"`
	TzDst                     bool      `json:"tz_dst"`
	TzDstAuto                 bool      `json:"tz_dst_auto"`
	Time                      string    `json:"time"`
	UnixTime                  int       `json:"unixtime"`
	LedStatusDisable          bool      `json:"led_status_disable"`
	DebugEnable               bool      `json:"debug_enable"`
	AllowCrossOrigin          bool      `json:"allow_cross_origin"`
	WifiRecoveryRebootEnabled bool      `json:"wifirecovery_reboot_enabled"`
}
