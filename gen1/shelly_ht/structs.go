package shelly_ht

// Device represents the device configuration
type Device struct {
	SleepMode bool `json:"sleep_mode"`
}

// Actions represents the actions configuration
type Actions struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

// Sensors represents the sensors configuration
type Sensors struct {
	TemperatureThreshold int    `json:"temperature_threshold"`
	TemperatureUnit      string `json:"temperature_unit"`
	HumidityThreshold    int    `json:"humidity_threshold"`
}

// SleepMode represents the sleep mode configuration
type SleepMode struct {
	Period int    `json:"period"`
	Unit   string `json:"unit"`
}

// Configuration represents the overall configuration
type Settings struct {
	Device            Device    `json:"device"`
	Actions           Actions   `json:"actions"`
	Sensors           Sensors   `json:"sensors"`
	SleepMode         SleepMode `json:"sleep_mode"`
	ExternalPower     int       `json:"external_power"`
	TemperatureOffset int       `json:"temperature_offset"`
	HumidityOffset    int       `json:"humidity_offset"`
}

// Action represents a generic action configuration
type Action struct {
	Index   int      `json:"index"`
	Urls    []string `json:"urls"`
	Enabled bool     `json:"enabled"`
}

// TempOverAction represents the configuration for temp_over_url actions
type TempOverAction struct {
	Action
	TempOverValue   float64 `json:"temp_over_value"`
	TempOverOnetime bool    `json:"temp_over_onetime"`
}

// TempUnderAction represents the configuration for temp_under_url actions
type TempUnderAction struct {
	Action
	TempUnderValue   float64 `json:"temp_under_value"`
	TempUnderOnetime bool    `json:"temp_under_onetime"`
}

// HumOverAction represents the configuration for hum_over_url actions
type HumOverAction struct {
	Action
	HumOverValue   int  `json:"hum_over_value"`
	HumOverOnetime bool `json:"hum_over_onetime"`
}

// HumUnderAction represents the configuration for hum_under_url actions
type HumUnderAction struct {
	Action
	HumUnderValue   int  `json:"hum_under_value"`
	HumUnderOnetime bool `json:"hum_under_onetime"`
}

// Actions represents the overall actions configuration
type ActionsSettings struct {
	ReportURL    []Action          `json:"report_url"`
	TempOverURL  []TempOverAction  `json:"temp_over_url"`
	TempUnderURL []TempUnderAction `json:"temp_under_url"`
	HumOverURL   []HumOverAction   `json:"hum_over_url"`
	HumUnderURL  []HumUnderAction  `json:"hum_under_url"`
}

// Tmp represents the temperature details
type Tmp struct {
	Value   int     `json:"value"`
	Units   string  `json:"units"`
	TC      int     `json:"tC"`
	TF      float64 `json:"tF"`
	IsValid bool    `json:"is_valid"`
}

// Hum represents the humidity details
type Hum struct {
	Value   int  `json:"value"`
	IsValid bool `json:"is_valid"`
}

// Bat represents the battery details
type Bat struct {
	Value   int     `json:"value"`
	Voltage float64 `json:"voltage"`
}

// Status represents the overall configuration
type Status struct {
	IsValid        bool     `json:"is_valid"`
	Tmp            Tmp      `json:"tmp"`
	Hum            Hum      `json:"hum"`
	Bat            Bat      `json:"bat"`
	ActReasons     []string `json:"act_reasons"`
	ConnectRetries int      `json:"connect_retries"`
}
