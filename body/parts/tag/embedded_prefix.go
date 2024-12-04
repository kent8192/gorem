package tag

import "fmt"

// EmbeddedPrefix構造体は、埋め込みフィールドのカラム名に付加するプレフィックスを指定します。
type EmbeddedPrefix struct {
	isLegitimatelyInitialized bool
	prefix                    string
}

// InitTagメソッドでEmbeddedPrefixタグを初期化
func (ep *EmbeddedPrefix) InitTag(value string) *EmbeddedPrefix {
	ep.isLegitimatelyInitialized = true
	ep.prefix = value
	return ep
}

// RenderTagメソッドでタグ文字列を生成
func (ep *EmbeddedPrefix) RenderTag() string {
	return fmt.Sprintf("embeddedPrefix:'%s'", ep.prefix)
}

// IsValidメソッドでEmbeddedPrefixタグの妥当性をチェック
func (ep *EmbeddedPrefix) IsValid() (bool, string) {
	if !ep.isLegitimatelyInitialized {
		return false, "embeddedPrefix tag is not initialized"
	}
	if ep.prefix == "" {
		return false, "embeddedPrefix cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (ep *EmbeddedPrefix) ValidateTag() error {
	if valid, errMessage := ep.IsValid(); !valid {
		return fmt.Errorf("invalid embeddedPrefix tag: %s", errMessage)
	}
	return nil
}

func (e *EmbeddedPrefix) GetIsLegitimatelyInitialized() bool {
	return e.isLegitimatelyInitialized
}
