package models

type (
	Sepa struct {
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Locale   string `json:"locale"`
	}
	CreditCard struct {
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Locale   string `json:"locale"`
	}

	CheckoutData struct {
		ClientSecret string `json:"client_secret"`
	}
	Supporter struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	Payment struct {
		Id        string    `json:"id"`
		Provider  string    `json:"provider"`
		Supporter Supporter `json:"supporter"`
	}
)
