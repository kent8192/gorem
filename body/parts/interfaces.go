package parts

type TagValidator[T any] interface {
	Validate(value T) error
}

type TagRenderer interface {
	RenderTag() string
}
