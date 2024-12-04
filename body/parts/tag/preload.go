package tag

import "fmt"

type Preload struct {
	isLegitimatelyInitialized bool
	preload                   bool
}

// InitTagメソッドでpreloadを初期化
func (p *Preload) InitTag(value bool) *Preload {
	p.isLegitimatelyInitialized = true
	p.preload = value
	if err := p.ValidateTag(); err != nil {
		panic(err)
	}
	return p
}

// RenderTagメソッドでタグ文字列を生成
func (p *Preload) RenderTag() string {
	if !p.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("preload:%t", p.preload)
}

// IsValidメソッドでpreloadの妥当性をチェック
func (p *Preload) IsValid() (bool, string) {
	if !p.isLegitimatelyInitialized {
		return false, "preload is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (p *Preload) ValidateTag() error {
	if valid, errMessage := p.IsValid(); !valid {
		return fmt.Errorf("invalid preload: %s", errMessage)
	}
	return nil
}

func (p *Preload) GetIsLegitimatelyInitialized() bool {
	return p.isLegitimatelyInitialized
}
