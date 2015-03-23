package lang

import (
	"reflect"
	"testing"
)

func TestNewESValWithValidValues(t *testing.T) {
	val, err := NewESValueWithType(ESTypeString, "hello")
	s, _ := val.String()
	if err != nil && s == "hello" {
		t.Error(err.Error())
	}
}

func TestNewESValWithInvalidValues(t *testing.T) {
	val, err := NewESValueWithType(ESTypeString, 10)
	if err == nil {
		t.Errorf("No error was returned, an invalid ESValue was created: %s", val.DebugString())
	}
}

func TestAnyUintValueCanBeUsed(t *testing.T) {
	var u uint8 = 10
	val := NewESValue(u)
	i, _ := val.Int()
	iVal := reflect.ValueOf(i)
	if val.Type() != ESTypeInteger || i != int64(10) || iVal.Kind() != reflect.Int64 {
		t.Error("Unable to coerce uint values to int64 type for storage by ESValue!")
	}
}

func TestAnyIntValueCanBeUsed(t *testing.T) {
	var u int8 = 10
	val := NewESValue(u)
	i, _ := val.Int()
	iVal := reflect.ValueOf(i)
	if val.Type() != ESTypeInteger || i != int64(10) || iVal.Kind() != reflect.Int64 {
		t.Error("Unable to coerce int values to int64 type for storage by ESValue!")
	}
}

func TestSettingInvalidValues(t *testing.T) {
	val := NewESValue(10)
	err := val.Set("hello!")
	if err == nil {
		t.Error("Successfully set an invalid value for an ESValue type of Integer")
	}
}

func TestTypeCastingToInvalidTypes(t *testing.T) {
	val := NewESValue(10.4)
	if _, err := val.Float(); err != nil {
		t.Errorf("Unable to cast ESValue to proper type: %s", val.DebugString())
	}
	if _, err := val.Int(); err == nil {
		t.Errorf("Successfully cast ESValue to invalid type: %s", err)
	}
}

func TestESNilValueIsNil(t *testing.T) {
	val := ESNilValue()
	if !val.IsNil() {
		t.Error("ESNilValue does not return a nil ESValue!")
	}
}

func TestESValuesAreNotNil(t *testing.T) {
	val := NewESValue("Hello")
	if val.IsNil() {
		t.Errorf("ESValue (%s) is reporting as nil!", val.DebugString())
	}
}
