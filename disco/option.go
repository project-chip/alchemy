package disco

type Option func(b *Ball)

func LinkAttributes(b *Ball) {
	b.ShouldLinkAttributes = true
}
