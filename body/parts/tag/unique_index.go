package tag

import "fmt"

// UniqueIndex 構造体
type UniqueIndex struct {
	isLegitimatelyInitialized bool
	name                      string
	sort                      string
}

// InitTag メソッドでユニークインデックス名を初期化
func (u *UniqueIndex) InitTag(name string, sort string) *UniqueIndex {
	u.isLegitimatelyInitialized = true
	u.name = name
	u.sort = sort
	return u
}

// RenderTag メソッドでタグ文字列を生成
func (u *UniqueIndex) RenderTag() string {
	if !u.isLegitimatelyInitialized {
		return ""
	}
	if u.sort != "" {
		return fmt.Sprintf("uniqueIndex:%s,sort:%s", u.name, u.sort)
	}
	return fmt.Sprintf("uniqueIndex:%s", u.name)
}

// IsValid メソッドでユニークインデックスの妥当性をチェック
func (u *UniqueIndex) IsValid() (bool, string) {
	if !u.isLegitimatelyInitialized {
		return false, "uniqueIndex is not initialized legitimately"
	}
	if u.name == "" {
		return false, "uniqueIndex name cannot be empty"
	}
	if u.sort != "" && u.sort != "asc" && u.sort != "desc" {
		return false, "sort must be 'asc' or 'desc' if specified"
	}
	return true, ""
}

// ValidateTag メソッドでエラーメッセージを返す
func (u *UniqueIndex) ValidateTag() error {
	if valid, errMessage := u.IsValid(); !valid {
		return fmt.Errorf("invalid uniqueIndex: %s", errMessage)
	}
	return nil
}

func (i *UniqueIndex) GetIsLegitimatelyInitialized() bool {
	return i.isLegitimatelyInitialized
}
