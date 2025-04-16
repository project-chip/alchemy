package render

type Option func(r *Renderer)

type RenderOptions struct {
	WrapLength int `name:"wrap" default:"0" help:"the maximum length of a line" group:"Output:"`
}

func (ro RenderOptions) ToOptions() []Option {
	return []Option{Wrap(ro.WrapLength)}
}
