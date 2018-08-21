package conf

import "os"

type AppConf struct {
	Env    string
	DBName string
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
	return AppConf{Env: "local", DBName: "account"}
}

// test environment
func getTestConf() AppConf {
	return AppConf{Env: "test", DBName: "test_account"}
}

// set application environment
func SetEnv(env string) {
	os.Setenv("env", env)
}
