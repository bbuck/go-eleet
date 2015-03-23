package lang

import (
	"fmt"
	"reflect"
)

type ESType int

// Eleetscript types
const (
	ESTypeString ESType = 1 << iota
	ESTypeInteger
	ESTypeFloat
	ESTypeBool
	ESTypeObject
	ESTypeNil
)

// Make ESType conform to the Stringer interface.
func (e ESType) String() string {
	switch e {
	case ESTypeString:
		return "ESString"
	case ESTypeInteger:
		return "ESInteger"
	case ESTypeFloat:
		return "ESFloat"
	case ESTypeBool:
		return "ESBoolean"
	// TODO: Add ESTypeObject case
	case ESTypeNil:
		return "ESNil"
	default:
		return "ESUnknownType"
	}
}

// Store an Eleetscript value for passing around Go Code. Use of various methods
// allow for accessing the value as a Go Value.
type ESValue struct {
	t     ESType
	value interface{}
}

// Create a new ESValue from the value given. This method will assign the
// appropreiate type to the ESValue and then assign the value (ignoring errors
// from assignment, as there should be none). If invalid types are given then
// an ESTypeNil value will be returned.
func NewESValue(val interface{}) *ESValue {
	esVal := &ESValue{value: val}

	// Guess the ESType of the Go value given.
	switch val.(type) {
	case string:
		esVal.t = ESTypeString
	case uint, uint8, uint16, uint32, uint64, int, int8, int16, int32, int64:
		esVal.t = ESTypeInteger
	case float32, float64:
		esVal.t = ESTypeFloat
	case bool:
		esVal.t = ESTypeBool
	// TODO:: Add case for ESTypeObject
	default:
		esVal.t = ESTypeNil
		esVal.value = nil
	}

	if err := esVal.Set(val); err != nil {
		esVal.t = ESTypeNil
		esVal.value = nil
	}

	return esVal
}

// Create a new ESValue with an explicit (no type guessing). This will return an
// error if invalid value is given that doesn't match the type assigned to it.
// This method shouldn't be preferred, but can be used.
func NewESValueWithType(t ESType, val interface{}) (*ESValue, error) {
	esVal := &ESValue{t: t}

	if err := esVal.Set(val); err != nil {
		return nil, err
	} else {
		return esVal, nil
	}
}

// Return an ESValue instance that represents Nil. This isn't a constant value,
// nor should it necessarily be.
func ESNilValue() *ESValue {
	return &ESValue{t: ESTypeNil, value: nil}
}

// Set the value of this ESType, this will return an error if the value given
// does not conform to the type of the ESValue.
func (e *ESValue) Set(val interface{}) (err error) {
	err = ESInvalidValueError{e, val}
	refVal := reflect.ValueOf(val)

	// Numeric switch, capture numeric types in such a way as to prevent
	// duplicating logic.
	switch refVal.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if e.t != ESTypeInteger {
			return
		} else {
			e.value = int64(refVal.Uint())
			return nil // don't finish with the type switch
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if e.t != ESTypeInteger {
			return
		} else {
			e.value = refVal.Int()
			return nil // don't finish with the type switch
		}
	case reflect.Float32, reflect.Float64:
		if e.t != ESTypeFloat {
			return
		} else {
			e.value = refVal.Float()
			return nil // don't finish with the type switch
		}
	}

	// Standard tpye switch, capture the different types other than numeric.
	switch t := val.(type) {
	case nil:
		if e.t != ESTypeNil {
			return
		} else {
			e.value = nil
		}
	case string:
		if e.t != ESTypeString {
			return
		} else {
			e.value = val
		}
	case bool:
		if e.t != ESTypeBool {
			return
		} else {
			e.value = t
		}
	// TODO: Add in case for ESTypeObject
	default:
		return
	}

	return nil
}

// Return a string with useful debug information (this method is defined due to
// the fact that "String()" was used for something else.
func (e *ESValue) DebugString() string {
	return fmt.Sprintf("ESValue{t: %s, value: %#v %T}", e.t, e.value, e.value)
}

// Return the ESType of the value, this is useful when building a switch or
// type checking internally before using an invalid type conversion method.
func (e *ESValue) Type() ESType {
	return e.t
}

// Return the value of the ESValue as a string. If the ESValue doesn't
// represent, or hold a string then an error will be returned.
func (e *ESValue) String() (string, error) {
	if e.t == ESTypeString {
		s, _ := e.value.(string)

		return s, nil
	} else {
		return "", ESTypeError{e, "string"}
	}
}

// Return the value of the ESValue as an int64. If the ESValue doesn't represent
// or hold an int then an error will be returned.
func (e *ESValue) Int() (int64, error) {
	if e.t == ESTypeInteger {
		i, _ := e.value.(int64)

		return i, nil
	} else {
		return 0, ESTypeError{e, "int64"}
	}
}

// Return the value of the ESValue as a float64. If the ESValue doesn't represent
// or hold a float then an error will be returned.
func (e *ESValue) Float() (float64, error) {
	if e.t == ESTypeFloat {
		f, _ := e.value.(float64)

		return f, nil
	} else {
		return 0.0, ESTypeError{e, "float64"}
	}
}

// Return the value of the ESValue as a bool, If the ESValue doesn't represent
// or hold a bool then an error will be returned.
func (e *ESValue) Bool() (bool, error) {
	if e.t == ESTypeBool {
		b, _ := e.value.(bool)

		return b, nil
	} else {
		return false, ESTypeError{e, "bool"}
	}
}

// Return true if this ESValue points to nil.
func (e *ESValue) IsNil() bool {
	return e.t == ESTypeNil
}
