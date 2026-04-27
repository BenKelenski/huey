package main

type ColorPreset struct {
	Name    string
	X, Y    float64
	Display string // lipgloss-compatible color for terminal preview
}

var colorPresets = [...]ColorPreset{
	{Name: "Red", X: 0.675, Y: 0.322, Display: "#FF2020"},
	{Name: "Orange", X: 0.561, Y: 0.403, Display: "#FF8C00"},
	{Name: "Yellow", X: 0.477, Y: 0.453, Display: "#FFE000"},
	{Name: "Green", X: 0.214, Y: 0.709, Display: "#00DD00"},
	{Name: "Cyan", X: 0.167, Y: 0.237, Display: "#00DDDD"},
	{Name: "Blue", X: 0.167, Y: 0.040, Display: "#4040FF"},
	{Name: "Purple", X: 0.263, Y: 0.103, Display: "#9900FF"},
	{Name: "Pink", X: 0.365, Y: 0.159, Display: "#FF69B4"},
	{Name: "Warm White", X: 0.447, Y: 0.407, Display: "#FFD580"},
	{Name: "Cool White", X: 0.313, Y: 0.329, Display: "#E8F0FF"},
}
