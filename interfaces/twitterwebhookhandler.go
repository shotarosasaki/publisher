package interfaces

import (
	"net/http"
)

type TwitterWebhookHandler struct {
	// TODO 必要に応じてフィールド追加
}

// TODO 必要に応じて引数追加
func NewTwitterWebhookHandler() *TwitterWebhookHandler {
	return &TwitterWebhookHandler{}
}

func (h *TwitterWebhookHandler) ParseRequest(r *http.Request) (customResponse, error) {
	// TODO 実装！
	return nil, nil
}
