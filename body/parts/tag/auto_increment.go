package tag

import (
	"fmt"
)

// AutoIncrement構造体
type AutoIncrement struct {
	isLegitimatelyInitialized bool
	value                     bool
}

// InitTagメソッドでAutoIncrementを初期化
func (a *AutoIncrement) InitTag(value bool) *AutoIncrement {
	a.isLegitimatelyInitialized = true
	a.value = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AutoIncrement) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	if a.value {
		return "autoIncrement:true"
	}
	return "autoIncrement:false"
}

// IsValidメソッドでAutoIncrementの妥当性をチェック
func (a *AutoIncrement) IsValid() (bool, string) {
	if !a.isLegitimatelyInitialized {
		return false, "autoIncrement is not initialized"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AutoIncrement) ValidateTag() error {
	if valid, errMsg := a.IsValid(); !valid {
		return fmt.Errorf("invalid autoIncrement: %s", errMsg)
	}
	return nil
}

func (a *AutoIncrement) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
