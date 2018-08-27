package command

import (
	"encoding/json"
	"os/exec"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	// encryption test
	expected := "225025e6\n"

	cmd := exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
		"String",
		"-e", "test",
	)

	result, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	if string(result) != expected {
		t.Errorf("got different values. actual: %v, expected: %v", string(result), expected)
	}

	// decryption test
	expected = "test\n"

	cmd = exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
		"String",
		"-d", "225025e6",
	)

	result, err = cmd.CombinedOutput()
	if err != nil {
		t.Errorf("commandline error: %v", err)
	}

	if string(result) != expected {
		t.Errorf("got different values. actual: %v, expected: %v", string(result), expected)
	}

	// error check
	cmd = exec.Command("go",
		"run", config.ProjectRoot+"/main.go",
		"String",
		"-d", "225025e6",
		"-e", "test",
	)

	result, _ = cmd.CombinedOutput()
	r := string(result)
	output := strings.Split(r, "\n")[0]

	type LogInfo struct {
		Message string `json:"message"`
	}

	var actualJSON, expectedJSON LogInfo

	expected = "{\"time\":\"2018-08-27T20:58:55.202269385+09:00\",\"level\":\"ERROR\",\"prefix\":\"-\",\"file\":\"strings.go\",\"line\":\"35\",\"message\":\"got too many arguments. Don't set `-e` and `-d` at the same time.\"}"

	json.Unmarshal([]byte(output), &actualJSON)
	json.Unmarshal([]byte(expected), &expectedJSON)

	if actualJSON.Message != expectedJSON.Message {
		t.Error("Got different sentences.")
		t.Errorf("actual\n%v", actualJSON.Message)
		t.Errorf("expected\n%v", expectedJSON.Message)
	}
}
