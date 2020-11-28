package apps

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dacharat/rlp-api/apps/models"
)

func createHeader(req *http.Request, path, body string) {
	nouce := time.Now().UTC()

	data := RLPConfig.ChannelSecret + path + body + nouce.String()

	h := hmac.New(sha256.New, []byte(RLPConfig.ChannelSecret))
	h.Write([]byte(data))
	hmacKey := base64.StdEncoding.EncodeToString(h.Sum(nil))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-LINE-ChannelId", RLPConfig.ChannelID)
	req.Header.Set("X-LINE-Authorization-Nonce", nouce.String())
	req.Header.Set("X-LINE-Authorization", hmacKey)
}

func RequestPayment() (*models.RequestRLPResponse, error) {
	path := "/v3/payments/request"
	url := RLPConfig.APIUrl + path
	body := createMockRequest()

	bodyString, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	createHeader(request, path, string(bodyString))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		res := &models.RequestRLPResponse{}

		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func ConfirmPayment(transactionId string) (*models.ConfirmRLPResponse, error) {
	path := fmt.Sprintf("/v3/payments/%s/confirm", transactionId)
	url := RLPConfig.APIUrl + path
	body := models.ConfirmRequest{
		Amount:   250,
		Currency: "THB",
	}

	bodyString, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	createHeader(request, path, string(bodyString))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		res := &models.ConfirmRLPResponse{}

		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func CapturePayment(transactionId string) (*models.CaptureResponse, error) {
	path := fmt.Sprintf("/v3/payments/authorizations/%s/capture", transactionId)
	url := RLPConfig.APIUrl + path
	body := models.ConfirmRequest{
		Amount:   250,
		Currency: "THB",
	}

	bodyString, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	createHeader(request, path, string(bodyString))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		res := &models.CaptureResponse{}

		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func VoidPayment(transactionId string) (*models.VoidResponse, error) {
	path := fmt.Sprintf("/v3/payments/authorizations/%s/void", transactionId)
	url := RLPConfig.APIUrl + path

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	createHeader(request, path, "")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		res := &models.VoidResponse{}

		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func RefundPayment(transactionId string) (*models.RefundResponse, error) {
	path := fmt.Sprintf("/v3/payments/%s/refund", transactionId)
	url := RLPConfig.APIUrl + path
	body := models.RefundRequest{
		RefundAmount: 250,
	}

	bodyString, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(bodyString))
	if err != nil {
		return nil, err
	}

	createHeader(request, path, string(bodyString))

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		res := &models.RefundResponse{}

		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func createMockRequest() models.RequestBody {
	return models.RequestBody{
		Amount:   250,
		Currency: "THB",
		OrderID:  fmt.Sprint(time.Now().Unix()),
		Packages: []models.Package{
			{
				ID:     "01A",
				Amount: 250,
				Name:   "Toy Package",
				Products: []models.Product{
					{
						Name:     "ตุ๊กตา Cony",
						Quantity: 1,
						Price:    100,
						ImageUrl: "https://firebasestorage.googleapis.com/v0/b/linedeveloper-63341.appspot.com/o/512x512bb.jpg?alt=media&token=7cfd10b0-6d01-4612-b42e-b1b4d0105acd",
					},
					{
						Name:     "ตุ๊กตา Sally",
						Quantity: 1,
						Price:    150,
						ImageUrl: "https://firebasestorage.googleapis.com/v0/b/linedeveloper-63341.appspot.com/o/8cd724371a6f169b977684fd69cc2339.jpg?alt=media&token=e2008ff7-1cad-4476-a2e4-cda5f8af6561",
					},
				},
			},
		},
		RedirectUrls: models.RedirectUrls{
			ConfirmUrl: "https://www.google.com",
			CancelUrl:  "https://www.google.com",
		},
	}
}
