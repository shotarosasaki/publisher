package interfaces

import (
	"net/http"

	"github.com/shotarosasaki/publisher/static"
)

func FacebookWebhookHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set(static.HttpHeaderContentType, static.ContentTypeJson)
	res.Write([]byte("facebook"))
}
