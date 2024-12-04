package tag

import (
	"fmt"
	"strings"
)

// Many2Many構造体は、GORMのmany2manyタグを管理します。
type Many2Many struct {
	isLegitimatelyInitialized bool
	tableName                 string
}

// InitTagメソッドは、many2manyタグを初期化します。
func (m *Many2Many) InitTag(value string) *Many2Many {
	m.isLegitimatelyInitialized = true
	m.tableName = value
	return m
}

// RenderTagメソッドは、many2manyタグの文字列を生成します。
func (m *Many2Many) RenderTag() string {
	if !m.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("many2many:%s", m.tableName)
}

// IsValidメソッドは、many2manyタグの妥当性をチェックします。
func (m *Many2Many) IsValid() (bool, string) {
	if !m.isLegitimatelyInitialized {
		return false, "many2many tag is not initialized legitimately"
	}
	if strings.TrimSpace(m.tableName) == "" {
		return false, "table name cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドは、many2manyタグのバリデーションを行い、エラーメッセージを返します。
func (m *Many2Many) ValidateTag() error {
	if valid, errMessage := m.IsValid(); !valid {
		return fmt.Errorf("invalid many2many tag: %s", errMessage)
	}
	return nil
}

func (m *Many2Many) GetIsLegitimatelyInitialized() bool {
	return m.isLegitimatelyInitialized
}
