package tag

import "fmt"

// WritePermission 構造体
type WritePermission struct {
	isLegitimatelyInitialized bool
	permission                string
}

// InitTag メソッドで書き込み権限を初期化
func (w *WritePermission) InitTag(value string) *WritePermission {
	w.isLegitimatelyInitialized = true
	w.permission = value
	return w
}

// RenderTag メソッドでタグ文字列を生成
func (w *WritePermission) RenderTag() string {
	if !w.isLegitimatelyInitialized {
		return ""
	}
	switch w.permission {
	case "create", "update", "false":
		return fmt.Sprintf("<-:%s", w.permission)
	case "":
		return "<-"
	default:
		return ""
	}
}

// IsValid メソッドで書き込み権限の妥当性をチェック
func (w *WritePermission) IsValid() (bool, string) {
	if !w.isLegitimatelyInitialized {
		return false, "write permission is not initialized legitimately"
	}
	validPermissions := map[string]bool{"create": true, "update": true, "false": true, "": true}
	if !validPermissions[w.permission] {
		return false, "invalid write permission value"
	}
	return true, ""
}

// ValidateTag メソッドでエラーメッセージを返す
func (w *WritePermission) ValidateTag() error {
	if valid, errMessage := w.IsValid(); !valid {
		return fmt.Errorf("invalid write permission: %s", errMessage)
	}
	return nil
}

func (w *WritePermission) GetIsLegitimatelyInitialized() bool {
	return w.isLegitimatelyInitialized
}
