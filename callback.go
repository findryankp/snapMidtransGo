package snapMidtransGo

import (
	"crypto/sha512"
	"encoding/hex"
)

type ResponseFromCallbackMidtrans struct {
	TransactionId     string `json:"transaction_id"`
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
	GrossAmount       string `json:"gross_amount"`
}

func ValidateSignatureKey(response ResponseFromCallbackMidtrans, orderId, ServerKey string) bool {
	str := orderId + response.StatusCode + ServerKey
	hash := sha512.Sum512([]byte(str))
	hashStr := hex.EncodeToString(hash[:])
	return string(hashStr) == response.SignatureKey
}
