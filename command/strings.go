package command

import (
	"flag"
	"fmt"
	"github.com/fuwalab/tools/util"
	"github.com/labstack/gommon/log"
	"os"
)

// String string related command
func String(args ...string) {
	// check parameters
	var plaintText, cipherText string
	var params = []string{
		"e",
		"d",
	}

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&plaintText, params[0], "", "set plain text")
	f.StringVar(&cipherText, params[1], "", "set encrypted text")
	f.Parse(args[1:])

	if f.NArg() == 1 {
		f.Usage()
		os.Exit(2)
		return
	}

	for 1 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if plaintText != "" && cipherText != "" {
		log.Error("got too many arguments. Don't set `-e` and `-d` at the same time.")
		f.Usage()
		os.Exit(2)
	}

	if plaintText != "" {
		fmt.Println(util.Encrypt(plaintText))
	}
	if cipherText != "" {
		fmt.Println(util.Decrypt(cipherText))
	}
}
