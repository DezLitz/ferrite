package markdownmode

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dogmatiq/ferrite/variable"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// Run generates environment variable usage instructions in markdown format.
func Run(reg *variable.Registry) string {
	w := &strings.Builder{}

	references := map[string]string{
		"dogmatiq/ferrite": "https://github.com/dogmatiq/ferrite",
	}

	fmt.Fprintln(w, "# Environment Variables")
	fmt.Fprintln(w, "")
	fmt.Fprintf(w, "This document describes the environment variables used by `%s`.\n", filepath.Base(os.Args[0]))

	variables := reg.Variables()

	if len(variables) == 0 {
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "**There do not appear to be any environment variables.**")
	}

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "The application may consume other undocumented environment variables; this")
	fmt.Fprintln(w, "document only shows those variables defined using [dogmatiq/ferrite].")

	if len(variables) != 0 {
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "## Index")
		fmt.Fprintln(w, "")

		// 		for _, r := range resolvers {
		// 			renderMarkdownSpecLink(w, r.Spec())
		// 		}

		// 		fmt.Fprintln(w, "")
		// 		fmt.Fprintln(w, "## Specification")

		// 		for _, r := range resolvers {
		// 			fmt.Fprintln(w, "")
		// 			renderMarkdownSpec(w, r.Spec())
		// 		}
	}

	fmt.Fprintln(w, "")
	renderReferences(w, references)

	return w.String()
}

// func renderMarkdownSpecLink(w *strings.Builder, s spec.Spec) {
// 	fmt.Fprintf(w, "- [`%s`](#%s) — %s\n", s.Name, s.Name, s.Description)
// }

// func renderMarkdownSpec(w *strings.Builder, s spec.Spec) {
// 	fmt.Fprintf(w, "### `%s`\n", s.Name)
// 	fmt.Fprintln(w, "")
// 	fmt.Fprintf(w, "> %s\n", s.Description)
// 	fmt.Fprintln(w, "")

// 	if s.IsOptional {

// 	} else if s.HasDefault {
// 		fmt.Fprintln(w, "This variable is **required**, although a default is provided.")
// 	}

// 	s.Schema.AcceptVisitor(&markdownSchemaRenderer{w})
// }

func renderReferences(w *strings.Builder, refs map[string]string) {
	fmt.Fprintln(w, "<!-- references -->")
	fmt.Fprintln(w, "")

	keys := maps.Keys(refs)
	slices.SortFunc(
		keys,
		func(a, b string) bool {
			return strings.Trim(a, "`") < strings.Trim(b, "`")
		},
	)

	for _, k := range keys {
		fmt.Fprintf(w, "[%s]: %s\n", k, refs[k])
	}
}

// type markdownSchemaRenderer struct {
// 	Output *strings.Builder
// }

// func (r *markdownSchemaRenderer) VisitOneOf(s spec.OneOf) {
// 	for i, c := range s {
// 		if i > 0 {
// 			r.Output.WriteString(" | ")
// 		}

// 		c.AcceptVisitor(r)
// 	}
// }

// func (r *markdownSchemaRenderer) VisitLiteral(s spec.Literal) {
// 	fmt.Fprintf(r.Output, "%s", s)
// }

// func (r *markdownSchemaRenderer) VisitType(s spec.Type) {
// 	fmt.Fprintf(r.Output, "<%s>", s.Type)
// }

//	func (r *markdownSchemaRenderer) VisitRange(s spec.Range) {
//		if s.Min != "" && s.Max != "" {
//			fmt.Fprintf(r.Output, "%s .. %s", s.Min, s.Max)
//		} else if s.Max != "" {
//			fmt.Fprintf(r.Output, "... %s", s.Max)
//		} else {
//			fmt.Fprintf(r.Output, "%s ...", s.Min)
//		}

// }
