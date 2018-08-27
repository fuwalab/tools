// Package main contains toolkit that may be needed.
//
// This tool has the following sub command so far.
// AddAccount, ShowAccount, CopyPassword
package main

import (
	"flag"
	"fmt"
	"github.com/fuwalab/tools/command"
	"github.com/fuwalab/tools/conf"
	"github.com/fuwalab/tools/db"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
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
		"String":       command.String,
	}

	c := args[0]

	if _, ok := subCommands[c]; ok {
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
  String: Show encrypted/decrypted string
	Run "String -h" for more detail.
`
}
