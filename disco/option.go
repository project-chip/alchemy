package disco

type Option func(b *Ball)

type options struct {
	linkAttributes bool
}

func LinkAttributes(link bool) Option {
	return func(b *Ball) {
		b.options.linkAttributes = link
	}
}

func AddMissingColumns(add bool) Option {
	return func(b *Ball) {
	}
}
