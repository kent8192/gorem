package tag

import "fmt"

// ReadPermission 構造体
type ReadPermission struct {
	isLegitimatelyInitialized bool
	permission                string
}

// InitTag メソッドで読み取り権限を初期化
func (r *ReadPermission) InitTag(value string) *ReadPermission {
	r.isLegitimatelyInitialized = true
	r.permission = value
	if err := r.ValidateTag(); err != nil {
		panic(err)
	}
	return r
}

// RenderTag メソッドでタグ文字列を生成
func (r *ReadPermission) RenderTag() string {
	if !r.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("->:%s", r.permission)
}

// IsValid メソッドで読み取り権限の妥当性をチェック
func (r *ReadPermission) IsValid() (bool, string) {
	if !r.isLegitimatelyInitialized {
		return false, "read permission is not initialized legitimately"
	}
	return true, ""
}

// ValidateTag メソッドでエラーメッセージを返す
func (r *ReadPermission) ValidateTag() error {
	if valid, errMessage := r.IsValid(); !valid {
		return fmt.Errorf("invalid read permission: %s", errMessage)
	}
	return nil
}

func (r *ReadPermission) GetIsLegitimatelyInitialized() bool {
	return r.isLegitimatelyInitialized
}
