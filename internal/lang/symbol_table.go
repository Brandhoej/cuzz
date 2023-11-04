package lang

type Symbol int

type SymbolTable struct {
	identifiers map[string]Symbol
	symbols     map[Symbol]string
}

func NewSymbolTable() SymbolTable {
	return SymbolTable{
		identifiers: make(map[string]Symbol),
		symbols:     make(map[Symbol]string),
	}
}

func (table *SymbolTable) Length() int {
	return len(table.identifiers)
}

func (table *SymbolTable) Symbols() []Symbol {
	index := 0
	symbols := make([]Symbol, len(table.symbols))
	for symbol := range table.symbols {
		symbols[index] = symbol
		index += 1
	}
	return symbols
}

func (table *SymbolTable) Store(identifier string) Symbol {
	// If the value already exists then we do nothing and return the existing symbol.
	// This gurantees that the returned symbol is always pointing at the same identifier.
	// If we did not have this rule and stored the same identifier twice then could result
	// in cases where lookup for the symbol returned a different identifier.
	if value, exists := table.identifiers[identifier]; exists {
		return Symbol(value)
	}

	// We start at one so the zero'th value can be reserved as a replacement for nil.
	// That way we can reduce the size of our edges and transitions and not make omittable outputs a pointer.
	symbol := Symbol(len(table.identifiers) + 1)
	table.identifiers[identifier] = symbol
	table.symbols[symbol] = identifier
	return Symbol(symbol)
}

func (table *SymbolTable) Lookup(symbol Symbol) (string, bool) {
	value, exists := table.symbols[symbol]
	return value, exists
}
