package lang

import "testing"

func TestNamespaceContextSetWithoutFailure(t *testing.T) {
	val := NewESValue(10)
	context := NewNamespaceContext()
	if err := context.Set("a", val); err != nil {
		t.Errorf("Failure setting variable in NamespaceContext: %s", err)
	}
}

func TestFetchingNonExistentFromNamespaceContextReturnsNil(t *testing.T) {
	c := NewNamespaceContext()
	val, err := c.Get("a")
	if !val.IsNil() || err != nil {
		t.Errorf("The value returned was not a nil ESValue: %s", val.DebugString())
	}
}

func TestSettingAndFetchingDataFromNamespaceContext(t *testing.T) {
	val := NewESValue(10)
	context := NewNamespaceContext()
	context.Set("a", val)
	newVal, _ := context.Get("a")
	i, _ := newVal.Int()
	if newVal.IsNil() || i != 10 {
		t.Errorf("Failed to fetch the correct data from the namespace context: Expected %s got %s", val.DebugString(), newVal.DebugString())
	}
}
