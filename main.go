package main

import (
	dbaccess "payment-service/DBAccess"
	"payment-service/Requests"
)

func main() {
	dbaccess.CreateDB()
	requests.HandleRequests()
}
