package db

import (
	"database/sql"
	"fmt"
	"github.com/fuwalab/tools/conf"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	conf.SetEnv("test")
	config = conf.GetAppConf()
	code := m.Run()

	os.Exit(code)
}

func TestRepo_InitDB(t *testing.T) {
	NewRepo(Conn()).InitDB()

	_, err := os.Stat(fmt.Sprintf("%s/%s.db", config.ProjectRoot, config.DBName))
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestRepo_Save(t *testing.T) {
	account := Account{
		Name:     "test",
		Account:  "hoge_user",
		Password: "password",
	}
	NewRepo(Conn()).Save(&account)
}

func TestRepo_FindAccountByName(t *testing.T) {
	account, _ := NewRepo(Conn()).FindAccountByName("test")

	expected := Account{
		Name:     "test",
		Account:  "hoge_user",
		Password: "password",
	}

	if *account != expected {
		t.Errorf("got unexpected account type.\nactual: %v, expected: %v", account, expected)
	}

	// failing test
	_, err := NewRepo(Conn()).FindAccountByName("tes")

	if err != sql.ErrNoRows {
		t.Errorf("got unexpected error.\nactual: %v, expected: %v", err, sql.ErrNoRows)
	}
}
