# mvte

MVTextEditorのGo移植

## 実装方針

- GUIはWebブラウザを使用する
  - localhostサーバでの実装
- CLIとしても使える
- ワンバイナリで動作
  - HTML, JS, CSSはgo-assetsでうめこむ

## ツクールMVの仕様

メッセージ構造

- parametersのとこに1行テキストを持つ
- 101codeから4行分のテキストを持って1セット

    {
      "code": 101,
      "indent": 0,
      "parameters": [
        "",
        0,
        0,
        2
      ]
    },
    {
      "code": 401,
      "indent": 0,
      "parameters": [
        "\\>ナビゲータ"
      ]
    },
    {
      "code": 401,
      "indent": 0,
      "parameters": [
        "「初めまして。私は本ソフト(MV Text Editor)の使い方のう"
      ]
    },
    {
      "code": 401,
      "indent": 0,
      "parameters": [
        "  ち、テキスト書式について説明させていただきます。ナビ"
      ]
    },
    {
      "code": 401,
      "indent": 0,
      "parameters": [
        "  ゲータと申します。」"
      ]
    },
