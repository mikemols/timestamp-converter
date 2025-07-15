package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput  textinput.Model
	converted  string
	err        error
	copyStatus string
}

type clipboardSuccessMsg struct{}
type clipboardErrorMsg struct{ err error }

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter a timestamp"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 60

	return model{
		textInput:  ti,
		converted:  "",
		err:        nil,
		copyStatus: "",
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handles keyboard input
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit // Exists the program
		case "enter":
			// Converts the timestamp
			result, err := convertTimestamp(m.textInput.Value())

			if err != nil {
				m.err = err
				m.converted = ""
				m.copyStatus = ""
			} else {
				m.err = nil
				m.converted = result
				m.copyStatus = "Copying..."
				return m, copyToClipboard(result)
			}
			return m, nil
		case "ctrl+l":
			m.textInput.SetValue("")
			m.converted = ""
			m.copyStatus = ""
			m.err = nil
			return m, nil
		}

	case clipboardSuccessMsg:
		m.copyStatus = "✓ Copied to clipboard!"
		return m, nil

	case clipboardErrorMsg:
		m.copyStatus = fmt.Sprintf("Copy failed: %v", msg.err)
		return m, nil
	}
	// returns current state unchanged for now
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func copyToClipboard(text string) tea.Cmd {
	return func() tea.Msg {
		err := clipboard.WriteAll(text)
		if err != nil {
			return clipboardErrorMsg{err}
		}

		return clipboardSuccessMsg{}
	}
}

func convertTimestamp(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	// Parses the input as integer
	timestamp, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return "", fmt.Errorf("Invalid number: %s", input)
	}

	// Auto-detect: seconds vs milliseconds based on digit count
	var t time.Time
	if len(input) == 13 {
		// Milliseconds (like your example: 1752574424823)
		t = time.Unix(0, timestamp*int64(time.Millisecond))
	} else if len(input) == 10 {
		// Seconds
		t = time.Unix(timestamp, 0)
	} else {
		return "", fmt.Errorf("timestamp must be 10 (seconds) or 13 (milliseconds) digits")
	}

	// Format the timestamp in local time
	return t.UTC().Format("Monday, January 2, 2006 15:04:05 MST"), nil
}

func (m model) View() string {
	convertedLine := m.converted

	if m.err != nil {
		convertedLine = fmt.Sprintf("Error: %s", m.err.Error())
	}

	copyStatusLine := m.copyStatus
	if copyStatusLine == "" {
		copyStatusLine = "Enter: Copy to clipboard | ESC: Quit"
	}

	return fmt.Sprintf(
		"┌─────────────────────────────────────────┐\n"+
			"│              Timestamp Converter        │\n"+
			"├─────────────────────────────────────────┤\n"+
			"│                                         │\n"+
			"│ Paste timestamp: %s                     │\n"+
			"│                                         │\n"+
			"│ Converted: %-29s │\n"+
			"│                                         │\n"+
			"├─────────────────────────────────────────┤\n"+
			"│ %-39s │\n"+
			"└─────────────────────────────────────────┘\n",
		m.textInput.View(),
		convertedLine,
		copyStatusLine,
	)
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
