package shelly_bulb

type Actions struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

type Light struct {
	IsOn          bool          `json:"ison"`
	Red           int           `json:"red"`
	Green         int           `json:"green"`
	Blue          int           `json:"blue"`
	White         int           `json:"white"`
	Gain          int           `json:"gain"`
	Temp          int           `json:"temp"`
	Brightness    int           `json:"brightness"`
	Effect        int           `json:"effect"`
	DefaultState  string        `json:"default_state"`
	AutoOn        int           `json:"auto_on"`
	AutoOff       int           `json:"auto_off"`
	Power         int           `json:"power"`
	Schedule      bool          `json:"schedule"`
	ScheduleRules []interface{} `json:"schedule_rules"`
}

type Settings struct {
	Mode    string  `json:"mode"`
	Actions Actions `json:"actions"`
	Lights  []Light `json:"lights"`
}

type URLAction struct {
	Index   int      `json:"index"`
	URLs    []string `json:"urls"`
	Enabled bool     `json:"enabled"`
}

type SettingsActionsURL struct {
	OutOnURL  []URLAction `json:"out_on_url"`
	OutOffURL []URLAction `json:"out_off_url"`
}

type SettingsActions struct {
	Actions SettingsActionsURL `json:"actions"`
}
