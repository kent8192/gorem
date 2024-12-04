package tag

import "fmt"

type NotNull struct {
	isLegitimatelyInitialized bool
	notNull                   bool
}

// InitTagメソッドでNotNullを初期化
func (n *NotNull) InitTag(value bool) *NotNull {
	n.isLegitimatelyInitialized = true
	n.notNull = value
	if err := n.ValidateTag(); err != nil {
		panic(err)
	}
	return n
}

// RenderTagメソッドでタグ文字列を生成
func (n *NotNull) RenderTag() string {
	if !n.isLegitimatelyInitialized {
		return ""
	}
	if n.notNull {
		return "not null"
	}
	return ""
}

// IsValidメソッドでNotNullの妥当性をチェック
func (n *NotNull) IsValid() (bool, string) {
	if !n.isLegitimatelyInitialized {
		return false, "NotNull is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (n *NotNull) ValidateTag() error {
	if valid, errMessage := n.IsValid(); !valid {
		return fmt.Errorf("invalid NotNull: %s", errMessage)
	}
	return nil
}

func (n *NotNull) GetIsLegitimatelyInitialized() bool {
	return n.isLegitimatelyInitialized
}
