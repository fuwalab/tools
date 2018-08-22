package batch

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"tools/conf"
)

func TestMain(m *testing.M) {
	conf.SetEnv("test")
	code := m.Run()

	config := conf.GetAppConf()
	if err := os.Remove(fmt.Sprintf("%s/%s.db", config.ProjectRoot, config.DBName)); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestAddAccount(t *testing.T) {
	cmd := exec.Command("go",
		"run", "../main.go",
		"AddAccount",
		"-s", "sample",
		"-u", "hoge_user",
		"-p", "password",
	)
	err := cmd.Run()
	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	// TODO: check it will be raised error if there is a missing parameter
}

func TestShowAccount(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "ShowAccount", "-s", "sample")
	err := cmd.Run()
	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	// TODO: check it will be raised error if there is a missing parameter
}
