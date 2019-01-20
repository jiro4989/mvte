package api

import (
	"fmt"
	"io"

	"github.com/jiro4989/tmpl-go/entity"
)

// WriteSettingsJSON はユーザ設定をファイル出力する
func WriteSettingsJSON(settings Settings) {

}

// WriteMapJSON はツクールMVのMapXXX.jsonを出力する
func WriteMapJSON(msgs []Message, settings Settings) {

}

func writeMap(w io.Writer, msgs []entity.Message, settings entity.Settings) {
	for _, msg := range msgs {
		// Messageのアクター名を括弧でくくる
		msg.ActorName = settings.ActorNameBracketStart + msg.ActorName + settings.ActorNameBracketEnd

		// Messageのアクター名の色を変更する
		colIdx := settings.ActorNameColor
		if 0 < colIdx {
			msg.ActorName = surroundColor(msg.ActorName, colIdx)
		}
	}
	// Messageの1行文字列が特定の長さ以上だった場合に改行(配列をわける)
	// Messageのアクター名とBodyが合わせて4行以上になった場合はMessageを分ける
	// Bodyを括弧で括る。括る際はインデントを揃える可能性がある
	// Bodyを括弧で括らない場合もインデントするかも知れない
	// ファイル出力用のツクールMV規定のJSON構造体に変換する
	// writerに書き込む
}

// surroundColor は文字列をツクール仕様の色文字で囲う
// デフォルト色：\c[0]
// 他：\c[1]
func surroundColor(s string, n int) string {
	return fmt.Sprintf(`\c[%d]%s\c[0]`, n, s)
}

// joinMessages は複数のテーブルのテキストを可能な限り結合する。
func joinMessages(r io.Writer, msgs []entity.Message) {

}
