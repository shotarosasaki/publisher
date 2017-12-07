package dummy
/*
【Facebook Webhook】
■要件
https://developers.facebook.com/docs/messenger-platform/webhook#security
・すべてのWebhookイベントに200 OKで応答する。
・すべてのWebhookイベントに20秒以下で応答する。

＜Webhookイベントの検証＞
HTTPリクエストにはX-Hub-Signatureヘッダーが含まれます。
このヘッダーには、リクエストペイロードのSHA1署名が含まれます。
署名のキーにはapp secretが使用され、sha1=のプリフィックスが追加されます。
コールバックエンドポイントではこの署名を確認することによって、ペイロードの整合性とソースを検証できます。

計算は、ペイロードのエスケープしたUnicodeバージョンで、小文字の16進数を使用して行われます。
たとえば、文字列äöåをエスケープすると\u00e4\u00f6\u00e5となります。
また、/をエスケープすると\/に、<は\u003Cに、%は\u0025に、@は\u0040になります。
デコードされたバイト数で計算すると、異なる署名が算出されます。
 */
/*
///////////////////////////////////////////////////////////////////////////////////////////////////
イベントの共通形式
///////////////////////////////////////////////////////////////////////////////////////////////////
{
  "object":"page",
  "entry":[
    {
      "id":"<PAGE_ID>",
      "time":1458692752478,
      "messaging":[
        {
          "sender":{
            "id":"<PSID>"
          },
          "recipient":{
            "id":"<PAGE_ID>"
          },

          ...
        }
      ]
    }
  ]
}
 */

/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★メッセージ受信イベントをフォローします
このコールバックは、ページにメッセージが送信されたときに発生

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messages
///////////////////////////////////////////////////////////////////////////////////////////////////
*/
/*
例：テキストメッセージ
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<PAGE_ID>"
  },
  "timestamp":1458692752478,
  "message":{
    "mid":"mid.1457764197618:41d102a3e1ae206a38",
    "text":"hello, world!",
    "quick_reply": {
      "payload": "<DEVELOPER_DEFINED_PAYLOAD>"
    }
  }
}
 */
/*
例：画像の添付ファイル付きのメッセージ
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<PAGE_ID>"
  },
  "timestamp":1458692752478,
  "message":{
    "mid":"mid.1458696618141:b4ef9d19ec21086067",
    "attachments":[
      {
        "type":"image",
        "payload":{
          "url":"<IMAGE_URL>"
        }
      }
    ]
  }
}
 */
/*
例：リンクのスクレイピングからフォールバックが添付されたメッセージ
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<PAGE_ID>"
  },
  "timestamp":1458692752478,
  "message":{
    "mid":"mid.1458696618141:b4ef9d19ec21086067",
    "text":"<URL_SENT_BY_THE_USER>",
    "attachments":[
      {
        "type":"fallback",
        "payload":null,
      	"title":"<TITLE_OF_THE_URL_ATTACHMENT>",
      	"URL":"<URL_OF_THE_ATTACHMENT>",
      }
    ]
  }
}
 */

/*
クイック返信のリクエストcurlの例
curl -X POST -H "Content-Type: application/json" -d '{
  "recipient":{
    "id":"<PSID>"
  },
  "message":{
    "text": "Here's a quick reply!",
    "quick_replies":[
      {
        "content_type":"text",
        "title":"Search",
        "payload":"<POSTBACK_PAYLOAD>",
        "image_url":"http://example.com/img/red.png"
      },
      {
        "content_type":"location"
      },
      {
        "content_type":"text",
        "title":"Something Else",
        "payload":"<POSTBACK_PAYLOAD>"
      }
    ]
  }
}' "https://graph.facebook.com/v2.6/me/messages?access_token=<PAGE_ACCESS_TOKEN>"
 */

/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★アカウントのリンクイベントをフォローします
アカウントリンクを使用する場合、このコールバックは、アカウントのリンクまたはアカウントのリンク解除ボタンがタップされたときに発生

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_account_linking
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
  "sender":{
    "id":"USER_ID"
  },
  "recipient":{
    "id":"PAGE_ID"
  },
  "timestamp":1234567890,
  "account_linking":{
    "status":"linked",
    "authorization_code":"PASS_THROUGH_AUTHORIZATION_CODE"
  }
}
 */
