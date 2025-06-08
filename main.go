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
