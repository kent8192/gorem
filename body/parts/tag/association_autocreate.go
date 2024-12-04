package tag

import "fmt"

// AssociationAutocreate構造体
type AssociationAutocreate struct {
	isLegitimatelyInitialized bool
	autoCreate                bool
}

// InitTagメソッドで自動作成フラグを初期化
func (a *AssociationAutocreate) InitTag(value bool) *AssociationAutocreate {
	a.isLegitimatelyInitialized = true
	a.autoCreate = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AssociationAutocreate) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("association_autocreate:%t", a.autoCreate)
}

// IsValidメソッドで自動作成フラグの妥当性をチェック
func (a *AssociationAutocreate) IsValid() (bool, string) {
	if !a.isLegitimatelyInitialized {
		return false, "association_autocreate is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AssociationAutocreate) ValidateTag() error {
	if valid, errMessage := a.IsValid(); !valid {
		return fmt.Errorf("invalid association_autocreate: %s", errMessage)
	}
	return nil
}

func (a *AssociationAutocreate) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
