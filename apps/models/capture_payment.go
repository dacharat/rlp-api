package models

type CaptureResponse struct {
	ReturnCode    string      `json:"returnCode"`
	ReturnMessage string      `json:"returnMessage"`
	Info          CaptureInfo `json:"info"`
}

type CaptureInfo struct {
	OrderID       string           `json:"orderId"`
	TransactionID int64            `json:"transactionId"`
	PayInfo       []CapturePayInfo `json:"payInfo,omitempty"`
}

type CapturePayInfo struct {
	Method string `json:"method"`
	Amount int    `json:"amount"`
}
