package shelly_flood

// Device structure
type Device struct {
	SleepMode bool `json:"sleep_mode"`
}

// Actions structure
type Actions struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

// Sensors structure
type Sensors struct {
	TemperatureThreshold int    `json:"temperature_threshold"`
	TemperatureUnit      string `json:"temperature_unit"`
}

// SleepMode structure
type SleepMode struct {
	Period int    `json:"period"`
	Unit   string `json:"unit"`
}

// Settings structure
type Settings struct {
	Device            Device    `json:"device"`
	Actions           Actions   `json:"actions"`
	Sensors           Sensors   `json:"sensors"`
	SleepMode         SleepMode `json:"sleep_mode"`
	RainSensor        bool      `json:"rain_sensor"`
	TemperatureOffset int       `json:"temperature_offset"`
}

// Tmp structure
type Tmp struct {
	Value   int    `json:"value"`
	Units   string `json:"units"`
	TC      int    `json:"tC"`
	TF      int    `json:"tF"`
	IsValid bool   `json:"is_valid"`
}

// Bat structure
type Bat struct {
	Value   int     `json:"value"`
	Voltage float64 `json:"voltage"`
}

// Main Settings structure
type Status struct {
	IsValid    bool     `json:"is_valid"`
	Flood      bool     `json:"flood"`
	Tmp        Tmp      `json:"tmp"`
	Bat        Bat      `json:"bat"`
	ActReasons []string `json:"act_reasons"`
	RainSensor bool     `json:"rain_sensor"`
}

// URLAction structure
type URLAction struct {
	Index   int      `json:"index"`
	URLs    []string `json:"urls"`
	Enabled bool     `json:"enabled"`
}

// TempOverURL structure
type TempOverURL struct {
	Index           int      `json:"index"`
	URLs            []string `json:"urls"`
	Enabled         bool     `json:"enabled"`
	TempOverValue   float64  `json:"temp_over_value"`
	TempOverOnetime bool     `json:"temp_over_onetime"`
}

// TempUnderURL structure
type TempUnderURL struct {
	Index            int      `json:"index"`
	URLs             []string `json:"urls"`
	Enabled          bool     `json:"enabled"`
	TempUnderValue   float64  `json:"temp_under_value"`
	TempUnderOnetime bool     `json:"temp_under_onetime"`
}

// Actions structure
type Setting_Actions_Slices struct {
	ReportURL        []URLAction    `json:"report_url"`
	FloodDetectedURL []URLAction    `json:"flood_detected_url"`
	FloodGoneURL     []URLAction    `json:"flood_gone_url"`
	TempOverURL      []TempOverURL  `json:"temp_over_url"`
	TempUnderURL     []TempUnderURL `json:"temp_under_url"`
}

// SettingsActions structure
type SettingsActions struct {
	Actions Actions `json:"actions"`
}
