package models

type (
	Money struct {
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}

	CheckoutData struct {
		ClientSecret string `json:"client_secret"`
	}
)
