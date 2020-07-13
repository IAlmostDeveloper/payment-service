package requests

import (
	dbaccess "PaymentAPI/DBAccess"
	entities "PaymentAPI/Entities"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func paymentRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var paymentResponse entities.PaymentSession
		json.NewDecoder(r.Body).Decode(&paymentResponse)
		payment := dbaccess.GetPayment(paymentResponse.SessionId)
		js, err := json.Marshal(payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
		break
	case "POST":
		var payment entities.Payment
		json.NewDecoder(r.Body).Decode(&payment)
		id, err := uuid.NewUUID()
		response := entities.PaymentSession{SessionId: id.String()}
		js, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		dbaccess.InsertPayment(payment, id.String())
		w.Write(js)
		fmt.Println(string(js))
		break
	case "PUT":
		var paymentResponse entities.PaymentSession
		json.NewDecoder(r.Body).Decode(&paymentResponse)
		dbaccess.MakePaymentComplete(paymentResponse.SessionId)
		break
	}
}

func validateCardRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		
	}
}

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/payment", paymentRequest)
	router.HandleFunc("/validate", validateCardRequest)
	log.Fatal(http.ListenAndServe(":8080", router))
}
