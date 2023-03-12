package ferrite

import (
	"github.com/dogmatiq/ferrite/variable"
)

// String configures an environment variable as a string.
//
// name is the name of the environment variable to read. desc is a
// human-readable description of the environment variable.
func String(name, desc string) *StringBuilder[string] {
	return StringAs[string](name, desc)
}

// StringAs configures an environment variable as a string using a user-defined
// type.
//
// name is the name of the environment variable to read. desc is a
// human-readable description of the environment variable.
func StringAs[T ~string](name, desc string) *StringBuilder[T] {
	b := &StringBuilder[T]{}

	b.spec.Name(name)
	b.spec.Description(desc)

	return b
}

// StringBuilder builds a specification for a string variable.
type StringBuilder[T ~string] struct {
	schema variable.TypedString[T]
	spec   variable.TypedSpecBuilder[T]
}

var _ isBuilderOf[string, *StringBuilder[string]]

// WithDefault sets a default value of the variable.
//
// It is used when the environment variable is undefined or empty.
func (b *StringBuilder[T]) WithDefault(v T) *StringBuilder[T] {
	b.spec.Default(v)
	return b
}

// WithConstraint adds a constraint to the variable.
//
// fn is called with the environment variable value after it is parsed. If fn
// returns false the value is considered invalid.
func (b *StringBuilder[T]) WithConstraint(
	desc string,
	fn func(T) bool,
) *StringBuilder[T] {
	b.spec.UserConstraint(desc, fn)
	return b
}

// WithSensitiveContent marks the variable as containing sensitive content.
//
// Values of sensitive variables are not printed to the console or included in
// generated documentation.
func (b *StringBuilder[T]) WithSensitiveContent() *StringBuilder[T] {
	b.spec.MarkSensitive()
	return b
}

// SeeAlso creates a relationship between this variable and those used by i.
func (b *StringBuilder[T]) SeeAlso(i Input, options ...SeeAlsoOption) *StringBuilder[T] {
	seeAlsoInput(&b.spec, i, options...)
	return b
}

// Required completes the build process and registers a required variable with
// Ferrite's validation system.
func (b *StringBuilder[T]) Required(options ...RequiredOption) Required[T] {
	return required(b.schema, &b.spec, options...)
}

// Optional completes the build process and registers an optional variable with
// Ferrite's validation system.
func (b *StringBuilder[T]) Optional(options ...OptionalOption) Optional[T] {
	return optional(b.schema, &b.spec, options...)
}

// Deprecated completes the build process and registers a deprecated variable
// with Ferrite's validation system.
func (b *StringBuilder[T]) Deprecated(options ...DeprecatedOption) Deprecated[T] {
	return deprecated(b.schema, &b.spec, options...)
}
