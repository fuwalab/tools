// Package db contains anything related to database
package db

import (
	"database/sql"
	"fmt"
	"github.com/fuwalab/tools/conf"
	"github.com/labstack/gommon/log"
)

// Repo database connection
type Repo struct {
	db *sql.DB
}

var config = &conf.AppConf{}

// Conn create a new connection
func Conn() *sql.DB {
	config = conf.GetAppConf()
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/%s.db", config.ProjectRoot, config.DBName))
	if err != nil {
		panic(err)
	}

	return db
}

// NewRepo create a new repository interface.
func NewRepo(conn *sql.DB) *Repo {
	return &Repo{db: conn}
}

// InitDB initialize database
func (r *Repo) InitDB() {
	_, err := r.db.Exec(
		"CREATE TABLE IF NOT EXISTS `account` ( " +
			"`name` VARCHAR(255) NOT NULL PRIMARY KEY ," +
			"`account` VARCHAR(255) NOT NULL, " +
			"`password` VARCHAR(255) NOT NULL ) ",
	)

	if err != nil {
		panic(err)
	}
}

// Account account information
type Account struct {
	Name     string
	Account  string
	Password string
}

// Save save account
func (r *Repo) Save(account *Account) {
	var a Account
	row := r.db.QueryRow(
		"SELECT * FROM account WHERE name = ?", &account.Name)
	err := row.Scan(&a.Name, &a.Account, &a.Password)

	if err != nil {
		log.Info("error:", err)

		if err == sql.ErrNoRows {
			// insert
			log.Info("insert")
			_, err := r.db.Exec(
				"INSERT INTO account (name, account, password) VALUES(?, ?, ?)",
				&account.Name, &account.Account, &account.Password,
			)
			if err != nil {
				panic(err)
			}
		} else {
			// update
			log.Info("update")
			log.Info("account: ", account)
			_, err := r.db.Exec(
				"UPDATE account SET account = ?, password = ? WHERE name = ?",
				&account.Account, &account.Password, &a.Name,
			)
			if err != nil {
				panic(err)
			}
		}
	}
}

// FindAccountByName retrieve row by name
func (r *Repo) FindAccountByName(name string) (*Account, error) {
	var account Account
	row := r.db.QueryRow("SELECT * FROM account WHERE name = ?", &name)
	err := row.Scan(&account.Name, &account.Account, &account.Password)

	return &account, err
}
