package lang

import "fmt"

// Used to signify a simple error when a type that doesn't match the given
// ESValue's type was requested.
type ESTypeError struct {
	esVal     *ESValue
	requested string
}

// Return a message explaining the type of the ESValue and what type was requested
// from it.
func (e ESTypeError) Error() string {
	return fmt.Sprintf("The ESValue is of type %s, requested %s", e.esVal.t, e.requested)
}

// Represents setting an ESValue with a value that doesn't match the type of
// the ESValue object.
type ESInvalidValueError struct {
	esVal *ESValue
	given interface{}
}

// Return a message explaining the type of the value and the type given to it.
func (e ESInvalidValueError) Error() string {
	return fmt.Sprintf("The ESValue is of type %s, the value given was %T", e.esVal, e.given)
}
