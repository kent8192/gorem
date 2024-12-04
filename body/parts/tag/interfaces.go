package tag

// TagAttributes は、タグの初期化、バリデーション、およびレンダリングを管理するためのインターフェース
type TagAttributes[T interface {
	TagInitializer[T, V]
	TagValidator
	TagRenderer
}, V any] interface {
	TagInitializer[T, V]
	TagValidator
	TagRenderer
}

// TagInitializer はジェネリック型 T を初期化するためのインターフェース
type TagInitializer[T interface {
	TagValidator
	TagRenderer
}, V any] interface {
	InitTag(value V) T
}

// TagValidator はタグの値をバリデーションするインターフェース
type TagValidator interface {
	// タグの値が有効かを判定する関数
	IsValid() (bool, string)
	// IsValid を用いてそのタグが有効かを判定する関数
	ValidateTag() error
}

// TagRenderer はタグの Gorm のタグ仕様を定義するインターフェース
type TagRenderer interface {
	// タグのインスタンスを Gorm のタグに変換する
	RenderTag() string
}
