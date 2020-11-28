package models

type ConfirmRequest struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type ConfirmRLPResponse struct {
	ReturnCode    string      `json:"returnCode"`
	ReturnMessage string      `json:"returnMessage"`
	Info          ConfirmInfo `json:"info"`
}

type ConfirmInfo struct {
	OrderID                 string           `json:"orderId"`
	TransactionID           int64            `json:"transactionId"`
	AuthorizationExpireDate string           `json:"authorizationExpireDate,omitempty"`
	RegKey                  string           `json:"regKey,omitempty"`
	PayInfo                 []ConfirmPayInfo `json:"payInfo,omitempty"`
	Shipping                ConfirmShipping  `json:"shipping,omitempty"`
}

type ConfirmPayInfo struct {
	Method                 string `json:"method"`
	Amount                 int    `json:"amount"`
	CreditCardNickname     string `json:"creditCardNickname"`
	CreditCardBrand        string `json:"creditCardBrand"`
	MaskedCreditCardNumber string `json:"maskedCreditCardNumber"`
}

type ConfirmPackage struct {
	ID            string `json:"id"`
	Amount        int    `json:"amount"`
	UserFeeAmount int    `json:"userFeeAmount"`
}

type ConfirmShipping struct {
}
