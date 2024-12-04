package tag

import "fmt"

// AssociationForeignKey構造体
type AssociationForeignKey struct {
	isLegitimatelyInitialized bool
	foreignKey                string
}

// InitTagメソッドで外部キーを初期化
func (a *AssociationForeignKey) InitTag(value string) *AssociationForeignKey {
	a.isLegitimatelyInitialized = true
	a.foreignKey = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AssociationForeignKey) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("associationForeignKey:%s", a.foreignKey)
}

// IsValidメソッドで外部キーの妥当性をチェック
func (a *AssociationForeignKey) IsValid() (bool, string) {
	if !a.isLegitimatelyInitialized {
		return false, "associationForeignKey is not initialized legitimately"
	}
	if a.foreignKey == "" {
		return false, "associationForeignKey cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AssociationForeignKey) ValidateTag() error {
	if valid, errMessage := a.IsValid(); !valid {
		return fmt.Errorf("invalid associationForeignKey: %s", errMessage)
	}
	return nil
}

func (a *AssociationForeignKey) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
