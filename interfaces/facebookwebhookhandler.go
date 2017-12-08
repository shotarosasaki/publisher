package interfaces

import (
	"io/ioutil"
	"net/http"
)

type FacebookWebhookHandler struct {
	// TODO 必要に応じてフィールド追加
}

// TODO 必要に応じて引数追加
func NewFacebookWebhookHandler() *FacebookWebhookHandler {
	return &FacebookWebhookHandler{}
}

func (h *FacebookWebhookHandler) ParseRequest(r *http.Request) (customResponse, error) {
	// TODO POSTのみに限る

	// リクエストボディ取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return &FacebookWebhookResponse{
		// TODO 実装！
		attributes: map[string]string{"Type": "text"},
		body:       body,
	}, nil
}

type FacebookWebhookResponse struct {
	attributes map[string]string
	body       []byte
}

func (r *FacebookWebhookResponse) getAttribute() map[string]string {
	return r.attributes
}

func (r *FacebookWebhookResponse) getBody() []byte {
	return r.body
}
