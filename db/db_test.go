package db

import (
	"fmt"
	"os"
	"testing"
	"tools/conf"
)

func init() {
	conf.SetEnv("test")
	config = conf.GetAppConf()
}

func TestMain(m *testing.M) {
	code := m.Run()

	if err := os.Remove(fmt.Sprintf("./%s.db", config.DBName)); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestRepo_InitDB(t *testing.T) {
	NewRepo(Conn()).InitDB()

	_, err := os.Stat(fmt.Sprintf("./%s.db", config.DBName))
	if err != nil {
		t.Errorf("%v", err)
	}
}
