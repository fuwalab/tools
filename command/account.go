// Package command contains all commands of toolkit.
package command

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/fuwalab/tools/conf"
	"github.com/fuwalab/tools/db"
	"github.com/fuwalab/tools/util"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

// Add add account
func Add(args ...string) {
	var config = conf.GetAppConf()
	var serviceName, userName string
	var params = []string{
		"s",
		"u",
		//"p",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, params[0], "", "set service name")
	f.StringVar(&userName, params[1], "", "set user account")
	//f.StringVar(&password, params[2], "", "set password")
	f.Parse(args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}
	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	validateStringParams(params, f)

	var password []byte
	if config.Env != "test" {
		fmt.Print("Enter password: ")
		pass, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		password = pass
	} else {
		pass := []byte(os.Getenv("test_password"))
		password = pass
	}

	account := db.Account{
		Name:     util.Encrypt(serviceName),
		Account:  util.Encrypt(userName),
		Password: util.Encrypt(string(password)),
	}

	// save
	db.NewRepo(db.Conn()).Save(&account)
}

// Show output account info
func Show(args ...string) {
	var serviceName string
	var params = []string{
		"s",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, params[0], "", "set service name")
	f.Parse(args[1:])

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
func CopyPassword(args ...string) {
	var serviceName string
	var params = []string{
		"s",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&serviceName, params[0], "", "set service name")
	f.Parse(args[1:])

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
