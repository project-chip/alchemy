package pipeline

type Data[T any] struct {
	Path    string
	Content T
}
