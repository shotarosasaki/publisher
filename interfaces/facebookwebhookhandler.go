package interfaces

import "net/http"

type FacebookWebhookHandler struct {
	// TODO 必要に応じてフィールド追加
}

// TODO 必要に応じて引数追加
func NewFacebookWebhookHandler() *FacebookWebhookHandler {
	return &FacebookWebhookHandler{}
}

func (h *FacebookWebhookHandler) ParseRequest(r *http.Request) {
	// TODO 実装！
}
