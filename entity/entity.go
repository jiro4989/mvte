package entity

type (
	WindowColor    int
	WindowPosition int
	ActorNameColor int
)

//go:generate stringer -type=WindowColor
//go:generate stringer -type=WindowPosition

const (
	WCWindow WindowColor = iota
	WCDark
	WCTransparent
	WPTop WindowPosition = iota
	WPMiddle
	WPBottom
)

type Message struct {
	ActorName      string
	Body           []string
	WindowColor    WindowColor
	WindowPosition WindowPosition
}

// Settings はMapXXXJSON出力設定
// See https://qiita.com/aosho235/items/ef7a99396f69442d2cf2
type Settings struct {
	ActorNameBracketStart    string         // アクター名を括る括弧(開き)
	ActorNameBracketEnd      string         // アクター名を括る括弧(閉じ)
	ActorNameInstantShowFlag bool           // アクター名の瞬間表示フラグ
	BodyBracketStart         string         // 本文を括る括弧(開き)
	BodyBracketEnd           string         // 本文を括る括弧(閉じ)
	Indent                   string         // インデント
	BracketAlign             bool           // 括弧の位置に揃えるか否か
	ActorNameColor           ActorNameColor // アクター名につける色。0未満の場合はつけない
	LineReturnLength         int            // 1行に表示する文字の長さ
}
