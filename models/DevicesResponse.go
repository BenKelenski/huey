package models

type DevicesResponse struct {
	Data []HueDevice `json:"data"`
}

type HueDevice struct {
	Metadata Metadata  `json:"metadata"`
	Services []Service `json:"services"`
}

type Metadata struct {
	Archetype string `json:"archetype"`
	Name      string `json:"name"`
}

type Service struct {
	RID   string `json:"rid"`
	RType string `json:"rtype"`
}
