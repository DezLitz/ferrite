package ferrite

import (
	"github.com/dogmatiq/ferrite/variable"
)

// Deprecated is the application-facing interface for a value that is sourced
// from deprecated environment variables.
//
// It is obtained by calling Deprecated() on a variable builder.
type Deprecated[T any] interface {
	// DeprecatedValue returns the parsed and validated value of the environment
	// variable, if it is defined.
	//
	// If the environment variable is not defined (and there is no default
	// value), ok is false; otherwise, ok is true and v is the value.
	//
	// It panics if the environment variable is defined but invalid.
	DeprecatedValue() (T, bool)
}

// DeprecatedOption is an option that can be applied to a deprecated variable.
type DeprecatedOption interface {
	variable.RegisterOption
}

// deprecated is a convenience function that registers and returns a
// deprecated[T] that maps one-to-one with an environment variable of the same
// type.
func deprecated[T any, S variable.TypedSchema[T]](
	schema S,
	spec *variable.TypedSpecBuilder[T],
	options []DeprecatedOption,
) Deprecated[T] {
	spec.MarkDeprecated()

	v := variable.Register(
		spec.Done(schema),
		options...,
	)

	// interface is currently empty so we don't need an implementation
	return deprecatedFunc[T]{
		func() (T, bool, error) {
			return v.NativeValue()
		},
	}
}

// deprecatedFunc is an implementation of Deprecated[T] that obtains the value
// from an arbitrary function.
type deprecatedFunc[T any] struct {
	fn func() (T, bool, error)
}

func (d deprecatedFunc[T]) DeprecatedValue() (T, bool) {
	n, ok, err := d.fn()
	if err != nil {
		panic(err.Error())
	}
	return n, ok
}
