package snapGoMidtrans

import (
	"encoding/base64"
)

func RequestSnapMidtrans(postData DataPostMidtrans) (string, error) {
	stringBody, errBody := stringBody(postData)
	if errBody != nil {
		return "", errBody
	}

	link := "https://app.midtrans.com/snap/v1/transactions"
	encodedString := base64.StdEncoding.EncodeToString([]byte(postData.ServerKey))
	dataBody := []byte(stringBody)

	return postDataTransaction(link, encodedString, dataBody)
}

func GetTransaction(postData DataPostMidtrans) (map[string]any, error) {
	baseUrl := "https://api.midtrans.com/v2"
	link := baseUrl + "/" + postData.OrderId + "/status"
	encodedString := base64.StdEncoding.EncodeToString([]byte(postData.ServerKey))

	return getDataTransaction(link, encodedString)
}
