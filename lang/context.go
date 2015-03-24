package lang

type Context interface {
	Get(string) *ESValue
	Set(string, *ESValue) error
	Parent() Context
}

type NamespaceContext struct {
	globals, locals map[string]*ESValue
	namespaces      map[string]*NamespaceContext
}

func NewNamespaceContext() *NamespaceContext {
	return &NamespaceContext{
		globals:    make(map[string]*ESValue),
		locals:     make(map[string]*ESValue),
		namespaces: make(map[string]*NamespaceContext)}
}

func (n *NamespaceContext) Get(name string) (*ESValue, error) {
	if len(name) == 0 {
		return nil, InvalidIdentifierError{name, "the identifier must have a length"}
	}

	switch name[0] {
	case '$':
		if val, ok := n.globals[name]; ok {
			return val, nil
		} else {
			return ESNilValue(), nil
		}
	default:
		if val, ok := n.locals[name]; ok {
			return val, nil
		} else {
			return ESNilValue(), nil
		}
	}
}

func (n *NamespaceContext) Set(name string, val *ESValue) error {
	if len(name) == 0 {
		return InvalidIdentifierError{name, "the identifier must have a length"}
	}

	switch name[0] {
	default:
		n.locals[name] = val
	case '$':
		n.globals[name] = val
	}

	return nil
}
