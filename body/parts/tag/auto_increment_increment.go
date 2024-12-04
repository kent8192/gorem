package tag

import "fmt"

// AutoIncrementIncrement構造体
type AutoIncrementIncrement struct {
	isLegitimatelyInitialized bool
	value                     int
}

// InitTagメソッドでAutoIncrementIncrementを初期化
func (a *AutoIncrementIncrement) InitTag(value int) *AutoIncrementIncrement {
	a.isLegitimatelyInitialized = true
	a.value = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AutoIncrementIncrement) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("autoIncrementIncrement:%d", a.value)
}

// IsValidメソッドでAutoIncrementIncrementの妥当性をチェック
func (a *AutoIncrementIncrement) IsValid() (bool, string) {
	if !a.isLegitimatelyInitialized {
		return false, "autoIncrementIncrement is not initialized"
	}
	if a.value <= 0 {
		return false, "autoIncrementIncrement must be a positive integer"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AutoIncrementIncrement) ValidateTag() error {
	if valid, errMsg := a.IsValid(); !valid {
		return fmt.Errorf("invalid autoIncrementIncrement: %s", errMsg)
	}
	return nil
}

func (a *AutoIncrementIncrement) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
