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

	// TODO ここですべて定義せず、設定ファイル等に記載した内容に応じて使うハンドラーを制御するかどうかも検討！
	http.HandleFunc(static.RoutingPathFacebookWebhook, interfaces.CreateHandlerFunc(interfaces.NewFacebookWebhookHandler(), cfg))
	http.HandleFunc(static.RoutingPathTwitterWebhook, interfaces.CreateHandlerFunc(interfaces.NewTwitterWebhookHandler(), cfg))

	if err := http.ListenAndServe(cfg.Listen, nil); err != nil {
		return err
	}

	return nil
}
