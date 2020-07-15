package dbaccess

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"payment-service/Entities"
)

func CreateDB() {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("create table if not exists payments(id integer primary key autoincrement , " +
		"sum integer, purpose text, session_id text, expire_time text, completed numeric, card text)")
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
		err := result.Scan(&p.Id, &p.Sum, &p.Purpose, &p.SessionId, &p.ExpireTime, &p.Completed, &p.Card)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(p)
	}
	return p
}

func InsertPayment(payment entities.Payment, session_id string, expire_time string) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("insert into payments(sum, purpose, session_id, expire_time, completed, card) values($1, $2, $3, $4, false, '')",
		payment.Sum, payment.Purpose, session_id, expire_time)
}

func MakePaymentComplete(session_id string, card_number string) {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Exec("update payments set completed=true, card=$1 where session_id=$2", card_number, session_id)
}
