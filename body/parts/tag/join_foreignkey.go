package tag

import "fmt"

// JoinForeignKey構造体
type JoinForeignKey struct {
	isLegitimatelyInitialized bool
	fieldName                 string
}

// InitTagメソッドで結合テーブルの外部キーのフィールド名を初期化
func (j *JoinForeignKey) InitTag(value string) *JoinForeignKey {
	j.isLegitimatelyInitialized = true
	j.fieldName = value
	return j
}

// RenderTagメソッドでタグ文字列を生成
func (j *JoinForeignKey) RenderTag() string {
	if !j.isLegitimatelyInitialized {
		return ""
	}
	return fmt.Sprintf("joinForeignKey:%s", j.fieldName)
}

// IsValidメソッドで結合テーブルの外部キーの妥当性をチェック
func (j *JoinForeignKey) IsValid() (bool, string) {
	if !j.isLegitimatelyInitialized {
		return false, "joinForeignKey is not initialized legitimately"
	}
	if j.fieldName == "" {
		return false, "joinForeignKey cannot be empty"
	}
	return true, ""
}

// ValidateTagメソッドでエラーメッセージを返す
func (j *JoinForeignKey) ValidateTag() error {
	if valid, errMessage := j.IsValid(); !valid {
		return fmt.Errorf("invalid joinForeignKey: %s", errMessage)
	}
	return nil
}

func (j *JoinForeignKey) GetIsLegitimatelyInitialized() bool {
	return j.isLegitimatelyInitialized
}
