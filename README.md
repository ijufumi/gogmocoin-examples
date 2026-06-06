# gogmocoin サンプルコード

このリポジトリには、[gogmocoin](https://github.com/ijufumi/gogmocoin) ライブラリを使用したGMO Coin APIのサンプルコードが含まれています。

## 概要

gogmocoinは、GMO CoinのREST APIとWebSocket APIをGoから簡単に利用できるラッパーライブラリです。このリポジトリには、公開API(Public API)とプライベートAPI(Private API)の両方のサンプルコードが用意されています。

## ディレクトリ構成

```
app/
├── public/          # 公開APIのサンプル
│   ├── rest/        # 公開REST APIのサンプル
│   └── ws/          # 公開WebSocket APIのサンプル
└── private/         # プライベートAPIのサンプル
    ├── rest/        # プライベートREST APIのサンプル
    └── ws/          # プライベートWebSocket APIのサンプル
```

## セットアップ

### 必要なもの

- Go 1.25以上
- GMO CoinのAPIキーとシークレット（プライベートAPIを使用する場合）

### インストール

```bash
go mod download
```

### 環境変数の設定

プライベートAPIを使用するには、APIキーとシークレットを環境変数として設定する必要があります。

1. `.env.example` を `.env` にコピー:

```bash
cp .env.example .env
```

2. `.env` ファイルを編集し、あなたのAPIキーとシークレットを設定:

```
DEBUG_MODE=false
API_KEY=your_actual_api_key
API_SECRET=your_actual_api_secret
```

**注意**: `DEBUG_MODE=true` にするとリクエスト/レスポンスの詳細（署名済みの API-KEY ヘッダを含む）が標準出力に出力されます。認証情報がログに残るため、通常は `false` のままにし、デバッグ時のみ一時的に有効化してください。

**注意**: `.env` ファイルにはAPIの認証情報が含まれるため、絶対にgitにコミットしないでください。このファイルは既に `.gitignore` に含まれています。

## 公開API (Public API)

公開APIは認証不要で利用できます。

### REST API サンプル

- **status**: 取引所のステータス確認
- **orderbooks**: オーダーブック（板情報）の取得
- **trades**: 約定履歴の取得

#### 実行例

```bash
# ステータス確認
go run app/public/rest/status/main.go

# オーダーブック取得
go run app/public/rest/orderbooks/main.go

# 約定履歴取得
go run app/public/rest/trades/main.go
```

### WebSocket API サンプル

- **ticker**: ティッカー情報のリアルタイム受信
- **orderbooks**: オーダーブック更新のリアルタイム受信
- **trades**: 約定情報のリアルタイム受信

#### 実行例

```bash
# ティッカー情報の受信
go run app/public/ws/ticker/main.go

# オーダーブック更新の受信
go run app/public/ws/orderbooks/main.go

# 約定情報の受信
go run app/public/ws/trades/main.go
```

## プライベートAPI (Private API)

プライベートAPIを使用するには、`.env` ファイルにAPIキーとシークレットを設定する必要があります。

### REST API サンプル

- **account_margin**: 証拠金情報の取得
- **account_assets**: 資産残高の取得
- **orders**: 注文情報の取得
- **executions**: 約定情報の取得
- **last_executions**: 最新の約定情報の取得
- **position_summary**: ポジション情報のサマリー取得
- **change_order**: 注文の変更
- **ws_auth**: WebSocket認証トークンの管理

#### 実行例

```bash
# 証拠金情報の取得
go run app/private/rest/account_margin/main.go

# 資産残高の取得
go run app/private/rest/account_assets/main.go

# 注文情報の取得
go run app/private/rest/orders/main.go
```

### WebSocket API サンプル

- **execution_events**: 約定イベントのリアルタイム受信
- **order_events**: 注文イベントのリアルタイム受信
- **position_events**: ポジションイベントのリアルタイム受信
- **position_summary_events**: ポジションサマリーイベントのリアルタイム受信

#### 実行例

```bash
# 約定イベントの受信
go run app/private/ws/execution_events/main.go

# 注文イベントの受信
go run app/private/ws/order_events/main.go

# ポジションイベントの受信
go run app/private/ws/position_events/main.go
```

## コードの特徴

### 環境変数の自動読み込み

プライベートAPIのサンプルコードでは、`godotenv/autoload` を使用して `.env` ファイルから環境変数を自動的に読み込みます:

```go
import (
    "github.com/ijufumi/gogmocoin/v2/api/private/rest"
    _ "github.com/joho/godotenv/autoload"
)

func main() {
    // rest.New() は環境変数から自動的にAPI_KEYとAPI_SECRETを読み込みます
    client := rest.New()
    // ...
}
```

### エラーハンドリング

サンプルコードでは、適切なエラーハンドリングを実装しています:

- REST API: エラーが発生した場合は `log.Println` でログを出力してreturn
- WebSocket API: 共通の受信ループ（`internal/wsrunner`）を使用し、`Ctrl-C`（SIGINT/SIGTERM）を受信すると購読を解除してから安全に終了します

### WebSocket の共通受信ループ

WebSocket のサンプルは、購読・受信ループ・購読解除の共通処理を `internal/wsrunner` パッケージに集約しています。これにより各サンプルはクライアント生成と `wsrunner.Run` の呼び出しだけで完結し、シグナル受信時の graceful shutdown も共通化されています:

```go
client := ws.NewTicker(consts.SymbolBTCJPY)
if err := wsrunner.Run(client.Subscribe, client.Stream, client.Unsubscribe); err != nil {
    log.Fatal(err)
}
```

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。詳細は[LICENSE](LICENSE)ファイルを参照してください。

## リンク

- [gogmocoin ライブラリ](https://github.com/ijufumi/gogmocoin)
- [GMO Coin API ドキュメント](https://api.coin.z.com/docs/)

## コントリビューション

バグ報告や機能追加のリクエストは、GitHubのIssueでお願いします。Pull Requestも歓迎します。
