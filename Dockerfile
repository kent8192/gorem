# ビルドステージ：Go 1.23.3を使用してアプリケーションをビルド
FROM golang:1.23.3-bookworm AS builder

# 作業ディレクトリを設定
WORKDIR /gorem

# 依存関係ファイルをコピー
COPY go.mod ./

RUN go mod tidy

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .
