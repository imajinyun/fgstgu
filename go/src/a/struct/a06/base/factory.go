package base

// Class interface.
type Class interface {
	Do()
}

var factoryByName = make(map[string]func() Class)

// Register func.
func Register(name string, factory func() Class) {
	factoryByName[name] = factory
}

// Create func.
func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	}
	panic("Name not found")
}
