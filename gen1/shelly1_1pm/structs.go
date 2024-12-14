package shelly1_1pm

type Status struct {
	Relays             []Relay           `json:"relays"`
	Meters             []Meter           `json:"meters"`
	Inputs             []Input           `json:"inputs"`
	Ext_Sensors        []Ext_Sensor      `json:"ext_sensors"`
	Ext_Temperature    []Ext_Temperature `json:"ext_temperature"`
	Ext_Humidity       []Ext_Humidity    `json:"ext_humidity"`
	Temperature        float64           `json:"temperature"`
	Overtemperature    bool              `json:"overtemperature"`
	Tmp                Temp              `json:"tmp"`
	Temperature_Status string            `json:"temperature_status"`
	Ext_Switch         []Ext_Switch      `json:"ext_switch"`
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

type Meter struct {
	Power     int     `json:"power"`
	Overpower float64 `json:"overpower"`
	Is_Valid  bool    `json:"is_valid"`
	Timestamp int     `json:"timestamp"`
	Counters  []int   `json:"counters"`
	Total     int     `json:"total"`
}

type Input struct {
	Input     int    `json:"input"`
	Event     string `json:"event"`
	Event_Cnt int    `json:"event_cnt"`
}

// TODO: Find this structure
type Ext_Sensor struct {
}

// TODO: Find this structure
type Ext_Temperature struct {
}

// TODO: Find this structure
type Ext_Humidity struct {
}

type Temp struct {
	TemperatureC float64 `json:"tC"`
	TemperatureF float64 `json:"tF"`
	Is_Valid     bool    `json:"is_valid"`
}

type Ext_Switch struct {
	Zero InputInt `json:"0"`
}

type InputInt struct {
	Input int `json:"input"`
}
