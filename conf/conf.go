// Package conf contains configurations of toolkit
package conf

import "os"

// AppConf Configuration of this project.
type AppConf struct {
	Env         string
	DBName      string
	ProjectRoot string
}

// GetAppConf get application setting
func GetAppConf() *AppConf {
	var env = os.Getenv("env")

	switch env {
	case "test":
		return getTestConf()
	case "local":
		return getLocalConf()
	default:
		return getLocalConf()
	}
}

// local environment
func getLocalConf() *AppConf {
	return &AppConf{Env: "local", DBName: "account", ProjectRoot: getProjectRoot()}
}

// test environment
func getTestConf() *AppConf {
	return &AppConf{Env: "test", DBName: "test_account", ProjectRoot: getProjectRoot()}
}

// SetEnv set application environment
func SetEnv(env string) {
	os.Setenv("env", env)
}

// get project root
func getProjectRoot() (path string) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		path = os.Getenv("HOME") + "/go/src/github.com/fuwalab/tools"
	} else {
		path = goPath + "/src/github.com/fuwalab/tools"
	}
	return
}
