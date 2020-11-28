package models

type RefundRequest struct {
	RefundAmount int `json:"refundAmount"`
}

type RefundResponse struct {
	ReturnCode    string     `json:"returnCode"`
	ReturnMessage string     `json:"returnMessage"`
	Info          RefundInfo `json:"info"`
}

type RefundInfo struct {
	RefundTransactionId   int    `json:"refundTransactionId"`
	RefundTransactionDate string `json:"refundTransactionDate"`
}
