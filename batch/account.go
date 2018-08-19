package batch

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/labstack/gommon/log"
	"tools/db"
	"tools/util"
)

// add account
func AddAccount() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 4 {
		log.Error("Invalid number of arguments: ", len(args))
		return
	}

	account := db.Account{
		Name:     util.Encrypt(args[1]),
		Account:  util.Encrypt(args[2]),
		Password: util.Encrypt(args[3]),
	}

	// save
	db.NewRepo(db.Conn()).Save(account)
}

// output account info
func ShowAccount() {
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Error("Invalid number of arguments: ", len(args))
		return
	}

	name := util.Encrypt(args[1])

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
