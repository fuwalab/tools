package conf

import (
	"os"
	"strings"
	"testing"
)

func TestSetEnv(t *testing.T) {
	SetEnv("test")

	actual := os.Getenv("env")
	expected := "test"

	if actual != expected {
		t.Errorf("Not same env.\nactual: %v, expected: %v", actual, expected)
	}
}

func TestGetAppConf(t *testing.T) {
	config := GetAppConf()

	expected := AppConf{
		Env:    "test",
		DBName: "test_account",
	}

	expectedProjectRootContains := "/src/github.com/fuwalab/tools"

	if config.DBName != expected.DBName {
		t.Errorf("Not same DBName.\nactual: %v, expected: %v", config.DBName, expected.DBName)
	}

	if config.Env != expected.Env {
		t.Errorf("Not same Env.\nactual: %v, expected: %v", config.Env, expected.Env)
	}

	if !strings.Contains(config.ProjectRoot, expectedProjectRootContains) {
		t.Errorf("Not contains path.\nactual: %v, expected: %v",
			config.ProjectRoot, expectedProjectRootContains)
	}
}
