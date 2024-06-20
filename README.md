# ari

## バックエンド起動方法
```
1. aws cliを導入
2. AWS側で cloude(一応全部) と SDXLv1.0を使える状態にする. S3も (bucket-nameは`ai-phone`で)
3. docker compose up してmysql起動
4. go run cmd/app/main.go　でサーバ起動 (port :8080)