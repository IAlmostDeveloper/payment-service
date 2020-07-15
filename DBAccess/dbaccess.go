package dbaccess

import (
	"payment-service/Entities"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("create table if not exists payments(id integer primary key autoincrement , " +
		"sum integer, purpose text, session_id text, completed numeric)")
}

func GetPayment(session_id string) entities.PaymentFromDB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query("select * from payments where session_id=$1", session_id)
	p := entities.PaymentFromDB{}
	for result.Next() {
		err := result.Scan(&p.Id, &p.Sum, &p.Purpose, &p.SessionId, &p.Completed)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p)
	}
	return p
}

func InsertPayment(payment entities.Payment, session_id string) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("insert into payments(sum, purpose, session_id, completed) values($1, $2, $3, false)",
		payment.Sum, payment.Purpose, session_id)
}

func MakePaymentComplete(session_id string) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("update payments set completed=true where session_id=$1", session_id)
}
