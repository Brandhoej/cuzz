package lang

/*type Scope struct {
	symbols   SymbolTable
	types     Types
	functions FunctionSet
}

func (scope *Scope) AddType(identifier string) (Symbol, Type) {
	symbol := scope.symbols.Store(identifier)
	t := Type{
		identifier: symbol,
	}
	scope.types.mapping[symbol] = t
	return symbol, t
}

func (scope *Scope) AddFunction(function Function) {
	scope.functions.set = append(scope.functions.set, function)
}

type Language struct {
	global Scope
}

func CreateGoLang() Language {
	global := Scope{
		symbols: NewSymbolTable(),
		types: Types{
			mapping: make(map[Symbol]Type),
		},
		functions: FunctionSet{
			set: []Function{},
		},
	}

	// Types:
	anySymbol, _ := global.AddType("any")
	comparableSymbol, _ := global.AddType("comparable")

	booleanSymbol, _ := global.AddType("boolean")

	uintSymbol, _ := global.AddType("uint")
	uint8Symbol, _ := global.AddType("uint8")
	uint16Symbol, _ := global.AddType("uint16")
	uint32Symbol, _ := global.AddType("uint32")
	uint64Symbol, _ := global.AddType("uint64")

	intSymbol, _ := global.AddType("int")
	int8Symbol, _ := global.AddType("int8")
	int16Symbol, _ := global.AddType("int16")
	int32Symbol, _ := global.AddType("int32")
	int64Symbol, _ := global.AddType("int64")

	float32Symbol, _ := global.AddType("float32")
	float64Symbol, _ := global.AddType("float64")

	complex64Symbol, _ := global.AddType("complex64")
	complex128Symbol, _ := global.AddType("complex128")

	byteSymbol, _ := global.AddType("byte")
	runeSymbol, _ := global.AddType("rune")

	typeTree := TypeTree{
		relations: map[Symbol][]Symbol{
			booleanSymbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			uintSymbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			uint8Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			uint16Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			uint32Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			uint64Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			intSymbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			int8Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			int16Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			int32Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			int64Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			float32Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			float64Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			complex64Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			complex128Symbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			byteSymbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
			runeSymbol: []Symbol{
				anySymbol,
				comparableSymbol,
			},
		},
	}

	// Literals:
	trueFn := Function{
		identifier: global.symbols.Store("true"),
		generics:   []Generic{},
		parameters: []Symbol{},
		returnType: booleanSymbol,
	}
	falseFn := Function{
		identifier: global.symbols.Store("false"),
		generics:   []Generic{},
		parameters: []Symbol{},
		returnType: booleanSymbol,
	}

	// Operators:
	lessThanFn := Function{
		identifier: global.symbols.Store("less than"),
		generics:   []Generic{},
		parameters: []Symbol{
			int32Symbol,
			int32Symbol,
		},
	}

	return Language{
		global,
	}
}*/
