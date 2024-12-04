package tag

import "fmt"

type Size struct {
	isLegitimatelyInitialized bool
	size                      uint
}

func (s *Size) InitTag(value uint) *Size {
	s.isLegitimatelyInitialized = true
	s.size = value
	if err := s.ValidateTag(); err != nil {
		panic(err)
	}
	return s
}

func (s *Size) RenderTag() string {
	if !s.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("size:%d", s.size)
}

// IsValidメソッドでエラー内容を返すように修正
func (s *Size) IsValid() (bool, string) {
	if !s.isLegitimatelyInitialized {
		return false, "size is not initialized legitimately"
	}
	if s.size < 1 {
		return false, "size must be greater than or equal to 1"
	}
	return true, ""
}

func (s *Size) ValidateTag() error {
	// エラー内容をIsValidから受け取る
	if valid, errMessage := s.IsValid(); !valid {
		return fmt.Errorf("invalid size: %s (size=%d)", errMessage, s.size)
	}
	return nil
}

func (s *Size) GetIsLegitimatelyInitialized() bool {
	return s.isLegitimatelyInitialized
}
