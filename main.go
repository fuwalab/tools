package main

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"tools/command"
	"tools/conf"
	"tools/db"
)

func init() {
	conf.SetEnv("local")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Print(usage())
		return
	}

	subCommands := map[string]func(){
		"AddAccount":   command.Add,
		"ShowAccount":  command.Show,
		"CopyPassword": command.CopyPassword,
	}

	c := args[0]

	if _, ok := subCommands[c]; ok {
		log.Info("Executing ", c)
		db.NewRepo(db.Conn()).InitDB()

		f := subCommands[c]
		f()
	} else {
		log.Info("subcommand ", args[0], " is not exist")
	}
}

func usage() string {
	return `Usage:
  AddAccount: Add a new account information.
	Run "AddAccount -h" for more detail.
  ShowAccount: Show account/user name of a particular service.
	Run "ShowAccount -h" for more detail.
  CopyPassword: Copy password of the particular service to clipboard.
	Run "CopyPassword -h" for more detail.
`
}
