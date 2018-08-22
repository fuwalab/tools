package main

import (
	"encoding/json"
	"os/exec"
	"testing"
)

func TestUsage(t *testing.T) {
	// no argument
	cmd := exec.Command("go",
		"run", "main.go",
	)
	result, _ := cmd.CombinedOutput()

	actual := string(result)
	expected := usage()

	if actual != expected {
		t.Error("Got different centences.")
		t.Errorf("actual\n%v", actual)
		t.Errorf("expected\n%v", expected)
	}

	// incorrect argument
	cmd = exec.Command("go",
		"run", "main.go", "DoSomething",
	)
	result, _ = cmd.CombinedOutput()

	actual = string(result)
	expected = `{"time":"2018-08-23T02:21:03.812357433+09:00","level":"INFO","prefix":"-","file":"main.go","line":"40","message":"subcommand DoSomething is not exist"}`

	type LogInfo struct {
		Message string `json:"message"`
	}

	var actualJson, expectedJson LogInfo

	json.Unmarshal([]byte(actual), &actualJson)
	json.Unmarshal([]byte(expected), &expectedJson)

	if actualJson.Message != expectedJson.Message {
		t.Error("Got different centences.")
		t.Errorf("actual\n%v", actualJson.Message)
		t.Errorf("expected\n%v", expectedJson.Message)
	}
}
