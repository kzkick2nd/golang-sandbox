# go build option の調査

```
# DWARF と シンボルテーブル 情報を削ると結構軽くなる
$ go build --ldflags="-s -w"

# upx かけるともっと小さくなる
$ upx --best --ultra-brute Binary
```

# Reference
[Building smaller Go binaries - Javier Provecho Fernández at dotGo 2016](https://www.dotconferences.com/2016/10/javier-provecho-fernandez-building-smaller-go-binaries)