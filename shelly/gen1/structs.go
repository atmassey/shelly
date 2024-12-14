package gen1

type Shelly struct {
	Type         string `json:"type"`
	Mac          string `json:"mac"`
	Auth         bool   `json:"auth"`
	Fw           string `json:"fw"`
	LongId       int    `json:"longid"`
	Discoverable bool   `json:"discoverable"`
}