/*
{
  "sender":{
    "id":"USER_ID"
  },
  "recipient":{
    "id":"PAGE_ID"
  },
  "timestamp":1234567890,
  "account_linking":{
    "status":"unlinked"
  }
}
 */

/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★メッセージ配信イベントをフォローします
このコールバックは、ページが送信されたメッセージが配信されたときに発生

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/message-deliveries
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<PAGE_ID>"
  },
   "delivery":{
      "mids":[
         "mid.1458668856218:ed81099e15d3f4f233"
      ],
      "watermark":1458668856253,
      "seq":37
   }
}
 */

/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★メッセージエコーイベントをフォローします
このコールバックは、ページによってメッセージが送信されたときに発生します。
テキストメッセージや添付ファイル付きのメッセージ（画像、ビデオ、音声、テンプレート、フォールバック）を受け取ることがあります。
ペイロードには、送信者が送信したオプションのカスタムメタデータと対応するapp_idも含まれます。

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/message-echoes
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<USER_ID>"
  },
  "timestamp":1457764197627,
  "message":{
    "is_echo":true,
    "app_id":1517776481860111,
    "metadata": "<DEVELOPER_DEFINED_METADATA_STRING>",
    "mid":"mid.1457764197618:41d102a3e1ae206a38",
    ...
  }
}
 */
/*
例：Text message
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<USER_ID>"
  },
  "timestamp":1457764197627,
  "message":{
    "is_echo":true,
    "app_id":1517776481860111,
    "metadata": "<DEVELOPER_DEFINED_METADATA_STRING>",
    "mid":"mid.1457764197618:41d102a3e1ae206a38",
    "text":"hello, world!"
  }
}
 */
/*
例：Message with image, audio, video or file attachment
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<USER_ID>"
  },
  "timestamp":1458696618268,
  "message":{
    "is_echo":true,
    "app_id":1517776481860111,
    "metadata": "<DEVELOPER_DEFINED_METADATA_STRING>",
    "mid":"mid.1458696618141:b4ef9d19ec21086067",
    "attachments":[
      {
        "type":"image",
        "payload":{
          "url":"<IMAGE_URL>"
        }
      }
    ]
  }
}
 */
/*
例：Message with template attachment
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<USER_ID>"
  },
  "timestamp":1458696618268,
  "message":{
    "is_echo":true,
    "app_id":1517776481860111,
    "metadata": "<DEVELOPER_DEFINED_METADATA_STRING>",
    "mid":"mid.1458696618141:b4ef9d19ec21086067",
    "attachments":[
      {
        "type":"template",
        "payload":{
          "template_type":"button",
          "buttons":[
            {
              "type":"web_url",
              "url":"https:\/\/www.messenger.com\/",
              "title":"Visit Messenger"
            }
          ]
        }
      }
    ]
  }
}
 */
/*
例：Message with fallback attachment
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<USER_ID>"
  },
  "timestamp":1458696618268,
  "message":{
    "is_echo":true,
    "app_id":1517776481860111,
    "metadata": "<DEVELOPER_DEFINED_METADATA_STRING>",
    "mid":"mid.1458696618141:b4ef9d19ec21086067",
    "attachments":[
      {
        "title":"Legacy Attachment",
        "url":"https:\/\/www.messenger.com\/",
        "type":"fallback",
        "payload":null
      }
    ]
  }
}
 */

/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★プラグインのオプトインイベントをフォローします
このコールバックは、Messengerプラグインへの送信がタップされたとき、
ユーザーが顧客の一致を使用してメッセージ要求を受け入れたとき、
またはユーザーがチェックボックスプラグイン経由でメッセージを受信するように選択したときに発生

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_optins
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
  "sender": {
    "id": "<PSID>"
  },
  "recipient": {
    "id": "<PAGE_ID>"
  },
  "timestamp": 1234567890,
  "optin": {
    "ref": "<PASS_THROUGH_PARAM>",
    "user_ref": "<REF_FROM_CHECKBOX_PLUGIN>"
  }
}

