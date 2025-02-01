package shellyi3

type Duration struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Actions struct {
	Active bool     `json:"active"`
	Names  []string `json:"names"`
}

type Input struct {
	Name       *string `json:"name"`
	BtnType    string  `json:"btn_type"`
	BtnReverse int     `json:"btn_reverse"`
}

type Settings struct {
	LongPushDurationMs           Duration `json:"longpush_duration_ms"`
	MultiPushTimeBetweenPushesMs struct {
		Max int `json:"max"`
	} `json:"multipush_time_between_pushes_ms"`
	Actions          Actions `json:"actions"`
	Inputs           []Input `json:"inputs"`
	LedStatusDisable bool    `json:"led_status_disable"`
}

type SettingsAction struct {
	Index   int      `json:"index"`
	Urls    []string `json:"urls"`
	Enabled bool     `json:"enabled"`
}

type SettingsActions struct {
	OutOnUrl             []SettingsAction `json:"out_on_url"`
	OutOffUrl            []SettingsAction `json:"out_off_url"`
	ShortPushUrl         []SettingsAction `json:"shortpush_url"`
	LongPushUrl          []SettingsAction `json:"longpush_url"`
	DoubleShortPushUrl   []SettingsAction `json:"double_shortpush_url"`
	TripleShortPushUrl   []SettingsAction `json:"triple_shortpush_url"`
	ShortPushLongPushUrl []SettingsAction `json:"shortpush_longpush_url"`
	LongPushShortPushUrl []SettingsAction `json:"longpush_shortpush_url"`
	// Add other fields as needed...
}

type ActionConfig struct {
	Actions SettingsActions `json:"actions"`
}
