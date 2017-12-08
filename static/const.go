package static

// TODO とりあえず定数。随時、定数でよいのか、設定ファイル保持がよいのか、プログラム引数や環境変数渡しがよいのか検討！

const (
	RoutingPathPing = "/ping"

	// TODO Facebook側の仕様を考慮したラインアップにする（設定ファイル保持にする場合、Facebook/Twitter差分を吸収した形式にするかも検討）
	RoutingPathFacebookWebhook = "/facebook/webhook"

	// TODO Twitter側の仕様を考慮したラインアップにする（設定ファイル保持にする場合、Facebook/Twitter差分を吸収した形式にするかも検討）
	RoutingPathTwitterWebhook = "/twitter/webhook"
)

const (
	HttpHeaderContentType = "Content-Type"
	ContentTypeJson       = "application/json"
)
