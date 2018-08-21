package db

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(conn *sql.DB) *Repo {
	return &Repo{db: conn}
}

func Conn() *sql.DB {
	db, err := sql.Open("sqlite3", "./account.db")
	if err != nil {
		panic(err)
	}

	return db
}

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

type Account struct {
	Name     string
	Account  string
	Password string
}

// save account
func (r *Repo) Save(account Account) {
	var a Account
	row := r.db.QueryRow(
		"SELECT * FROM account WHERE name = ?", &account.Name)
	err := row.Scan(&a.Name, &a.Password)

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
			&account.Account, &account.Password, &account.Name,
		)
		if err != nil {
			panic(err)
		}
	}
}

// retrieve row by name
func (r *Repo) FindAccountByName(name string) (*Account, error) {
	var account Account
	row := r.db.QueryRow("SELECT * FROM account WHERE name = ?", &name)
	err := row.Scan(&account.Name, &account.Account, &account.Password)

	return &account, err
}
