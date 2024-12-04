package tag

import "fmt"

// AssociationSaveReference構造体
type AssociationSaveReference struct {
	isLegitimatelyInitialized bool
	saveReference             bool
}

// InitTagメソッドで参照保存フラグを初期化
func (a *AssociationSaveReference) InitTag(value bool) *AssociationSaveReference {
	a.isLegitimatelyInitialized = true
	a.saveReference = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AssociationSaveReference) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("association_save_reference:%t", a.saveReference)
}

// IsValidメソッドで参照保存フラグの妥当性をチェック
func (a *AssociationSaveReference) IsValid() (bool, string) {
	if !a.isLegitimatelyInitialized {
		return false, "association_save_reference is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AssociationSaveReference) ValidateTag() error {
	if valid, errMessage := a.IsValid(); !valid {
		return fmt.Errorf("invalid association_save_reference: %s", errMessage)
	}
	return nil
}

func (a *AssociationSaveReference) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
