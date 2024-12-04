package tag

import "fmt"

type Comment struct {
	isLegitimatelyInitialized bool
	comment                   string
}

// InitTagメソッドでコメントを初期化
func (c *Comment) InitTag(value string) *Comment {
	c.isLegitimatelyInitialized = true
	c.comment = value
	if err := c.ValidateTag(); err != nil {
		panic(err)
	}
	return c
}

// RenderTagメソッドでタグ文字列を生成
func (c *Comment) RenderTag() string {
	if !c.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("comment:'%s'", c.comment)
}

// IsValidメソッドでコメントの妥当性をチェック
func (c *Comment) IsValid() (bool, string) {
	if !c.isLegitimatelyInitialized {
		return false, "comment is not initialized legitimately"
	}
	if c.comment == "" {
		return false, "comment cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (c *Comment) ValidateTag() error {
	if valid, errMessage := c.IsValid(); !valid {
		return fmt.Errorf("invalid comment: %s", errMessage)
	}
	return nil
}

func (c *Comment) GetIsLegitimatelyInitialized() bool {
	return c.isLegitimatelyInitialized
}
