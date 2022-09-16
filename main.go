package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type Struct struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
}

var Config = Struct{
	Host:     "localhost",
	Port:     "5432",
	User:     "",
	Password: "",
	DBName:   "",
	Sslmode:  "verify-full",
}

func init() {
	if os.Getenv("HOST") != "" {
		Config.Host = os.Getenv("HOST")
	}
	if os.Getenv("PORT") != "" {
		Config.Port = os.Getenv("PORT")
	}
	if os.Getenv("USER") != "" {
		Config.User = os.Getenv("USER")
	}
	if os.Getenv("PASSWORD") != "" {
		Config.Password = os.Getenv("PASSWORD")
	}
	if os.Getenv("DBNAME") != "" {
		Config.DBName = os.Getenv("DBNAME")
	}
	if os.Getenv("SSLMODE") != "" {
		Config.Sslmode = os.Getenv("SSLMODE")
	}
}

type ClientDB struct {
	db *sql.DB
}

func newClient(host string, port string, user string, password string, dbname string, sslmode string) *ClientDB {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	//defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	log.Info().Msg("Connected on database" + dbname)

	return &ClientDB{db}

}

func (c *ClientDB) insertDttm(date string, id string) {

	sql := `insert into "logs"("dttm", "id") values($1, $2)`

	_, err := c.db.Exec(sql, date, id)

	CheckError(err)
	log.Info().Msg("dttm added in logs table")


}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	db := newClient(Config.Host, Config.Port, Config.User, Config.Password, Config.DBName, Config.Sslmode)

	db.insertDttm("2022-09-16","2021516")

	db.db.Close()

}
