package commands

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type CmdFunc = func(parent *Cmd, args []string) tea.Model

type Cmd struct {
	isRoot      bool
	Name        string
	Alias       []string
	Description string
	Commands    []*Cmd
	Run         CmdFunc
	help        string
	parent      *Cmd
}

func NewRoot(helpFS fs.FS, cmds ...*Cmd) (*Cmd, error) {
	r := &Cmd{Commands: cmds, isRoot: true}
	setParents(r)
	if err := loadHelp(helpFS, "help", r); err != nil {
		return nil, err
	}
	return r, nil
}

func loadHelp(helpFS fs.FS, path string, cmd *Cmd) error {
	if len(cmd.Commands) == 0 {
		f, err := helpFS.Open(path + ".md")
		if err != nil {
			return err
		}
		defer f.Close()
		b, err := io.ReadAll(f)
		if err != nil {
			return err
		}
		cmd.help = string(b)
		return nil
	}
	for _, c := range cmd.Commands {
		if err := loadHelp(helpFS, path+"/"+c.Name, c); err != nil {
			return err
		}
	}
	return nil
}

func setParents(c *Cmd) {
	for _, cmd := range c.Commands {
		cmd.parent = c
		setParents(cmd)
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

type menu struct {
	table  table.Model
	cmd    *Cmd
	isHelp bool
}

func (m menu) Init() tea.Cmd {
	return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			selected := m.table.SelectedRow()
			if selected[0] == ".." {
				parent := m.cmd.parent
				return menu{cmd: parent, isHelp: m.isHelp, table: parent.table(m.isHelp)}, nil
			}
			for _, c := range m.cmd.Commands {
				if c.Name == selected[0] {
					if c.Run != nil {
						if m.isHelp {
							hv, err := newHelpView(c, true)
							if err != nil {
								panic(err)
							}
							return hv, nil
						}
						return c.Run(c.parent, []string{}), nil
					}
					return menu{cmd: c, isHelp: m.isHelp, table: c.table(m.isHelp)}, nil
				}
			}
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m menu) View() string {
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
		return newTeaProgram(menu{cmd: rootCmd, table: rootCmd.table(false)})
	}
	if args[0] == "help" || args[0] == "h" {
		cmd, _, err := findCommand(rootCmd, args[1:])
		if err != nil {
			return err
		}
		if cmd.Run != nil {
			hv, err := newHelpView(cmd, false)
			if err != nil {
				return err
			}
			return newTeaProgram(hv)
		}
		return newTeaProgram(menu{cmd: cmd, isHelp: true, table: cmd.table(true)})
	}
	cmd, args, err := findCommand(rootCmd, args)
	if err != nil {
		return err
	}
	if cmd.Run != nil {
		return newTeaProgram(cmd.Run(cmd.parent, args))
	}
	return newTeaProgram(menu{cmd: cmd, table: cmd.table(false)})
}

type helpView struct {
	viewport  viewport.Model
	cmd       *Cmd
	canGoBack bool
}

func renderMarkdown(content string, width int) (string, error) {
	r, err := glamour.NewTermRenderer(
		glamour.WithWordWrap(width),
		glamour.WithStandardStyle("dark"),
	)
	if err != nil {
		return "", err
	}
	return r.Render(content)
}

func newHelpView(cmd *Cmd, canGoBack bool) (helpView, error) {
	w, h, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return helpView{}, err
	}
	out, err := renderMarkdown(cmd.help, w)
	if err != nil {
		return helpView{}, err
	}
	vp := viewport.New(w, h)
	vp.SetContent(out)
	vp.KeyMap = viewport.DefaultKeyMap()
	vp.MouseWheelEnabled = true
	return helpView{viewport: vp, cmd: cmd, canGoBack: canGoBack}, nil
}

func (hv helpView) Init() tea.Cmd { return nil }

func (hv helpView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "b", "esc":
			if hv.canGoBack {
				parent := hv.cmd.parent
				return menu{cmd: parent, isHelp: true, table: parent.table(true)}, nil
			}
		case "q", "ctrl+c":
			return hv, tea.Quit
		}
	case tea.WindowSizeMsg:
		rendered, err := renderMarkdown(hv.cmd.help, msg.Width)
		if err != nil {
			panic(err)
		}
		hv.viewport = viewport.New(msg.Width, hv.viewport.Height)
		hv.viewport.SetContent(rendered)
	}
	hv.viewport, cmd = hv.viewport.Update(msg)
	return hv, cmd
}

func (hv helpView) View() string {
	msg := "\n["
	if hv.canGoBack {
		msg += "b to go back, "
	}
	msg += "q to exit]"
	return hv.viewport.View() + msg
}
