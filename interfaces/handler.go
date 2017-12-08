package interfaces

import (
	"context"
	"net/http"

	"github.com/shotarosasaki/publisher/config"
	"github.com/shotarosasaki/publisher/global"
	"go.uber.org/zap"
)

type Handler struct {
	cfg *config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) Start(context context.Context) error {
	global.Logger.Debug("Start", zap.String("key", "val"))
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

	err := http.ListenAndServe(h.cfg.Listen, nil)
	if err != nil {
		return err
	}

	return nil
}
