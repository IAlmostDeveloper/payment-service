package main

import (
	dbaccess "PaymentAPI/DBAccess"
	"PaymentAPI/Requests"
)

func main() {
	dbaccess.CreateDB()
	requests.HandleRequests()
}
