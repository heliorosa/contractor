package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CmdFunc = func(parent *Cmd, args []string) tea.Model

type Cmd struct {
	isRoot      bool
	Name        string
	Alias       []string
	Description string
	Commands    []*Cmd
	Run         CmdFunc
	parent      *Cmd
}

func NewRoot(cmds ...*Cmd) *Cmd {
	r := &Cmd{Commands: cmds, isRoot: true}
	r.setParents()
	return r
}

func (c *Cmd) setParents() {
	for _, cmd := range c.Commands {
		cmd.parent = c
		cmd.setParents()
	}
}

func (c *Cmd) tableRow(isHelp bool) table.Row {
	r := append(make(table.Row, 0, 3), c.Name)
	if isHelp {
		r = append(r, strings.Join(c.Alias, ", "))
	}
	return append(r, c.Description)
}

func lastRow(isHelp bool) table.Row {
	r := table.Row{"..", "", "Parent"}
	if !isHelp {
		return slices.Delete(r, 1, 2)
	}
	return r
}

func (c *Cmd) table(isHelp bool) table.Model {
	cols := append(make([]table.Column, 0, 3), table.Column{Title: "Command", Width: 7})
	if isHelp {
		cols = append(cols, table.Column{Title: "Alias", Width: 5})
	}
	cols = append(cols, table.Column{Title: "Description", Width: 11})
	rows := make([]table.Row, 0, len(c.Commands)+1)
	for _, cmd := range c.Commands {
		row := cmd.tableRow(isHelp)
		for i, entry := range row {
			if l := len(entry); cols[i].Width < l {
				cols[i].Width = l
			}
		}
		rows = append(rows, row)
	}
	if !c.isRoot {
		rows = append(rows, lastRow(isHelp))
	}
	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(rows)+1),
	)
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}

type model struct {
	table  table.Model
	cmd    *Cmd
	isHelp bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			selected := m.table.SelectedRow()
			if selected[0] == ".." {
				parent := m.cmd.parent
				return model{cmd: parent, isHelp: m.isHelp, table: parent.table(m.isHelp)}, nil
			}
			for _, c := range m.cmd.Commands {
				if c.Name == selected[0] {
					if c.Run != nil {
						args := make([]string, 0, 1)
						if m.isHelp {
							args = append(args, "-h")
						}
						return c.Run(c.parent, args), nil
					}
					return model{cmd: c, isHelp: m.isHelp, table: c.table(m.isHelp)}, nil
				}
			}
		case "q", "esc":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		Render(m.table.View()) + "\n(Use arrow keys to navigate, enter to select, escape or q to quit)"
}

func findCommand(cmd *Cmd, args []string) (*Cmd, []string, error) {
	if len(args) == 0 || cmd.Run != nil {
		return cmd, nil, nil
	}
	for _, c := range cmd.Commands {
		if c.Name == args[0] || slices.Contains(c.Alias, args[0]) {
			return findCommand(c, args[1:])
		}
	}
	return nil, nil, fmt.Errorf("unknown command: %s", args[0])
}

func newTeaProgram(model tea.Model) error {
	_, err := tea.NewProgram(model).Run()
	return err
}

func Run(rootCmd *Cmd, args []string) error {
	if len(args) == 0 {
		return newTeaProgram(model{cmd: rootCmd, table: rootCmd.table(false)})
	}
	if args[0] == "help" || args[0] == "h" {
		cmd, _, err := findCommand(rootCmd, args[1:])
		if err != nil {
			return err
		}
		if cmd.Run != nil {
			return newTeaProgram(cmd.Run(cmd.parent, []string{"-h"}))
		}
		return newTeaProgram(model{cmd: cmd, isHelp: true, table: cmd.table(true)})
	}
	cmd, args, err := findCommand(rootCmd, args)
	if err != nil {
		return err
	}
	if cmd.Run != nil {
		return newTeaProgram(cmd.Run(cmd.parent, args))
	}
	return newTeaProgram(model{cmd: cmd, table: cmd.table(false)})
}
