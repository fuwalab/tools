package main

import (
	"encoding/json"
	"github.com/fuwalab/tools/conf"
	"github.com/fuwalab/tools/db"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	conf.SetEnv("test")
	db.NewRepo(db.Conn()).InitDB()
	code := m.Run()

	os.Exit(code)
}

func TestUsage(t *testing.T) {
	config := conf.GetAppConf()
	// no argument
	cmd := exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
	)
	result, _ := cmd.CombinedOutput()

	actual := string(result)
	expected := usage()

	if actual != expected {
		t.Error("Got different sentences.")
		t.Errorf("actual\n%v", actual)
		t.Errorf("expected\n%v", expected)
	}

	// incorrect argument
	cmd = exec.Command("go",
		"run", "main.go", "DoSomething",
	)
	result, _ = cmd.CombinedOutput()

	expected = `{"time":"2018-08-23T02:21:03.812357433+09:00","level":"INFO","prefix":"-","file":"main.go","line":"40","message":"subcommand DoSomething is not exist"}`

	type LogInfo struct {
		Message string `json:"message"`
	}

	var actualJSON, expectedJSON LogInfo

	json.Unmarshal(result, &actualJSON)
	json.Unmarshal([]byte(expected), &expectedJSON)

	if actualJSON.Message != expectedJSON.Message {
		t.Error("Got different sentences.")
		t.Errorf("actual\n%v", actualJSON.Message)
		t.Errorf("expected\n%v", expectedJSON.Message)
	}
}
