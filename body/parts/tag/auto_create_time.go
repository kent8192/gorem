package tag

import "fmt"

// AutoCreateTime構造体
type AutoCreateTime struct {
	isLegitimatelyInitialized bool
	value                     string
}

// InitTagメソッドでAutoCreateTimeを初期化
func (a *AutoCreateTime) InitTag(value string) *AutoCreateTime {
	a.isLegitimatelyInitialized = true
	a.value = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AutoCreateTime) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	if a.value == "" {
		return "autoCreateTime"
	}
	return fmt.Sprintf("autoCreateTime:%s", a.value)
}

// IsValidメソッドでAutoCreateTimeの妥当性をチェック
func (a *AutoCreateTime) IsValid() (bool, string) {
	validValues := []string{"", "nano", "milli"}
	if !a.isLegitimatelyInitialized {
		return false, "autoCreateTime is not initialized"
	}
	for _, v := range validValues {
		if a.value == v {
			return true, ""
		}
	}
	return false, "autoCreateTime must be '', 'nano', or 'milli'"
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AutoCreateTime) ValidateTag() error {
	if valid, errMsg := a.IsValid(); !valid {
		return fmt.Errorf("invalid autoCreateTime: %s", errMsg)
	}
	return nil
}

func (a *AutoCreateTime) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
