package shelly_3em

type Settings struct {
	Actions struct {
		Active bool     `json:"active"`
		Names  []string `json:"names"`
	} `json:"actions"`
	Relays []struct {
		Name          *string  `json:"name"`
		IsOn          bool     `json:"ison"`
		HasTimer      bool     `json:"has_timer"`
		OverPower     bool     `json:"overpower"`
		DefaultState  string   `json:"default_state"`
		AutoOn        int      `json:"auto_on"`
		AutoOff       int      `json:"auto_off"`
		Schedule      bool     `json:"schedule"`
		ScheduleRules []string `json:"schedule_rules"`
	} `json:"relays"`
	Emeters []struct {
		ApplianceType string `json:"appliance_type"`
		MaxPower      int    `json:"max_power"`
	} `json:"emeters"`
	LedStatusDisable bool `json:"led_status_disable"`
}

type SettingActions struct {
	Actions struct {
		OutOnURL []struct {
			Index   int      `json:"index"`
			URLs    []string `json:"urls"`
			Enabled bool     `json:"enabled"`
		} `json:"out_on_url"`
		OutOffURL []struct {
			Index   int      `json:"index"`
			URLs    []string `json:"urls"`
			Enabled bool     `json:"enabled"`
		} `json:"out_off_url"`
		OverPowerURL []struct {
			Index                 int      `json:"index"`
			URLs                  []string `json:"urls"`
			Enabled               bool     `json:"enabled"`
			OverPowerURLThreshold int      `json:"over_power_url_threshold"`
		} `json:"over_power_url"`
		UnderPowerURL []struct {
			Index                  int      `json:"index"`
			URLs                   []string `json:"urls"`
			Enabled                bool     `json:"enabled"`
			UnderPowerURLThreshold int      `json:"under_power_url_threshold"`
		} `json:"under_power_url"`
	} `json:"actions"`
}

type Relay struct {
	Name          *string  `json:"name"`
	IsOn          bool     `json:"ison"`
	HasTimer      bool     `json:"has_timer"`
	OverPower     bool     `json:"overpower"`
	DefaultState  string   `json:"default_state"`
	AutoOn        int      `json:"auto_on"`
	AutoOff       int      `json:"auto_off"`
	Schedule      bool     `json:"schedule"`
	ScheduleRules []string `json:"schedule_rules"`
}

type Status struct {
	Relays []struct {
		IsOn           bool   `json:"ison"`
		HasTimer       bool   `json:"has_timer"`
		TimerStarted   int    `json:"timer_started"`
		TimerDuration  int    `json:"timer_duration"`
		TimerRemaining int    `json:"timer_remaining"`
		OverPower      bool   `json:"overpower"`
		IsValid        bool   `json:"is_valid"`
		Source         string `json:"source"`
	} `json:"relays"`
	Emeters []struct {
		Power         float64 `json:"power"`
		Pf            float64 `json:"pf"`
		Current       float64 `json:"current"`
		Voltage       float64 `json:"voltage"`
		IsValid       bool    `json:"is_valid"`
		Total         float64 `json:"total"`
		TotalReturned float64 `json:"total_returned"`
	} `json:"emeters"`
	TotalPower float64 `json:"total_power"`
	FsMounted  bool    `json:"fs_mounted"`
}

type Emeter struct {
	Power         float64 `json:"power"`
	Pf            float64 `json:"pf"`
	Current       float64 `json:"current"`
	Voltage       float64 `json:"voltage"`
	IsValid       bool    `json:"is_valid"`
	Total         float64 `json:"total"`
	TotalReturned float64 `json:"total_returned"`
}
