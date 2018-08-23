package batch

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"os/exec"
	"testing"
	"tools/conf"
)

var config *conf.AppConf

func TestMain(m *testing.M) {
	conf.SetEnv("test")
	config = conf.GetAppConf()
	code := m.Run()

	if err := os.Remove(fmt.Sprintf("%s/%s.db", config.ProjectRoot, config.DBName)); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestAddAccount(t *testing.T) {
	cmd := exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
		"AddAccount",
		"-s", "sample",
		"-u", "hoge_user",
		"-p", "password",
	)
	err := cmd.Run()
	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	// check it will be raised error if there is a missing parameter
	cmd = exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
		"AddAccount",
		"-s", "sample",
		"-u", "hoge_user",
	)
	err = cmd.Run()
	if err == nil {
		t.Errorf("commandline error: %v", err)
	}
}

func TestShowAccount(t *testing.T) {
	cmd := exec.Command("go", "run", config.ProjectRoot+"/main.go", "ShowAccount", "-s", "sample")
	err := cmd.Run()
	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	// check it will be raised error if there is a missing parameter
	cmd = exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
		"AddAccount",
	)
	err = cmd.Run()
	if err == nil {
		t.Errorf("commandline error: %v", err)
	}
}

func TestCopyPassword(t *testing.T) {
	expected := "password"
	cmd := exec.Command("go", "run", config.ProjectRoot+"/main.go", "CopyPassword", "-s", "sample")
	err := cmd.Run()

	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	actual, _ := clipboard.ReadAll()

	if actual != expected {
		t.Errorf("Not same value.\nactual: %v, expected: %v", actual, expected)
	}

	cmd = exec.Command("go", "run", config.ProjectRoot+"/main.go", "CopyPassword", "-s", "sampl")
	err = cmd.Run()

	if err == nil {
		t.Errorf("commandline error: %v", err)
	}
}
