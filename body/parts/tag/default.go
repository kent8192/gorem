package tag

import "fmt"

type Default struct {
	isLegitimatelyInitialized bool
	defaultValue              string
}

// InitTagメソッドでデフォルト値を初期化
func (d *Default) InitTag(value string) *Default {
	d.isLegitimatelyInitialized = true
	d.defaultValue = value
	if err := d.ValidateTag(); err != nil {
		panic(err)
	}
	return d
}

// RenderTagメソッドでタグ文字列を生成
func (d *Default) RenderTag() string {
	if !d.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("default:'%s'", d.defaultValue)
}

// IsValidメソッドでデフォルト値の妥当性をチェック
func (d *Default) IsValid() (bool, string) {
	if !d.isLegitimatelyInitialized {
		return false, "default is not initialized legitimately"
	}
	if d.defaultValue == "" {
		return false, "default value cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (d *Default) ValidateTag() error {
	if valid, errMessage := d.IsValid(); !valid {
		return fmt.Errorf("invalid default: %s", errMessage)
	}
	return nil
}

func (c *Default) GetIsLegitimatelyInitialized() bool {
	return c.isLegitimatelyInitialized
}
