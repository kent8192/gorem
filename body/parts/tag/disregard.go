package tag

import (
	"fmt"
)

// Level は、DisregardField.level の列挙型です。
type Level int

const (
	LevelAll Level = iota
	LevelMigration
	LevelEmpty
)

// String は、Level を文字列として返します。
func (l Level) String() string {
	return [...]string{"all", "migration", ""}[l]
}

// IsValidLevel は、与えられたレベル文字列が有効かどうかを確認します。
func IsValidLevel(level string) bool {
	validLevels := []string{
		LevelAll.String(),
		LevelMigration.String(),
		LevelEmpty.String(),
	}
	for _, valid := range validLevels {
		if level == valid {
			return true
		}
	}
	return false
}

// DisregardField 構造体は、GORMでフィールドを無視するためのタグ(`-`)を管理します。
type DisregardField struct {
	level                     string
	isLegitimatelyInitialized bool
}

// InitTag メソッドは、DisregardField を初期化します。
func (d *DisregardField) InitTag(value string) *DisregardField {
	d.isLegitimatelyInitialized = true
	d.level = value
	if err := d.ValidateTag(); err != nil {
		panic(err)
	}
	return d
}

// RenderTag メソッドは、GORMでフィールドを無視するためのタグ文字列を生成します。
func (d *DisregardField) RenderTag() string {
	if !d.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("-:%s", d.level)
}

// IsValid メソッドは、DisregardField の妥当性をチェックします。
func (d *DisregardField) IsValid() (bool, string) {
	// level が有効な値かチェック
	if !IsValidLevel(d.level) {
		return false, fmt.Sprintf("DisregardField.level '%s' is invalid", d.level)
	}

	// 正しく初期化されているかチェック
	if !d.isLegitimatelyInitialized {
		return false, "DisregardField is not initialized legitimately"
	}

	return true, ""
}

// ValidateTag メソッドは、タグの妥当性を検証し、エラーメッセージを返します。
func (d *DisregardField) ValidateTag() error {
	if valid, errMessage := d.IsValid(); !valid {
		return fmt.Errorf("invalid DisregardField: %s", errMessage)
	}
	return nil
}

func (d *DisregardField) GetIsLegitimatelyInitialized() bool {
	return d.isLegitimatelyInitialized
}
