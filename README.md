# mvte

MVTextEditorのGo移植

## 実装方針

- GUIはWebブラウザを使用する
  - localhostサーバでの実装
- CLIとしても使える
- ワンバイナリで動作
  - HTML, JS, CSSはgo-assetsでうめこむ
