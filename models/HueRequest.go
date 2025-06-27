package models

type HueRequest struct {
	On      *OnState      `json:"on,omitempty"`
	Color   *ColorState   `json:"color,omitempty"`
	Dimming *DimmingState `json:"dimming,omitempty"`
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
