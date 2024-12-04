package tag

import "fmt"

// Polymorphicタグを表現する構造体
type Polymorphic struct {
	isLegitimatelyInitialized bool
	fieldName                 string
}

// InitTagメソッドでPolymorphicタグを初期化
func (p *Polymorphic) InitTag(value string) *Polymorphic {
	p.isLegitimatelyInitialized = true
	p.fieldName = value
	if err := p.ValidateTag(); err != nil {
		panic(err)
	}
	return p
}

// RenderTagメソッドでタグ文字列を生成
func (p *Polymorphic) RenderTag() string {
	if !p.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("polymorphic:%s", p.fieldName)
}

// IsValidメソッドでPolymorphicタグの妥当性をチェック
func (p *Polymorphic) IsValid() (bool, string) {
	if !p.isLegitimatelyInitialized {
		return false, "polymorphic tag is not initialized legitimately"
	}
	if p.fieldName == "" {
		return false, "polymorphic field name cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (p *Polymorphic) ValidateTag() error {
	if valid, errMessage := p.IsValid(); !valid {
		return fmt.Errorf("invalid polymorphic tag: %s", errMessage)
	}
	return nil
}

func (p *Polymorphic) GetIsLegitimatelyInitialized() bool {
	return p.isLegitimatelyInitialized
}
