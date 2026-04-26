package main

import (
	"fmt"
	"os"
	"time"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const listWidth = 40
const listHeight = 20

func initialModel() model {
	s := spinner.New(spinner.WithSpinner(spinner.Dot))
	return model{
		spinner: s,
		loading: true,
	}
}

func fetchRooms() tea.Msg {
	rooms, err := GetRooms()
	return roomsLoadedMsg{rooms: rooms, err: err}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchRooms)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
	case roomsLoadedMsg:
		if msg.err != nil {
			fmt.Printf("error while getting rooms: %s\n", msg.err)
			return m, tea.Quit
		}
		items := make([]list.Item, len(msg.rooms))
		for i, r := range msg.rooms {
			items[i] = r
		}
		delegate := list.NewDefaultDelegate()
		l := list.New(items, delegate, listWidth, listHeight)
		l.Title = "Rooms"
		m.list = l
		m.loading = false
		return m, nil
	}

	if m.loading {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() tea.View {
	var content string
	if m.loading {
		content = fmt.Sprintf("%s Loading rooms…", m.spinner.View())
	} else {
		content = m.list.View()
	}
	centered := lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		content,
	)
	return tea.View{Content: centered, AltScreen: true}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
