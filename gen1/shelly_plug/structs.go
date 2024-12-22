package shelly_plug

// Action represents the actions within the JSON
type Action struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

// Relay represents each relay within the JSON
type RelaySettings struct {
	Name          *string  `json:"name"`
	ApplianceType string   `json:"appliance_type"`
	Ison          bool     `json:"ison"`
	HasTimer      bool     `json:"has_timer"`
	Overpower     bool     `json:"overpower"`
	DefaultState  string   `json:"default_state"`
	AutoOn        int      `json:"auto_on"`
	AutoOff       int      `json:"auto_off"`
	Schedule      bool     `json:"schedule"`
	ScheduleRules []string `json:"schedule_rules"`
	MaxPower      int      `json:"max_power"`
}

// Settings represents the entire JSON structure
type Settings struct {
	MaxPower        int             `json:"max_power"`
	LedPowerDisable bool            `json:"led_power_disable"`
	Actions         Action          `json:"actions"`
	Relays          []RelaySettings `json:"relays"`
}

type Relay struct {
	IsOn           bool   `json:"ison"`
	HasTimer       bool   `json:"has_timer"`
	TimerStarted   bool   `json:"timer_started"`
	TimerDuration  int    `json:"timer_duration"`
	TimerRemaining int    `json:"timer_remaining"`
	Overpower      bool   `json:"overpower"`
	Source         string `json:"source"`
}
