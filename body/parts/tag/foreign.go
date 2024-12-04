package tag

import "fmt"

// ForeignKey構造体
type ForeignKey struct {
	isLegitimatelyInitialized bool
	fieldName                 string
}

// InitTagメソッドで外部キーのフィールド名を初期化
func (f *ForeignKey) InitTag(value string) *ForeignKey {
	f.isLegitimatelyInitialized = true
	f.fieldName = value
	return f
}

// RenderTagメソッドでタグ文字列を生成
func (f *ForeignKey) RenderTag() string {
	if !f.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("foreignKey:%s", f.fieldName)
}

// IsValidメソッドで外部キーの妥当性をチェック
func (f *ForeignKey) IsValid() (bool, string) {
	if !f.isLegitimatelyInitialized {
		return false, "foreignKey is not initialized legitimately"
	}
	if f.fieldName == "" {
		return false, "foreignKey cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (f *ForeignKey) ValidateTag() error {
	if valid, errMessage := f.IsValid(); !valid {
		return fmt.Errorf("invalid foreignKey: %s", errMessage)
	}
	return nil
}

func (f *ForeignKey) GetIsLegitimatelyInitialized() bool {
	return f.isLegitimatelyInitialized
}
