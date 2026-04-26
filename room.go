package main

type ResourceRef struct {
	RID   string `json:"rid"`
	RType string `json:"rtype"`
}

type RoomMetadata struct {
	Name      string `json:"name"`
	Archetype string `json:"archetype"`
}

type Room struct {
	ID       string        `json:"id"`
	IDV1     string        `json:"id_v1"`
	Children []ResourceRef `json:"children"`
	Services []ResourceRef `json:"services"`
	Metadata RoomMetadata  `json:"metadata"`
	Type     string        `json:"type"`
}

func (r Room) Title() string       { return r.Metadata.Name }
func (r Room) Description() string { return r.Metadata.Archetype }
func (r Room) FilterValue() string { return r.Metadata.Name }

type RoomsResponse struct {
	Errors []any  `json:"errors"`
	Data   []Room `json:"data"`
}
