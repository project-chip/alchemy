package disco

type Option func(b *Ball)

func LinkAttributes(link bool) Option {
	return func(b *Ball) {
		b.ShouldLinkAttributes = link
	}
}
