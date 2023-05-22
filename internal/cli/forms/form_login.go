package forms

import (
	"databuseasyway/internal/cli/forms/styles"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type inputsFormServerDB struct {
	nameUser     textinput.Model
	passwordUser textinput.Model
}

type FormLogin struct {
	focusIndex       int
	focusIndexDialog int
	buttonsFocus     []bool

	inputs inputsFormServerDB

	rows  []table.Row
	table table.Model

	uid string
	new bool

	message string

	spinner    spinner.Model
	spinnerRun bool

	dialog bool
}

func NewFormLogin() *FormLogin {
	f := &FormLogin{}
	return f

}

func (m *FormLogin) Init() tea.Cmd {
	return nil
}

func (m *FormLogin) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m *FormLogin) View() string {
	var b strings.Builder
	b.WriteRune('\n')
	title := fmt.Sprintf("%s Input name user and password %s", styles.ShortLine, styles.Line)
	b.WriteString(fmt.Sprintf(" %s\n\n", title))

	// Dialog.
	subtle := lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	buttonStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#888B7E")).
		Padding(0, 3).
		MarginTop(1)

	activeButtonStyle := buttonStyle.Copy().
		Foreground(lipgloss.Color("#FFF7DB")).
		Background(lipgloss.Color("#F25D94")).
		MarginRight(2).
		Underline(true)

	width := 96

	okButton := buttonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Cancel")
	if m.focusIndexDialog == 2 {
		okButton = activeButtonStyle.Render("Yes")
	}
	if m.focusIndexDialog == 3 {
		cancelButton = activeButtonStyle.Render("Cancel")
	}

	inputs := fmt.Sprintf("\n%s %s\n%s %s\n",
		styles.СaptionStyleFB.Render("User:"), m.inputs.nameUser.View(),
		styles.СaptionStyleFB.Render("Password:"), m.inputs.passwordUser.View())

	question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).
		Render("Enter the database name and password")
	buttons := lipgloss.JoinHorizontal(lipgloss.Top, okButton, " ", cancelButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, question, inputs, buttons)

	dialog := lipgloss.Place(width, 9,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars("猫咪"),
		lipgloss.WithWhitespaceForeground(subtle),
	)

	b.WriteString(dialog + "\n\n")

	return b.String()
}
