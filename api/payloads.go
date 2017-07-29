package main

type NewGroupPayload struct {
	Code string `json:"code"`
}

type SignupPayload struct {
	Code  string `json:"code"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type MessagePayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
