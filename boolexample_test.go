package ferrite_test

import (
	"fmt"
	"os"

	"github.com/dogmatiq/ferrite"
)

func ExampleBool() {
	ferrite.DefaultRegistry.Reset()
	os.Setenv("FERRITE_BOOL", "true")
	defer os.Unsetenv("FERRITE_BOOL")

	value := ferrite.Bool(
		"FERRITE_BOOL",
		"example boolean variable",
	)

	ferrite.ValidateEnvironment()

	fmt.Println("value is", value.Value())

	// Output:
	// value is true
}

func ExampleBool_default() {
	ferrite.DefaultRegistry.Reset()

	value := ferrite.Bool(
		"FERRITE_BOOL",
		"example boolean variable",
	).
		Default(true)

	ferrite.ValidateEnvironment()

	fmt.Println("value is", value.Value())

	// Output:
	// value is true
}

func ExampleBool_customLiterals() {
	ferrite.DefaultRegistry.Reset()
	os.Setenv("FERRITE_BOOL", "yes")
	defer os.Unsetenv("FERRITE_BOOL")

	value := ferrite.Bool(
		"FERRITE_BOOL",
		"example boolean variable",
	).
		Literals("yes", "no")

	ferrite.ValidateEnvironment()

	fmt.Println("value is", value.Value())

	// Output:
	// value is true
}
