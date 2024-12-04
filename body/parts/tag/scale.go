package tag

import "fmt"

// Scale構造体の定義
type Scale struct {
	isLegitimatelyInitialized bool
	scale                     uint
}

// InitTagメソッド: Scaleの初期化
func (s *Scale) InitTag(value uint) *Scale {
	s.isLegitimatelyInitialized = true
	s.scale = value
	if err := s.ValidateTag(); err != nil {
		panic(err)
	}
	return s
}

// RenderTagメソッド: Scaleタグの文字列表現を生成
func (s *Scale) RenderTag() string {
	if !s.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("scale:%d", s.scale)
}

// IsValidメソッド: Scaleのバリデーションとエラー内容を返す
func (s *Scale) IsValid() (bool, string) {
	if !s.isLegitimatelyInitialized {
		return false, "scale is not initialized legitimately"
	}
	return true, ""
}

// ValidateTagメソッド: バリデーションの結果をエラーとして返す
func (s *Scale) ValidateTag() error {
	// エラー内容をIsValidから受け取る
	if valid, errMessage := s.IsValid(); !valid {
		return fmt.Errorf("invalid scale: %s (scale=%d)", errMessage, s.scale)
	}
	return nil
}

func (s *Scale) GetIsLegitimatelyInitialized() bool {
	return s.isLegitimatelyInitialized
}
