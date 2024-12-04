package tag

import "fmt"

type References struct {
	isLegitimatelyInitialized bool
	referenceField            string
}

// InitTagメソッドで参照フィールドを初期化
func (r *References) InitTag(value string) *References {
	r.isLegitimatelyInitialized = true
	r.referenceField = value
	return r
}

// RenderTagメソッドでタグ文字列を生成
func (r *References) RenderTag() string {
	if !r.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("references:%s", r.referenceField)
}

// IsValidメソッドで参照フィールドの妥当性をチェック
func (r *References) IsValid() (bool, string) {
	if !r.isLegitimatelyInitialized {
		return false, "references tag is not initialized legitimately"
	}
	if r.referenceField == "" {
		return false, "reference field cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (r *References) ValidateTag() error {
	if valid, errMessage := r.IsValid(); !valid {
		return fmt.Errorf("invalid references tag: %s", errMessage)
	}
	return nil
}

func (r *References) GetIsLegitimatelyInitialized() bool {
	return r.isLegitimatelyInitialized
}
