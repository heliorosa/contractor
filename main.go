package main

import (
	"embed"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/heliorosa/contractor/commands"
)

func notImplemented(parent *commands.Cmd, args []string) tea.Model {
	fmt.Println("::: args:", args)
	panic("NOT IMPLEMENTED")
}

//go:embed help
var helpDir embed.FS

var rootCmd *commands.Cmd

func init() {
	var err error
	rootCmd, err = commands.NewRoot(helpDir,
		&commands.Cmd{
			Name:        "method",
			Alias:       []string{"m"},
			Description: "Call a contract method",
			Run:         notImplemented,
		},
		&commands.Cmd{
			Name:        "networks",
			Alias:       []string{"net", "n"},
			Description: "Manage networks",
			Commands: []*commands.Cmd{
				{
					Name:        "list",
					Alias:       []string{"l", "ls"},
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
					Alias:       []string{"del", "d", "rm"},
					Description: "Delete network",
					Run:         notImplemented,
				},
				{
					Name:        "rename",
					Alias:       []string{"ren", "r", "mv"},
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
			Description: "Manage accounts",
			Commands: []*commands.Cmd{
				{
					Name:        "list",
					Alias:       []string{"l", "ls"},
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
					Alias:       []string{"del", "d", "rm"},
					Description: "Delete account",
					Run:         notImplemented,
				}, {
					Name:        "rename",
					Alias:       []string{"ren", "r", "mv"},
					Description: "Rename account",
					Run:         notImplemented,
				},
			},
		},
		&commands.Cmd{
			Name:        "contracts",
			Alias:       []string{"c"},
			Description: "Manage contracts",
			Commands: []*commands.Cmd{
				{
					Name:        "list",
					Alias:       []string{"l", "ls"},
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
					Alias:       []string{"del", "d", "rm"},
					Description: "Delete contract",
					Run:         notImplemented,
				},
				{
					Name:        "rename",
					Alias:       []string{"ren", "r", "mv"},
					Description: "Rename contract",
					Run:         notImplemented,
				},
				{
					Name:        "edit",
					Alias:       []string{"ed", "e"},
					Description: "Edit contract",
					Run:         notImplemented,
				},
			},
		},
		&commands.Cmd{
			Name:        "abis",
			Description: "Manage ABIs",
			Commands: []*commands.Cmd{
				{
					Name:        "list",
					Alias:       []string{"l", "ls"},
					Description: "List ABIs",
					Run:         notImplemented,
				},
				{
					Name:        "add",
					Alias:       []string{"a"},
					Description: "Add ABI",
					Run:         notImplemented,
				},
				{
					Name:        "delete",
					Alias:       []string{"del", "d", "rm"},
					Description: "Delete ABI",
					Run:         notImplemented,
				},
				{
					Name:        "rename",
					Alias:       []string{"ren", "r", "mv"},
					Description: "Rename ABI",
					Run:         notImplemented,
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
}

func main() {
	err := commands.Run(rootCmd, os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
