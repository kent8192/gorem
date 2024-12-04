package tag

import "fmt"

type Unique struct {
	isLegitimatelyInitialized bool
	unique                    bool
}

// InitTagメソッドでユニーク制約を初期化
func (u *Unique) InitTag(value bool) *Unique {
	u.isLegitimatelyInitialized = true
	u.unique = value
	if err := u.ValidateTag(); err != nil {
		panic(err)
	}
	return u
}

// RenderTagメソッドでタグ文字列を生成
func (u *Unique) RenderTag() string {
	if !u.isLegitimatelyInitialized {
		return ""
	}
	if u.unique {
		return "unique"
	}
	return ""
}

// IsValidメソッドでユニーク制約の妥当性をチェック
func (u *Unique) IsValid() (bool, string) {
	if !u.isLegitimatelyInitialized {
		return false, "unique is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (u *Unique) ValidateTag() error {
	if valid, errMessage := u.IsValid(); !valid {
		return fmt.Errorf("invalid unique: %s", errMessage)
	}
	return nil
}

func (r *Unique) GetIsLegitimatelyInitialized() bool {
	return r.isLegitimatelyInitialized
}
