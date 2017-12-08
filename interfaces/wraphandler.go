package interfaces

import (
	"fmt"
	"net/http"

	"github.com/shotarosasaki/publisher/config"
	"github.com/shotarosasaki/publisher/domain"
	"github.com/shotarosasaki/publisher/infrastructure/messaging"
)

// TODO リクエストパーサー、バリデーター、パブリッシャー、TDロガー等をfunc型として定義し、wrapHandlerの引数に追加！
type custumiser interface {
	// TODO I/F検討
	// TODO バリデーションを含めるかメソッド分けるか検討
	ParseRequest(r *http.Request)
}

func CreateHandlerFunc(c custumiser, cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	pub, err := messaging.NewPublisher(cfg.Queue)
	if err != nil {
		// TODO エラーログ
		return nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// TODO UUID生成

		// TODO リクエストパーサーのI/F見直し！（パース後の型を定義して、その結果に対して後続処理！）
		c.ParseRequest(r)

		// TODO とりあえずべたで
		data := []byte("Hello world !!!")
		attributes := map[string]string{"Type": "text"}
		in := &domain.PublishInput{
			Data:       data,
			Attributes: attributes,
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
