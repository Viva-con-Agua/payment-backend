package models

type (
	//Billing stripe billing information
	Billing struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Interval string `json:"interval"`
		Locale   string `json:"locale"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
		Type     string `json:"type"`
		Product  string `json:"product"`
	}
	//BillingPaypal represents the initial model for paypal subscribtion
	BillingPaypal struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Interval    string `json:"interval"`
		Locale      string `json:"locale"`
		Amount      string `json:"amount"`
		Currency    string `json:"currency"`
		Type        string `json:"type"`
		Product     string `json:"product"`
	}

	//ResponseMessage message for response
	ResponseMessage struct {
		Message string `json:"message"`
	}

	//PaypalProduct represents an model for creating an product by using paypal api.
	PaypalProduct struct {
		Name        string `json:"name"`
		Description string `json:"description" validate:"required"`
		Type        string `bson:"type" json:"type" validate:"required"`
		Category    string `json:"category" bson:"category"`
		ImageURL    string `json:"image_url"`
		HomeURL     string `json:"home_url"`
	}
)
