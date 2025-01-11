package shelly4_pro

type Relay struct {
	Name          *string       `json:"name"`
	IsOn          bool          `json:"ison"`
	HasTimer      bool          `json:"has_timer"`
	Overpower     bool          `json:"overpower"`
	DefaultState  string        `json:"default_state"`
	BtnType       string        `json:"btn_type"`
	AutoOn        int           `json:"auto_on"`
	AutoOff       int           `json:"auto_off"`
	MaxPower      int           `json:"max_power"`
	Schedule      bool          `json:"schedule"`
	ScheduleRules []interface{} `json:"schedule_rules"` // Assuming schedule rules can be of any type
}

type Meter struct {
	Power     int   `json:"power"`
	IsValid   bool  `json:"is_valid"`
	Timestamp int   `json:"timestamp"`
	Counters  []int `json:"counters"`
	Total     int   `json:"total"`
}

type Device struct {
	Relays []Relay `json:"relays"`
	Meters []Meter `json:"meters"`
}