※「ref」・・・エントリポイントで定義されたdata-ref属性
※「user_ref」・・・チェックボックスプラグインのみ。チェック・ボックス・プラグインに定義されたuser_ref属性が含まれています。
 */

/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★ポリシー順守のためのイベントをフォローします
アプリは、管理しているページでポリシー施行アクションが実行されると、このコールバックを受け取ります。

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_policy_enforcement
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
  "recipient":{
    "id":"PAGE_ID"
  },
  "timestamp":1458692752478,
  "policy-enforcement":{
    "action":"block",
    "reason":"The bot violated our Platform Policies (https://developers.facebook.com/policy/#messengerplatform). Common violations include sending out excessive spammy messages or being non-functional."
  }
}
 */
/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★ポストバック受信イベントをフォローします
ポストバックは、[ポストバック]ボタン、[スタート]ボタン、固定メニューがタップされたときに発生

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_postbacks
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
  "sender":{
    "id":"<PSID>"
  },
  "recipient":{
    "id":"<PAGE_ID>"
  },
  "timestamp":1458692752478,
  "postback":{
    "title": "<TITLE_FOR_THE_CTA>",
    "payload": "<USER_DEFINED_PAYLOAD>",
    "referral": {
      "ref": "<USER_DEFINED_REFERRAL_PARAM>",
      "source": "<SHORTLINK>",
      "type": "OPEN_THREAD",
    }
  }
}
 */
/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★メッセージ読み取りイベントをフォローします
このコールバックは、ページが送信されたメッセージがユーザーによって読み取られたときに発生

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/message-reads
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
{
   "sender":{
      "id":"<PSID>"
   },
   "recipient":{
      "id":"<PAGE_ID>"
   },
   "timestamp":1458668856463,
   "read":{
      "watermark":1458668856253,
      "seq":38
   }
}
 */
/*
///////////////////////////////////////////////////////////////////////////////////////////////////
★リファーラルイベントをフォローします
このコールバックは、すでにボットとのスレッドを持っている利用者が以下の経路でスレッドに移動してきたときに発生
・リファーラルパラメータが紐付けられたm.meリンクをたどった
・Messengerコンバージョン広告をクリックした
・パラメトリックMessengerコードをスキャンした

https://developers.facebook.com/docs/messenger-platform/reference/webhook-events/messaging_referrals
///////////////////////////////////////////////////////////////////////////////////////////////////
 */
/*
例：m.me
{
  "sender":{
    "id":"<USER ID>"
  },
  "recipient":{
    "id":"<PAGE ID>"
  },
  "timestamp":1458692752478,
  "referral": {
    "ref": <REF DATA PASSED IN M.ME PARAM>,
    "source": "SHORTLINK",
    "type": "OPEN_THREAD",
  }
}
 */
/*
例：広告リファーラル
{
  "sender":{
    "id":"<USER ID>"
  },
  "recipient":{
    "id":"<PAGE ID>
"
  },
  "timestamp":1458692752478,
  "referral": {
    "ref": <REF DATA IF SPECIFIED IN THE AD>,
    "ad_id": <ID OF THE AD>,
    "source": "ADS",
    "type": "OPEN_THREAD",
  }
}
 */
/*
例：パラメトリックMessengerコード
{
  "sender":{
    "id":"<USER_ID>"
  },
  "recipient":{
    "id":"<PAGE_ID>"
  },
  "timestamp":1458692752478,
  "referral": {
    "ref": <REF DATA PASSED IN CODE>,
    "source": "MESSENGER_CODE",
    "type": "OPEN_THREAD",
  }
}
 */
/*
例：[発見]タブの新しいスレッド
{
  "sender":{
    "id":"<USER ID>"
  },
  "recipient":{
    "id":"<PAGE ID>
"
  },
  "timestamp":1458692752478,
  "referral": {
    "source": "DISCOVER_TAB",
    "type": "OPEN_THREAD",
  }
}
 */


