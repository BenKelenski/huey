package main

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
)

type roomsLoadedMsg struct {
	rooms []Room
	err   error
}

type model struct {
	list         list.Model
	spinner      spinner.Model
	loading      bool
	windowWidth  int
	windowHeight int
}
