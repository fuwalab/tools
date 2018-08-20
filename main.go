package main

import (
	"flag"
	"fmt"
	"github.com/labstack/gommon/log"
	"tools/batch"
	"tools/db"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		usage()
		return
	}

	subCommands := map[string]func(){
		"AddAccount":   batch.AddAccount,
		"ShowAccount":  batch.ShowAccount,
		"CopyPassword": batch.CopyPassword,
	}

	command := args[0]

	if _, ok := subCommands[command]; ok {
		log.Info("Executing ", command)
		db.NewRepo(db.Conn()).InitDB()

		f := subCommands[command]
		f()
	} else {
		log.Info("subcommand ", args[0], " is not exist")
	}
}

func usage() {
	var usage = `
Usage:
	AddAccount: AddAccount [service name] [account] [password]
	ShowAccount:
	CopyPassword:
`
	fmt.Printf("%v\n", usage)
}
