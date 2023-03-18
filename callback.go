package snapMidtransGo

import (
	"crypto/sha512"
	"encoding/hex"
)

type ResponseFromCallback struct {
	TransactionStatus string `json:"transaction_status"`
	OrderId           string `json:"order_id"`
	StatusCode        string `json:"status_code"`
	SignatureKey      string `json:"signature_key"`
}

func ValidateSignatureKey(response ResponseFromCallback, orderId, ServerKey string) bool {
	str := orderId + response.StatusCode + ServerKey
	hash := sha512.Sum512([]byte(str))
	hashStr := hex.EncodeToString(hash[:])
	return string(hashStr) == response.SignatureKey
}
