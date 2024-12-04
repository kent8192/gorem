package tag

import "fmt"

// AutoUpdateTime構造体
type AutoUpdateTime struct {
	isLegitimatelyInitialized bool
	value                     string
}

// InitTagメソッドでAutoUpdateTimeを初期化
func (a *AutoUpdateTime) InitTag(value string) *AutoUpdateTime {
	a.isLegitimatelyInitialized = true
	a.value = value
	return a
}

// RenderTagメソッドでタグ文字列を生成
func (a *AutoUpdateTime) RenderTag() string {
	if !a.isLegitimatelyInitialized {
		return ""
	}
	if a.value == "" {
		return "autoUpdateTime"
	}
	return fmt.Sprintf("autoUpdateTime:%s", a.value)
}

// IsValidメソッドでAutoUpdateTimeの妥当性をチェック
func (a *AutoUpdateTime) IsValid() (bool, string) {
	validValues := []string{"", "nano", "milli"}
	if !a.isLegitimatelyInitialized {
		return false, "autoUpdateTime is not initialized"
	}
	for _, v := range validValues {
		if a.value == v {
			return true, ""
		}
	}
	return false, "autoUpdateTime must be '', 'nano', or 'milli'"
}

// ValidateTagメソッドでエラーメッセージを返す
func (a *AutoUpdateTime) ValidateTag() error {
	if valid, errMsg := a.IsValid(); !valid {
		return fmt.Errorf("invalid autoUpdateTime: %s", errMsg)
	}
	return nil
}

func (a *AutoUpdateTime) GetIsLegitimatelyInitialized() bool {
	return a.isLegitimatelyInitialized
}
