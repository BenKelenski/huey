package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const listWidth = 40
const listHeight = 20

var (
	titleStyle    = lipgloss.NewStyle().Bold(true).MarginBottom(1)
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(1)
	msgStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).MarginTop(1)
	errStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).MarginTop(1)
)

func initializeModel() model {
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

func setLight(room Room, on bool) tea.Cmd {
	return func() tea.Msg {
		err := SetRoomLights(room, on)
		return lightSetMsg{err: err}
	}
}

func setColor(room Room, preset ColorPreset) tea.Cmd {
	return func() tea.Msg {
		err := SetRoomColor(room, preset.X, preset.Y)
		return colorSetMsg{err: err}
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchRooms)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if m.currentView == listView {
				return m, tea.Quit
			}
			if m.currentView == colorView {
				m.currentView = roomView
				return m, nil
			}
			m.currentView = listView
			m.actionMsg = ""
			return m, nil
		case "esc", "backspace":
			if m.currentView == colorView {
				m.currentView = roomView
				return m, nil
			}
			if m.currentView == roomView {
				m.currentView = listView
				m.actionMsg = ""
				return m, nil
			}
		}

		if m.currentView == colorView {
			switch msg.String() {
			case "up", "k":
				if m.colorCursor > 0 {
					m.colorCursor--
				}
				return m, nil
			case "down", "j":
				if m.colorCursor < len(colorPresets)-1 {
					m.colorCursor++
				}
				return m, nil
			case "enter", " ":
				preset := colorPresets[m.colorCursor]
				return m, setColor(m.selectedRoom, preset)
			}
		}

		if m.currentView == roomView {
			switch msg.String() {
			case "up", "k":
				if m.roomCursor > 0 {
					m.roomCursor--
				}
				return m, nil
			case "down", "j":
				if m.roomCursor < 2 {
					m.roomCursor++
				}
				return m, nil
			case "enter", " ":
				switch m.roomCursor {
				case 0:
					return m, setLight(m.selectedRoom, true)
				case 1:
					return m, setLight(m.selectedRoom, false)
				case 2:
					m.colorCursor = 0
					m.currentView = colorView
					return m, nil
				}
			}
		}

		if m.currentView == listView && !m.loading {
			switch msg.String() {
			case "enter", " ":
				if item, ok := m.list.SelectedItem().(Room); ok {
					m.selectedRoom = item
					m.roomCursor = 0
					m.actionMsg = ""
					m.currentView = roomView
					return m, nil
				}
			}
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

	case lightSetMsg:
		if msg.err != nil {
			m.actionMsg = "error: " + msg.err.Error()
		} else {
			if m.roomCursor == 0 {
				m.actionMsg = "Lights turned on"
			} else {
				m.actionMsg = "Lights turned off"
			}
		}
		return m, nil

	case colorSetMsg:
		m.currentView = roomView
		if msg.err != nil {
			m.actionMsg = "error: " + msg.err.Error()
		} else {
			m.actionMsg = "Color set to " + colorPresets[m.colorCursor].Name
		}
		return m, nil
	}

	if m.loading {
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.currentView == listView {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {
	var content string

	if m.loading {
		content = fmt.Sprintf("%s Loading rooms…", m.spinner.View())
	} else if m.currentView == colorView {
		content = m.colorViewContent()
	} else if m.currentView == roomView {
		content = m.roomViewContent()
	} else {
		content = m.list.View() + "\n" + subtleStyle.Render("enter/space to select")
	}

	centered := lipgloss.Place(m.windowWidth, m.windowHeight,
		lipgloss.Center, lipgloss.Center,
		content,
	)
	return tea.View{Content: centered, AltScreen: true}
}

func (m model) roomViewContent() string {
	options := []string{"Turn On", "Turn Off", "Set Color"}
	out := titleStyle.Render(m.selectedRoom.Metadata.Name) + "\n"

	for i, opt := range options {
		if i == m.roomCursor {
			out += selectedStyle.Render("> " + opt)
		} else {
			out += normalStyle.Render("  " + opt)
		}
		out += "\n"
	}

	if m.actionMsg != "" {
		style := msgStyle
		if len(m.actionMsg) > 6 && m.actionMsg[:6] == "error:" {
			style = errStyle
		}
		out += style.Render(m.actionMsg) + "\n"
	}

	out += subtleStyle.Render("↑/↓ to move • enter/space to select • esc to go back")
	return out
}

func (m model) colorViewContent() string {
	out := titleStyle.Render("Set Color — "+m.selectedRoom.Metadata.Name) + "\n"

	for i, preset := range colorPresets {
		swatch := lipgloss.NewStyle().Foreground(lipgloss.Color(preset.Display)).Render("■ ")
		if i == m.colorCursor {
			out += selectedStyle.Render("> ") + swatch + selectedStyle.Render(preset.Name)
		} else {
			out += normalStyle.Render("  ") + swatch + normalStyle.Render(preset.Name)
		}
		out += "\n"
	}

	out += subtleStyle.Render("↑/↓ to move • enter/space to select • esc to go back")
	return out
}

func main() {
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Uh oh, there was an error: %v\n", err)
		os.Exit(1)
	}
}
