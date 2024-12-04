package tag

import (
	"fmt"
)

// Embedded構造体は、フィールドを埋め込みフィールドとして扱うことを示します。
type Embedded struct {
	isLegitimatelyInitialized bool
}

// InitTagメソッドでEmbeddedタグを初期化
func (e *Embedded) InitTag(value bool) *Embedded {
	e.isLegitimatelyInitialized = value
	if err := e.ValidateTag(); err != nil {
		panic(err)
	}
	return e
}

// RenderTagメソッドでタグ文字列を生成
func (e *Embedded) RenderTag() string {
	if !e.isLegitimatelyInitialized {
		return ""
	}
	return "embedded"
}

// IsValidメソッドでEmbeddedタグの妥当性をチェック
func (e *Embedded) IsValid() (bool, string) {
	if !e.isLegitimatelyInitialized {
		return false, "embedded tag is not initialized"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (e *Embedded) ValidateTag() error {
	if valid, errMessage := e.IsValid(); !valid {
		return fmt.Errorf("invalid embedded tag: %s", errMessage)
	}
	return nil
}

func (e *Embedded) GetIsLegitimatelyInitialized() bool {
	return e.isLegitimatelyInitialized
}
