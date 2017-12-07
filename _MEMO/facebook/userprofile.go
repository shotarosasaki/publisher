package dummy
/*
ユーザープロフィールAPI

https://developers.facebook.com/docs/messenger-platform/identity/user-profile
 */
/*
■リクエスト
curl -X GET "https://graph.facebook.com/v2.6/<PSID>?fields=first_name,last_name,profile_pic&access_token=<PAGE_ACCESS_TOKEN>"
 */
/*
■レスポンス
{
  "first_name": "Peter",
  "last_name": "Chang",
  "profile_pic": "https://fbcdn-profile-a.akamaihd.net/hprofile-ak-xpf1/v/t1.0-1/p200x200/13055603_10105219398495383_8237637584159975445_n.jpg?oh=1d241d4b6d4dac50eaf9bb73288ea192&oe=57AF5C03&__gda__=1470213755_ab17c8c8e3a0a447fed3f272fa2179ce",
  "locale": "en_US",
  "timezone": -7,
  "gender": "male",
  "last_ad_referral": {
    "source": "ADS",
    "type": "OPEN_THREAD",
    "ad_id": "6045246247433"
  }
}

※「last_ad_referral」＝最後に参照された広告
 */