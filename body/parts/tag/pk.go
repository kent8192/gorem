package tag

import (
	"fmt"
)

// PrimaryKey構造体は、primaryKeyタグの情報を保持します。
type PrimaryKey struct {
	isLegitimatelyInitialized bool
	isPrimaryKey              bool
}

// InitTagメソッドは、primaryKeyタグを初期化します。
func (p *PrimaryKey) InitTag(value bool) *PrimaryKey {
	p.isLegitimatelyInitialized = true
	p.isPrimaryKey = value
	if err := p.ValidateTag(); err != nil {
		panic(err)
	}
	return p
}

// RenderTagメソッドは、primaryKeyタグの文字列を生成します。
func (p *PrimaryKey) RenderTag() string {
	if !p.isLegitimatelyInitialized {
		return ""
	}
	if p.isPrimaryKey {
		return "primaryKey"
	}
	return ""
}

// IsValidメソッドは、primaryKeyタグの妥当性をチェックします。
func (p *PrimaryKey) IsValid() (bool, string) {
	if !p.isLegitimatelyInitialized {
		return false, "primaryKey is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドは、primaryKeyタグのエラーメッセージを返します。
func (p *PrimaryKey) ValidateTag() error {
	if valid, errMessage := p.IsValid(); !valid {
		return fmt.Errorf("invalid primaryKey: %s", errMessage)
	}
	return nil
}

func (p *PrimaryKey) GetIsLegitimatelyInitialized() bool {
	return p.isLegitimatelyInitialized
}
