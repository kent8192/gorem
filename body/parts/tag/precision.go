package tag

import (
	"fmt"
)

type Precision struct {
	isLegitimatelyInitialized bool
	precision                 int
}

// InitTagメソッドで精度を初期化
func (p *Precision) InitTag(value int) *Precision {
	p.isLegitimatelyInitialized = true
	p.precision = value
	if err := p.ValidateTag(); err != nil {
		panic(err)
	}
	return p
}

// RenderTagメソッドでタグ文字列を生成
func (p *Precision) RenderTag() string {
	if !p.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("precision:%d", p.precision)
}

// IsValidメソッドで精度の妥当性をチェック
func (p *Precision) IsValid() (bool, string) {
	if !p.isLegitimatelyInitialized {
		return false, "precision is not initialized legitimately"
	}
	if p.precision < 0 {
		return false, "precision cannot be negative"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (p *Precision) ValidateTag() error {
	if valid, errMessage := p.IsValid(); !valid {
		return fmt.Errorf("invalid precision: %s", errMessage)
	}
	return nil
}

func (p *Precision) GetIsLegitimatelyInitialized() bool {
	return p.isLegitimatelyInitialized
}
