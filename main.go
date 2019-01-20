package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// /Hello/:langにハンドルされているHello関数
func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	lang := p.ByName("lang") // langパラメーターを取得する
	fmt.Fprintf(w, lang)     // レスポンスに値を書き込む
}

// /ExampleにハンドルされているExample関数
func Example(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close() // Example関数が終了する時に実行されるdeferステートメント

	// リクエストボディを読み取る
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// リクエストボディの読み取りに失敗した => 400 Bad Requestエラー
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// JSONパラメーターを構造体にする為の定義
	type ExampleParameter struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var param ExampleParameter

	// ExampleParameter構造体に変換
	err = json.Unmarshal(bodyBytes, &param)
	if err != nil {
		// JSONパラメーターを構造体への変換に失敗した => 400 Bad Requestエラー
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 構造体に変換したExampleParameterを文字列にしてレスポンスに書き込む
	fmt.Fprintf(w, fmt.Sprintf("%+v\n", param))
}

func main() {
	router := httprouter.New() // HTTPルーターを初期化

	// /HelloにGETリクエストがあったらHello関数にハンドルする
	// :langはパラメーターとして扱われる
	router.GET("/Hello/:lang", Hello)

	// /ExampleにPOSTリクエストがあったらExample関数にハンドルする
	router.POST("/Example", Example)
	router.PATCH("/api/v1/map", Example)       // MapXXX.jsonの更新
	router.GET("/api/v1/settings", Example)    // ユーザ設定の取得
	router.POST("/api/v1/settings", Example)   // ユーザ設定の作成
	router.PATCH("/api/v1/settings", Example)  // ユーザ設定の更新
	router.DELETE("/api/v1/settings", Example) // ユーザ設定の削除

	// Webサーバーを8080ポートで立ち上げる
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
