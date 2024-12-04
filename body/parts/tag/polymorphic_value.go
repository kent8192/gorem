package tag

import "fmt"

// PolymorphicValueタグを表現する構造体
type PolymorphicValue struct {
	isLegitimatelyInitialized bool
	value                     string
}

// InitTagメソッドでPolymorphicValueタグを初期化
func (pv *PolymorphicValue) InitTag(value string) *PolymorphicValue {
	pv.isLegitimatelyInitialized = true
	pv.value = value
	return pv
}

// RenderTagメソッドでタグ文字列を生成
func (pv *PolymorphicValue) RenderTag() string {
	if !pv.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("polymorphicValue:'%s'", pv.value)
}

// IsValidメソッドでPolymorphicValueタグの妥当性をチェック
func (pv *PolymorphicValue) IsValid() (bool, string) {
	if !pv.isLegitimatelyInitialized {
		return false, "polymorphicValue tag is not initialized legitimately"
	}
	if pv.value == "" {
		return false, "polymorphicValue cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (pv *PolymorphicValue) ValidateTag() error {
	if valid, errMessage := pv.IsValid(); !valid {
		return fmt.Errorf("invalid polymorphicValue tag: %s", errMessage)
	}
	return nil
}

func (p *PolymorphicValue) GetIsLegitimatelyInitialized() bool {
	return p.isLegitimatelyInitialized
}
