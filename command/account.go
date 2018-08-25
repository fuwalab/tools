package command

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/labstack/gommon/log"
	"os"
	"tools/db"
	"tools/util"
)

// Add add account
func Add() {
	var serviceName, userName, password string
	var params = []string{
		"s",
		"u",
		"p",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, params[0], "", "set service name")
	f.StringVar(&userName, params[1], "", "set user account")
	f.StringVar(&password, params[2], "", "set password")
	f.Parse(os.Args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}
	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	validateStringParams(params, f)

	account := db.Account{
		Name:     util.Encrypt(serviceName),
		Account:  util.Encrypt(userName),
		Password: util.Encrypt(password),
	}

	// save
	db.NewRepo(db.Conn()).Save(account)
}

// Show output account info
func Show() {
	var serviceName string
	var params = []string{
		"s",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, params[0], "", "set service name")
	f.Parse(os.Args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}
	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	validateStringParams(params, f)

	name := util.Encrypt(serviceName)

	// select
	account, err := db.NewRepo(db.Conn()).FindAccountByName(name)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Println(util.Decrypt(account.Account))
}

// CopyPassword copy password to clipboard
func CopyPassword() {
	var serviceName string
	var params = []string{
		"s",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, params[0], "", "set service name")
	f.Parse(os.Args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}
	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	validateStringParams(params, f)

	name := util.Encrypt(serviceName)

	// select
	account, err := db.NewRepo(db.Conn()).FindAccountByName(name)
	if err != nil {
		log.Error(err)
		os.Exit(2)
		return
	}
	clipboard.WriteAll(util.Decrypt(account.Password))
}

// validate if param value has set
func validateStringParams(names []string, f *flag.FlagSet) {
	hasError := false

	for _, name := range names {
		fl := f.Lookup(name)
		if fl.Value.String() == "" {
			hasError = true
		}
	}
	if hasError {
		f.Usage()
		os.Exit(2)
	}
}
