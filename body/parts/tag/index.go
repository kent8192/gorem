package tag

import "fmt"

// Index 構造体
type Index struct {
	isLegitimatelyInitialized bool
	name                      string
	sort                      string
}

// InitTag メソッドでインデックス名を初期化
func (i *Index) InitTag(name string, sort string) *Index {
	i.isLegitimatelyInitialized = true
	i.name = name
	i.sort = sort
	if err := i.ValidateTag(); err != nil {
		panic(err)
	}
	return i
}

// RenderTag メソッドでタグ文字列を生成
func (i *Index) RenderTag() string {
	if !i.isLegitimatelyInitialized {
		return ""
	}
	if i.sort != "" {
		return fmt.Sprintf("index:%s,sort:%s", i.name, i.sort)
	}
	return fmt.Sprintf("index:%s", i.name)
}

// IsValid メソッドでインデックスの妥当性をチェック
func (i *Index) IsValid() (bool, string) {
	if !i.isLegitimatelyInitialized {
		return false, "index is not initialized legitimately"
	}
	if i.name == "" {
		return false, "index name cannot be empty"
	}
	if i.sort != "" && i.sort != "asc" && i.sort != "desc" {
		return false, "sort must be 'asc' or 'desc' if specified"
	}
	return true, ""
}

// ValidateTag メソッドでエラーメッセージを返す
func (i *Index) ValidateTag() error {
	if valid, errMessage := i.IsValid(); !valid {
		return fmt.Errorf("invalid index: %s", errMessage)
	}
	return nil
}

func (i *Index) GetIsLegitimatelyInitialized() bool {
	return i.isLegitimatelyInitialized
}
