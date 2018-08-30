package command

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/fuwalab/tools/conf"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/exec"
	"testing"
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

func TestAdd(t *testing.T) {
	params := []string{
		"AddAccount",
		"-s", "sample",
		"-u", "test_user",
	}

	os.Setenv("test_password", "password")
	Add(params...)
}

func TestShow(t *testing.T) {
	params := []string{
		"ShowAccount",
		"-s", "sample",
	}
	Show(params...)
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
	// Note: Not sure why it requires `TestAdd()` again for Travis CI.
	TestAdd(t)
	expected := "password"
	cmd := exec.Command("go", "run", config.ProjectRoot+"/main.go", "CopyPassword", "-s", "sample")
	result, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("commandline output: %v", string(result))
		t.Errorf("commandline error: %v", err)
	}

	actual, _ := clipboard.ReadAll()

	// TODO: fix here for travis CI
	if actual == "" {
		t.Skip("skipped. got empty value.")
	}

	if actual != expected {
		t.Errorf("Not same value.\nactual: %v, expected: %v", actual, expected)
	}

	cmd = exec.Command("go", "run", config.ProjectRoot+"/main.go", "CopyPassword", "-s", "sample_user")
	result, err = cmd.CombinedOutput()

	if err == nil {
		t.Errorf("commandline output: %v", string(result))
		t.Errorf("commandline error: %v", err)
	}
}
