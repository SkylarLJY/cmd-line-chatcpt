package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	messages    []string
	err         error
	textarea    textarea.Model
	viewport    viewport.Model
	senderStyle lipgloss.Style
	resStyle    lipgloss.Style
}

func initModel() model {
	ta := textarea.New()
	ta.Focus()
	ta.SetHeight(5)
	ta.SetWidth(60)
	ta.Prompt = "|"
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false

	vp := viewport.New(60, 20)
	vp.MouseWheelEnabled = true

	return model{
		messages:    []string{},
		textarea:    ta,
		err:         nil,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("#579BB1")),
		resStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#ECA869")),
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var taCmd, vpCmd tea.Cmd
	m.textarea, taCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msgType := msg.(type) {
	case tea.KeyMsg:
		switch msgType.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			m.messages = append(m.messages, m.senderStyle.Render("You: ")+wrapStr(m.textarea.Value(), m.viewport.Width))

			reqMsg := m.textarea.Value()
			m.textarea.Reset()
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.viewport.GotoBottom()
			m.textarea.Blur()
			res, err := sendMessage(reqMsg)

			if err != nil {
				m.err = err
				fmt.Fprintln(os.Stderr, err)
				return m, nil
			}
			// res = strings.TrimSpace(res)
			m.messages = append(m.messages, m.resStyle.Render("Res: ")+wrapStr(strings.Trim(res, "\n")+"\n", m.viewport.Width))
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			m.textarea.Focus()
			m.viewport.GotoBottom()
		}
	case error:
		m.err = msgType
		return m, nil
	}

	return m, tea.Batch(taCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf("%s\n\n%s", m.viewport.View(), m.textarea.View()) + "\n\n"
}

// TODO: what if a word is longer than the set width
func wrapStr(str string, width int) string {
	if str == "" {
		return ""
	}
	words := strings.Split(str, " ")
	res := ""
	lineLength := 0
	for _, w := range words {
		if lineLength+len(w) > width {
			res = strings.Trim(res, " ")
			res += "\n"
			lineLength = 0
		}
		res += w + " "
		lineLength += len(w) + 1

	}
	return strings.Trim(res, " ")
}
