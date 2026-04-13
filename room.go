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

type RoomsResponse struct {
	Errors []any  `json:"errors"`
	Data   []Room `json:"data"`
}
