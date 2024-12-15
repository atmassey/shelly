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
