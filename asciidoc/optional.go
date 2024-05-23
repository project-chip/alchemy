package asciidoc

type Optional[T comparable] struct {
	Value T
	IsSet bool
}

func One[T comparable](val T) Optional[T] {
	return Optional[T]{Value: val, IsSet: true}
}

func Default[T comparable](val T) Optional[T] {
	return Optional[T]{Value: val, IsSet: false}
}

func Maybe[T comparable](val any, defaulValue T) Optional[T] {
	if v, ok := val.(T); ok {
		return One(v)
	}
	return Optional[T]{Value: defaulValue, IsSet: false}
}

func (o Optional[T]) Equals(o2 Optional[T]) bool {
	return o.Value == o2.Value && o.IsSet == o2.IsSet
}
