package main

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
)

type appView int

const (
	listView appView = iota
	roomView
)

type roomsLoadedMsg struct {
	rooms []Room
	err   error
}

type lightSetMsg struct {
	err error
}

type model struct {
	list         list.Model
	spinner      spinner.Model
	loading      bool
	windowWidth  int
	windowHeight int
	currentView  appView
	selectedRoom Room
	roomCursor   int
	actionMsg    string
}
