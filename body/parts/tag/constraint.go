package tag

import (
	"fmt"
	"strings"
)

// Constraint構造体は、GORMのconstraintタグを管理します。
type Constraint struct {
	isLegitimatelyInitialized bool
	onUpdate                  string
	onDelete                  string
}

// InitTagメソッドは、OnUpdateおよびOnDeleteの制約を初期化します。
func (c *Constraint) InitTag(onUpdate, onDelete string) *Constraint {
	c.isLegitimatelyInitialized = true
	c.onUpdate = onUpdate
	c.onDelete = onDelete
	return c
}

// RenderTagメソッドは、constraintタグの文字列を生成します。
func (c *Constraint) RenderTag() string {
	if !c.isLegitimatelyInitialized {
		return ""
	}
	var constraints []string
	if c.onUpdate != "" {
		constraints = append(constraints, fmt.Sprintf("OnUpdate:%s", c.onUpdate))
	}
	if c.onDelete != "" {
		constraints = append(constraints, fmt.Sprintf("OnDelete:%s", c.onDelete))
	}
	return fmt.Sprintf("constraint:%s;", strings.Join(constraints, ","))
}

// IsValidメソッドは、設定された制約の妥当性をチェックします。
func (c *Constraint) IsValid() (bool, string) {
	if !c.isLegitimatelyInitialized {
		return false, "constraint is not initialized legitimately"
	}
	validActions := map[string]bool{
		"CASCADE":   true,
		"RESTRICT":  true,
		"SET NULL":  true,
		"NO ACTION": true,
	}
	if c.onUpdate != "" && !validActions[c.onUpdate] {
		return false, fmt.Sprintf("invalid OnUpdate action: %s", c.onUpdate)
	}
	if c.onDelete != "" && !validActions[c.onDelete] {
		return false, fmt.Sprintf("invalid OnDelete action: %s", c.onDelete)
	}
	return true, ""
}

// ValidateTagメソッドは、制約の妥当性を検証し、エラーメッセージを返します。
func (c *Constraint) ValidateTag() error {
	if valid, errMessage := c.IsValid(); !valid {
		return fmt.Errorf("invalid constraint: %s", errMessage)
	}
	return nil
}

func (c *Constraint) GetIsLegitimatelyInitialized() bool {
	return c.isLegitimatelyInitialized
}
