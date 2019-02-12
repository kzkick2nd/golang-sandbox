package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// TODO
// 1. json を返す OK
// 2. パラメーターを受け取る OK
// 3. テスト方法と設計
// https://golang.org/pkg/net/http/httptest/
// 4. Handler インターフェース [http - The Go Programming Language](https://golang.org/pkg/net/http/#Handler)
// 5. アクセスログ？
// 6. ダウンタイム？

// JSON 作り方
// フォーマット ベストプラクティス
// ステータス ベストプラクティス

func main() {

	http.HandleFunc("/params", parseParams)

	http.HandleFunc("/api/clock", apiClockHandler)

	http.Handle("/assets/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "world")
	})

	f := fuga(1)
	http.Handle("/", &f)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseParams(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", r.URL.Query())
}

func apiClockHandler(w http.ResponseWriter, r *http.Request) {
	// JSONにする構造体
	type ResponseBody struct {
		Time time.Time `json:"time"`
	}
	rb := &ResponseBody{
		Time: time.Now(),
	}

	// ヘッダーをセット
	w.Header().Set("Content-type", "application/json")

	// JSONにエンコードしてレスポンスに書き込む
	if err := json.NewEncoder(w).Encode(rb); err != nil {
		log.Fatal(err)
	}
}

type fuga int

func (f *fuga) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fuga type: %d", *f)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}
