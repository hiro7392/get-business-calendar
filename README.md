# 営業日取得 API

この Cloud Function は、指定された年月・地域に基づいて、その月の全営業日（＝通常の土日および祝日を除いた日）の一覧を返す API です

## 概要

- **API 名**: GetBusinessDays
- **通信プロトコル**: gRPC（Protocol Buffers 定義）
- **用途**: 指定した年・月・国・地域における、営業日（＝平日かつ祝日でない日）の一覧を取得する

## プロトコルバッファ定義

以下は、API の proto 定義の例です。

## 使い方

1. **リクエストを作成する**  
   以下の情報を入力してください:
    - **year**: 対象の年 (例: 2025)
    - **month**: 対象の月 (1～12)
    - **country**: 国コード (例: "JP")
    - **region**: 地域コード (例: 都道府県コードや市区町村コード、例 "13")

2. **API を呼び出す**  
   作成したリクエストを Cloud Function に送信します。Cloud Function 内部で営業日の計算を行い、結果を返します。

3. **レスポンスを受け取る**  
   レスポンスには、その月の各営業日の情報が含まれています。各日付は以下の情報で表されます:
    - **date**: 日付 (ISO8601形式、例: "YYYY-MM-DD")
    - **weekday**: 曜日 (0 が日曜、1 が月曜、…、6 が土曜)
    - **note**: 補足情報（例: 「振替休日」など、必要に応じて利用されます）

## リクエストとレスポンスの例

### リクエスト例

以下は、2025 年 5 月、国コード "JP"、地域コード "13" を指定した場合のリクエスト例です。

```json
{
  "year": 2025,
  "month": 5,
  "country": "JP",
  "region": "13"
}
```

### レスポンス例
```shell
{
  "days": [
    {
      "date": "2025-05-01",
      "weekday": 4,
      "note": ""
    },
    {
      "date": "2025-05-02",
      "weekday": 5,
      "note": ""
    },
    {
      "date": "2025-05-05",
      "weekday": 1,
      "note": "振替休日"
    }
    // その他の営業日が続きます
  ]
}
```


### proto定義
```proto
syntax = "proto3";

package holiday;

// サービス定義
service HolidayService {
  // 指定した年月・地域の営業日一覧を返す API
  rpc GetBusinessDays(BusinessDaysRequest) returns (BusinessDaysResponse);
}

// リクエスト: 年、月、国・地域情報を指定
message BusinessDaysRequest {
  // 対象の年（例: 2025）
  int32 year = 1;
  // 対象の月（1～12）
  int32 month = 2;
  // 国コード（例："JP"）
  string country = 3; 
}

// 個々の営業日の情報
message BusinessDay {
  // 営業日の日付（ISO8601形式："YYYY-MM-DD"）
  string date = 1;
  // 曜日（0: 日曜, 1: 月曜, …, 6: 土曜）
  int32 weekday = 2;
}

// レスポンス: 対象月の全営業日一覧
message BusinessDaysResponse {
  // 営業日リスト
  repeated BusinessDay days = 1;
}