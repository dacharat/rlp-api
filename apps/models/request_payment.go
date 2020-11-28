package models

type RequestBody struct {
	Amount       int          `json:"amount"`
	Currency     string       `json:"currency"`
	OrderID      string       `json:"orderId"`
	Packages     []Package    `json:"packages"`
	RedirectUrls RedirectUrls `json:"redirectUrls"`
}

type RequestRLPResponse struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	Info          Info   `json:"info"`
}

type Package struct {
	ID       string    `json:"id"`
	Amount   int       `json:"amount"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

type RedirectUrls struct {
	ConfirmUrl string `json:"confirmUrl"`
	CancelUrl  string `json:"cancelUrl"`
}

type Info struct {
	TransactionId      int        `json:"transactionId"`
	PaymentAccessToken string     `json:"paymentAccessToken"`
	PaymentUrl         PaymentUrl `json:"paymentUrl"`
}

type PaymentUrl struct {
	App string `json:"app"`
	Web string `json:"web"`
}
