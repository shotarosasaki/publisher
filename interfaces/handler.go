package interfaces

import (
	"context"
	"net/http"
)

type Handler struct {
	// TODO インフラ周りへのコネクション？
}

func (h *Handler) Start(context context.Context) error {
	go func() {
		// メインゴルーチンからのキャンセル通知を受けてWebサーバ停止
		// TODO いる？
		<-context.Done()
	}()

	// TODO 以降、各種定数（or設定ファイル持ち）化

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte("pong"))
	})

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		return err
	}

	return nil
}
