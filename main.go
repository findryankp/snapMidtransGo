package snapGoMidtrans

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DataPostMidtrans struct {
	OrderId   string
	Nominal   int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	ServerKey string
}

func checkRequirement(postData DataPostMidtrans) error {
	if postData.Email == "" || postData.FirstName == "" || postData.LastName == "" || postData.Nominal == 0 || postData.OrderId == "" || postData.ServerKey == "" || postData.Phone == "" {
		return errors.New("all requirement must be fill")
	}

	if postData.Nominal < 1000 {
		return errors.New("minimun transfer is Rp. 1.000")
	}

	return nil
}

func stringBody(postData DataPostMidtrans) (string, error) {
	if err := checkRequirement(postData); err != nil {
		return err.Error(), err
	}

	text := fmt.Sprintf(`{
		"transaction_details": {
			"order_id": "%s",
			"gross_amount": %d
		},
		"customer_details": {
			"first_name": "%s",
			"last_name": "%s",
			"email": "%s",
			"phone": "%s"
		}
	  }`, postData.OrderId, postData.Nominal, postData.FirstName, postData.LastName, postData.Email, postData.Phone)

	return text, nil
}

func postDataTransaction(link, encodedString string, dataBody []byte) (string, error) {
	req, err := http.NewRequest("POST", link, bytes.NewBuffer(dataBody))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Basic "+encodedString)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	response := map[string]any{}
	json.Unmarshal([]byte(body), &response)

	_, ok := response["redirect_url"]
	if ok {
		return response["redirect_url"].(string), nil
	}

	return "", errors.New("unautorized token")
}

func getDataTransaction(link, encodedString string) (map[string]any, error) {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return map[string]any{}, err
	}

	req.Header.Add("Authorization", "Basic "+encodedString)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return map[string]any{}, err
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	response := map[string]any{}
	json.Unmarshal([]byte(body), &response)
	return response, err
}
