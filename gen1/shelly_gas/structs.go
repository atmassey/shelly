package shelly_gas

type OK struct {
	Ok bool `json:"ok"`
}

type Status struct {
	Gas_Sensor    Gas_Sensor    `json:"gas_sensor"`
	Concentration Concentration `json:"concentration"`
	Valves        []States      `json:"valves"`
}

type Gas_Sensor struct {
	State      string `json:"state"`
	SelfTest   string `json:"self_test_state"`
	AlarmState string `json:"alarm_state"`
}

type Concentration struct {
	PPM     int  `json:"ppm"`
	IsValid bool `json:"is_valid"`
}

type States struct {
	State string `json:"state"`
}

type URLAction struct {
	Index   int      `json:"index"`
	URLs    []string `json:"urls"`
	Enabled bool     `json:"enabled"`
}

type Actions struct {
	AlarmOffURL   []URLAction `json:"alarm_off_url"`
	AlarmMildURL  []URLAction `json:"alarm_mild_url"`
	AlarmHeavyURL []URLAction `json:"alarm_heavy_url"`
}

type SettingsActions struct {
	Actions Actions `json:"actions"`
}

type Settings_Actions struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

type Settings struct {
	SetVolume int              `json:"set_volume"`
	Actions   Settings_Actions `json:"actions"`
}
