package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

const (
	User     = "xxx"
	Pwd      = "xxx"
	Database = "xxx"
)

func dbConn() (db *sql.DB, err error) {
	dbAddr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8", User, Pwd, Database)
	db, err = sql.Open("mysql", dbAddr)
	if err != nil {
		err = errors.Wrap(err, "Open database:"+Database+"fail")
		return
	}
	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "Ping database:"+Database+"fail")
		return
	}

	return
}

type Profile struct {
	id    int64
	name  string
	age   int8
	intro string
}

func queryProfile(db *sql.DB, sqlStr string) (err error) {
	rows, err := db.Query(sqlStr)
	if err != nil {
		return errors.Wrap(err, "Query fail:"+sqlStr)
	}

	for rows.Next() {
		var pro Profile
		err = rows.Scan(&pro.id, &pro.name, &pro.age, &pro.intro)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				err = errors.Wrap(err, "ErrNoRows")
			default:
				err = errors.Wrap(err, "Scan error")
			}
			return
		}
	}

	fmt.Println("queryProfile finished")

	return
}

func main() {
	sqlCon, err := dbConn()
	defer sqlCon.Close()
	if err != nil {
		fmt.Println("dbConn error:", err)
	}

	err = queryProfile(sqlCon, "select * from profile where id=1")
	if err != nil {
		fmt.Println("query sql error :", err)
	}
}
