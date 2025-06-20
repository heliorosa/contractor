package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/heliorosa/contractor/commands"
)

func notImplemented(parent *commands.Cmd, args []string) tea.Model {
	fmt.Println("::: args:", args)
	panic("NOT IMPLEMENTED")
}

var rootCmd = *commands.NewRoot(
	&commands.Cmd{
		Name:        "networks",
		Alias:       []string{"net", "n"},
		Description: "Networks management",
		Commands: []*commands.Cmd{
			{
				Name:        "list",
				Alias:       []string{"ls", "l"},
				Description: "List networks",
				Run:         notImplemented,
			},
			{
				Name:        "add",
				Alias:       []string{"a"},
				Description: "Add network",
				Run:         notImplemented,
			},
			{
				Name:        "delete",
				Alias:       []string{"del", "d"},
				Description: "Delete network",
				Run:         notImplemented,
			},
			{
				Name:        "rename",
				Alias:       []string{"ren", "r"},
				Description: "Rename network",
				Run:         notImplemented,
			},
			{
				Name:        "edit",
				Alias:       []string{"ed", "e"},
				Description: "Edit network",
				Run:         notImplemented,
			},
		},
	},
	&commands.Cmd{
		Name:        "accounts",
		Alias:       []string{"acc", "a"},
		Description: "Account management",
		Commands: []*commands.Cmd{
			{
				Name:        "list",
				Alias:       []string{"ls", "l"},
				Description: "List accounts",
				Run:         notImplemented,
			},
			{
				Name:        "add",
				Alias:       []string{"a"},
				Description: "Add account",
				Commands: []*commands.Cmd{
					{
						Name:        "plain",
						Alias:       []string{"p"},
						Description: "Add plain, unencrypted account",
						Run:         notImplemented,
					},
					{
						Name:        "encrypted",
						Alias:       []string{"e"},
						Description: "Add encrypted account",
						Run:         notImplemented,
					},
					{
						Name:        "hardware-wallet",
						Alias:       []string{"hw"},
						Description: "Add hardware wallet account",
						Run:         notImplemented,
					},
				},
			}, {
				Name:        "delete",
				Alias:       []string{"del", "d"},
				Description: "Delete account",
				Run:         notImplemented,
			}, {
				Name:        "rename",
				Alias:       []string{"ren", "r"},
				Description: "Rename account",
				Run:         notImplemented,
			},
		},
	},
	&commands.Cmd{
		Name:        "contracts",
		Alias:       []string{"c"},
		Description: "Contract management",
		Commands: []*commands.Cmd{
			{
				Name:        "list",
				Alias:       []string{"ls", "l"},
				Description: "List contracts",
				Run:         notImplemented,
			},
			{
				Name:        "add",
				Alias:       []string{"a"},
				Description: "Add contract",
				Run:         notImplemented,
			},
			{
				Name:        "delete",
				Alias:       []string{"del", "d"},
				Description: "Delete contract",
				Run:         notImplemented,
			},
			{
				Name:        "rename",
				Alias:       []string{"ren", "r"},
				Description: "Rename contract",
				Run:         notImplemented,
			},
		},
	},
	&commands.Cmd{
		Name:        "method",
		Alias:       []string{"m"},
		Description: "Call a contract method",
		Run:         notImplemented,
	})

func main() {
	err := commands.Run(&rootCmd, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

// type model struct {
// 	table table.Model
// 	view  string // can be "main" or "sub"
// }

// func main() {
// 	m := model{
// 		table: mainTable(),
// 		view:  "main",
// 	}
// 	if _, err := tea.NewProgram(m).Run(); err != nil {
// 		panic(err)
// 	}
// }

// func mainTable() table.Model {
// 	columns := []table.Column{
// 		{Title: "Main Entries", Width: 20},
// 	}
// 	rows := []table.Row{
// 		{"entryA"},
// 		{"entryB"},
// 		{"entryC"},
// 	}
// 	t := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithFocused(true))
// 	t.SetHeight(5)
// 	return t
// }

// func subTable() table.Model {
// 	columns := []table.Column{
// 		{Title: "Sub Entries", Width: 20},
// 	}
// 	rows := []table.Row{
// 		{"subA"},
// 		{"subB"},
// 		{"subC"},
// 	}
// 	t := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithFocused(true))
// 	t.SetHeight(5)
// 	return t
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "enter":
// 			selected := m.table.SelectedRow()[0]
// 			if m.view == "main" && selected == "entryA" {
// 				m.table = subTable()
// 				m.view = "sub"
// 			}
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		}
// 	}
// 	var cmd tea.Cmd
// 	m.table, cmd = m.table.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	style := lipgloss.NewStyle().
// 		BorderStyle(lipgloss.NormalBorder()).
// 		BorderForeground(lipgloss.Color("240")).
// 		Padding(1)

// 	return style.Render(m.table.View()) + "\n(Press enter to drill into entryA, q to quit)"
// }
