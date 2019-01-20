package api

// WriteSettingsJSON はユーザ設定をファイル出力する
func WriteSettingsJSON(settings Settings) {

}

// WriteMapJSON はツクールMVのMapXXX.jsonを出力する
func WriteMapJSON(msgs []Message, settings Settings) {

}

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
