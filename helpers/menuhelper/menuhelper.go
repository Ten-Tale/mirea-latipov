package menuhelper

import (
	"github.com/dixonwille/wmenu"
)

type Option struct {
	Title   string
	Handler func(wmenu.Opt) error
}

func ApplyOptionList(m *wmenu.Menu, ol []Option) {
	for _, o := range ol {
		m.Option(o.Title, nil, false, o.Handler)
	}
}
