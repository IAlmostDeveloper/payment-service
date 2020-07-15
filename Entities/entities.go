package entities

type Payment struct {
	Sum     int    `json:"sum"`
	Purpose string `json:"purpose"`
}

type PaymentFromDB struct {
	Id int `json:"id"`
	Sum int `json:"sum"`
	Purpose string `json:"purpose"`
	SessionId string `json:"session_id"`
	ExpireTime string `json:"expire_time"`
	Completed bool `json:"completed"`
	Card string `json:"card"`
}

type PaymentSession struct {
	SessionId string `json:"session_id"`
}

type CardData struct {
	User string `json:"user"`
	Number     string `json:"number"`
	CVV        int    `json:"cvv"`
	ExpireDate string `json:"expire_date"`
	SessionId  string `json:"session_id"`
}

type CardValidationResponse struct{
	Error string `json:"error"`
}
