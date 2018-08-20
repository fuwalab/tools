package batch

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/labstack/gommon/log"
	"os"
	"tools/db"
	"tools/util"
)

// add account
func AddAccount() {
	var serviceName, userName, password string

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, "s", "", "set service name")
	f.StringVar(&userName, "u", "", "set user account")
	f.StringVar(&password, "p", "", "set password")
	f.Parse(os.Args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}
	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	account := db.Account{
		Name:     util.Encrypt(serviceName),
		Account:  util.Encrypt(userName),
		Password: util.Encrypt(password),
	}

	// save
	db.NewRepo(db.Conn()).Save(account)
}

// output account info
func ShowAccount() {
	var serviceName string

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, "s", "", "set service name")
	f.Parse(os.Args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}
	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	validateParams("s", f)

	name := util.Encrypt(serviceName)

	// select
	account := db.NewRepo(db.Conn()).FindAccountByName(name)
	fmt.Println(util.Decrypt(account.Account))
}

// copy password to clipboard
func CopyPassword() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Error("Invalid number of arguments: ", len(args))
		return
	}

	name := util.Encrypt(args[1])

	// select
	account := db.NewRepo(db.Conn()).FindAccountByName(name)
	clipboard.WriteAll(util.Decrypt(account.Password))
}

// TODO: name should be slice. names []string
func validateParams(name string, f *flag.FlagSet) {
	fl := f.Lookup(name)
	if fl.Value.String() == "" {
		fmt.Println(fl.Usage)
		os.Exit(2)
	}
}
