package conf

import "os"

type AppConf struct {
	Env         string
	DBName      string
	ProjectRoot string
}

// get application setting
func GetAppConf() AppConf {
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
func getLocalConf() AppConf {
	return AppConf{Env: "local", DBName: "account", ProjectRoot: getProjectRoot()}
}

// test environment
func getTestConf() AppConf {
	return AppConf{Env: "test", DBName: "test_account", ProjectRoot: getProjectRoot()}
}

// set application environment
func SetEnv(env string) {
	os.Setenv("env", env)
}

// get project root
func getProjectRoot() (path string) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		path = os.Getenv("HOME") + "/go/src/tools"
	} else {
		path = goPath + "/src/tools"
	}
	return
}
