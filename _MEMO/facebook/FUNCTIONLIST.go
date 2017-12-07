package dummy
/*
Webhook		-	-	-
	message受信	システム	◎	Webhook	messagesイベント。メッセージ受信時、発生。
	message配信イベント	システム	◎	Webhook	message_deliveriesイベント。メッセージの送信が成功した際に発生。最後に送信したメッセージのタイムスタンプが記載されている。
	message既読イベント	システム	◎	Webhook	message_readsイベント。メッセージが既読された際に発生。
	messageエコーイベント	システム	◎	Webhook	message_echoesイベント。メッセージの送信が成功した際に発生。送信したメッセージの詳細な内容が記載。
	ポストバック受信イベント	システム	◎	Webhook	messaging_postbacksイベント。ポストバックボタン、スタートボタン、固定メニュークリック時に発生。
	オプトインイベント	システム	◎	Webhook	messaging_optinsイベント。外部サービスから、Messenger送信プラグイン（ボタン）により遷移した場合に発生。
	リファラールイベント	システム	◎	Webhook	messaging_referralsイベント。既にスレッドのやり取りがある利用者が、広告やショートリンク、MessengerコードからMessengerへ遷移した際に発生。
	アカウントリンクイベント	システム	◎	Webhook	messaging_account_linkingイベント。アカウントリンクボタンを押下された際に発生。
	ポリシー順守イベント	システム	◎	Webhook	messaging_policy_enforcementイベント。Messenger Platformポリシーに準拠していない、Facebookコミュニティの基準を満たしていない、Facebookページのガイドラインに違反している場合に発生。
エンキュー		-	-	-
	メッセージエンキュー	システム	◎	Webhook
通知処理		-	-	-
	message配信イベント通知	システム	◎	Webhook	messageが配信された事を通知。DBのステータス変更など。
	messageエコーイベント通知	システム	◎	Webhook	messageが配信された事を通知。DBのステータス変更など。
	アカウントリンクイベント処理	システム	◎	Webhook	渡された企業アカウント情報を、DB保存。
	ポリシー順守イベント通知	システム	◎	Webhook	渡されたポリシー違反情報を、メールやalertなどで通知。
 */