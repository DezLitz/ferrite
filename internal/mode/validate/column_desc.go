package validate

import "github.com/dogmatiq/ferrite/internal/variable"

// description renders a column containing the variable's human-readable
// description.
func description(v variable.Any) string {
	return v.Spec().Description()
}
