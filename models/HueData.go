package models

type HueData struct {
	Data []HueDevice `json:"data"`
}

type HueDevice struct {
	Id       string        `json:"id"`
	Metadata *Metadata     `json:"metadata,omitempty"`
	Services *[]Service    `json:"services,omitempty"`
	On       *OnState      `json:"on,omitempty"`
	Dimming  *DimmingState `json:"dimming,omitempty"`
	Color    *ColorState   `json:"color,omitempty"`
}

type Metadata struct {
	Archetype string `json:"archetype"`
	Name      string `json:"name"`
}

type Service struct {
	RID   string `json:"rid"`
	RType string `json:"rtype"`
}

type DimmingState struct {
	Brightness float64 `json:"brightness"`
}

type OnState struct {
	On bool `json:"on"`
}

type ColorState struct {
	XY XYState `json:"xy"`
}

type XYState struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
