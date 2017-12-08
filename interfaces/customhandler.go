package interfaces

import (
	"fmt"
	"net/http"

	"github.com/shotarosasaki/publisher/config"
	"github.com/shotarosasaki/publisher/domain"
	"github.com/shotarosasaki/publisher/infrastructure/messaging"
)

// TODO リクエストパーサー、バリデーター、パブリッシャー、TDロガー等をfunc型として定義し、CreateHandlerFuncの引数に追加！
type customHandler interface {
	// TODO I/F検討
	// TODO バリデーションを含めるかメソッド分けるか検討
	ParseRequest(r *http.Request) (customResponse, error)
}

type customResponse interface {
	getAttribute() map[string]string
	getBody() []byte
}

func CreateHandlerFunc(c customHandler, cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	pub, err := messaging.NewPublisher(cfg.Queue)
	if err != nil {
		// TODO エラーログ
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO UUID生成

		// TODO リクエストパーサーのI/F見直し！（パース後の型を定義して、その結果に対して後続処理！）
		parsed, err := c.ParseRequest(r)
		if err != nil {
			// TODO エラーログ
			// TODO エラー用レスポンス
			return
		}

		// TODO パース・バリデーションまで終わったらその時点でレスポンス返すようにする！
		// TODO 以降のCloudPubSubへの送信はゴルーチン内で行う！

		in := &domain.PublishInput{
			Data:       parsed.getBody(),
			Attributes: parsed.getAttribute(),
		}
		out, err := pub.Publish(in)
		if err != nil {
			// TODO エラーログ
			// TODO エラー用レスポンス
			return
		}

		// TODO TDログ出力！

		// TODO レスポンス
		fmt.Println(out)
	}
}
