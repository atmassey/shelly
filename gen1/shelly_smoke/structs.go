package shelly_smoke

type Status struct {
	IsValid    bool        `json:"is_valid"`
	Smoke      bool        `json:"smoke"`
	Tmp        Temperature `json:"tmp"`
	Bat        Battery     `json:"bat"`
	ActReasons []string    `json:"act_reasons"`
}

type Temperature struct {
	Value        int     `json:"value"`
	Units        string  `json:"units"`
	TemperatureC int     `json:"tC"`
	TemperatureF float64 `json:"tF"`
	IsValid      bool    `json:"is_valid"`
}

type Battery struct {
	Value   int     `json:"value"`
	Voltage float64 `json:"voltage"`
}

type Settings struct {
	Device            DeviceSettings    `json:"device"`
	Actions           ActionSettings    `json:"actions"`
	Sensors           SensorSettings    `json:"sensors"`
	SleepMode         SleepModeSettings `json:"sleep_mode"`
	TemperatureOffset int               `json:"temperature_offset"`
}

type DeviceSettings struct {
	SleepMode bool `json:"sleep_mode"`
}

type ActionSettings struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

type SensorSettings struct {
	TemperatureThreshold int    `json:"temperature_threshold"`
	TemperatureUnit      string `json:"temperature_unit"`
}

type SleepModeSettings struct {
	Period int    `json:"period"`
	Unit   string `json:"unit"`
}

type SettingsActions struct {
	Actions ActionsURL `json:"actions"`
}

type ActionsURL struct {
	FireDetectedURL []URLSetting `json:"fire_detected_url"`
	FireGoneURL     []URLSetting `json:"fire_gone_url"`
}

type URLSetting struct {
	Index   int      `json:"index"`
	URLs    []string `json:"urls"`
	Enabled bool     `json:"enabled"`
}
