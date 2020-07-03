package models

type (
	Billing struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Interval string `json:"interval"`
		Locale   string `json:"locale"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
		Type     string `json:"type"`
	}
	ResponseMessage struct {
		Message string `json:"message"`
	}
)
