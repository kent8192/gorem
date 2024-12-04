package tag

import "fmt"

// AssociationAutoupdate構造体
type AssociationAutoupdate struct {
	isLegitimatelyInitialized bool
	autoUpdate                bool
}

// InitTagメソッドで自動更新フラグを初期化
func (a *AssociationAutoupdate) InitTag(value bool) *AssociationAutoupdate {
	a.isLegitimatelyInitialized = true
	a.autoUpdate = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AssociationAutoupdate) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("association_autoupdate:%t", a.autoUpdate)
}

// IsValidメソッドで自動更新フラグの妥当性をチェック
func (a *AssociationAutoupdate) IsValid() (bool, string) {
	if !a.isLegitimatelyInitialized {
		return false, "association_autoupdate is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AssociationAutoupdate) ValidateTag() error {
	if valid, errMessage := a.IsValid(); !valid {
		return fmt.Errorf("invalid association_autoupdate: %s", errMessage)
	}
	return nil
}

func (a *AssociationAutoupdate) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
