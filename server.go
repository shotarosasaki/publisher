package publisher

import (
	"net/http"

	"github.com/shotarosasaki/publisher/config"
	"github.com/shotarosasaki/publisher/global"
	"github.com/shotarosasaki/publisher/interfaces"
	"github.com/shotarosasaki/publisher/static"
	"go.uber.org/zap"
)

func Serve(cfg *config.Config) error {
	global.Logger.Debug("Serve Start", zap.String("func", "publisher.Serve"))

	http.HandleFunc(static.RoutingPathPing, interfaces.PingHandler)

	// TODO 設定ファイル等で使うハンドラーを制御するか検討！
	http.HandleFunc(static.RoutingPathFacebookWebhook, interfaces.FacebookWebhookHandler)
	http.HandleFunc(static.RoutingPathTwitterWebhook, interfaces.TwitterWebhookHandler)

	err := http.ListenAndServe(cfg.Listen, nil)
	if err != nil {
		return err
	}

	return nil
}
