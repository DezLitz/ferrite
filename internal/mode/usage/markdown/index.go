package markdown

import (
	"github.com/dogmatiq/ferrite/variable"
)

func (r *renderer) renderIndex() {
	r.line("## Index")
	r.gap()

	for _, v := range r.Specs {
		r.renderIndexItem(v)
	}
}

func (r *renderer) renderIndexItem(s variable.Spec) {
	r.line(
		"- %s — %s",
		r.linkToSpec(s),
		s.Description(),
	)
}
