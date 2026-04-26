package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

const listWidth = 40
const listHeight = 20

func initialModel() model {
	rooms, err := GetRooms()
	if err != nil {
		fmt.Printf("error while getting rooms: %s\n", err)
		os.Exit(1)
	}

	items := make([]list.Item, len(rooms))
	for i, r := range rooms {
		items[i] = r
	}

	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, listWidth, listHeight)
	l.Title = "Rooms"

	return model{list: l}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	return tea.NewView(m.list.View())
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
