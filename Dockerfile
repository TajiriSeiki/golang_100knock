# Go の公式イメージを使用
FROM golang:1.20

# 作業ディレクトリを設定
WORKDIR /app

# アプリケーションコードをコンテナ内にコピー
COPY . .

# 必要に応じてモジュールをダウンロード
RUN go mod init myapp && go mod tidy

# デフォルトの実行コマンド（任意）
CMD ["go", "run", "main.go"]