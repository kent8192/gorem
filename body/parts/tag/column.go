package tag

import (
	"fmt"
)

type Column struct {
	column                    string
	isLegitimatelyInitialized bool
}

// 初期化メソッド
func (c *Column) InitTag(value string) *Column {
	c.isLegitimatelyInitialized = true
	c.column = value
	if err := c.ValidateTag(); err != nil {
		panic(err)
	}
	return c
}

// レンダリングメソッド
func (c *Column) RenderTag() string {
	if !c.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("column:%s", c.column)
}

// IsValidメソッドでエラー内容を返すように修正
func (c *Column) IsValid() (bool, string) {
	if !c.isLegitimatelyInitialized {
		return false, "column is not initialized legitimately"
	}
	if c.column == "" {
		return false, "column name cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッド
func (c *Column) ValidateTag() error {
	// エラー内容をIsValidから受け取る
	if valid, errMessage := c.IsValid(); !valid {
		return fmt.Errorf("invalid column: %s (column=%s)", errMessage, c.column)
	}
	return nil
}

func (c *Column) GetIsLegitimatelyInitialized() bool {
	return c.isLegitimatelyInitialized
}
