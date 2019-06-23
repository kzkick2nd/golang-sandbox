# go build option の調査

```
# DWARF と シンボルテーブル 情報を削ると結構軽くなる
$ go build --ldflags="-s -w"

# upx かけるともっと小さくなる
$ upx --best --ultra-brute Binary
```
