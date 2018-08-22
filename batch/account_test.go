package batch

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"tools/conf"
)

func init() {
	conf.SetEnv("test")
}

func TestMain(m *testing.M) {
	code := m.Run()

	config := conf.GetAppConf()
	if err := os.Remove(fmt.Sprintf("%s/%s.db", config.ProjectRoot, config.DBName)); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestAddAccount(t *testing.T) {
	fmt.Printf("env: %v\n", os.Getenv("env"))
	cmd := exec.Command("go",
		"run", "../main.go",
		"AddAccount",
		"-s", "sample",
		"-u", "hoge_user",
		"-p", "password",
	)
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("hoge: %v\n", err)
	}
	fmt.Printf("result: %v\n", string(result))
}

func TestShowAccount(t *testing.T) {
	fmt.Printf("env: %v\n", os.Getenv("env"))
	cmd := exec.Command("go", "run", "../main.go", "ShowAccount", "-s", "sample")
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("hoge: %v\n", err)
	}
	fmt.Printf("result: %v\n", string(result))
}
