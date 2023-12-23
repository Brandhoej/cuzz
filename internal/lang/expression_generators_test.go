package lang

import (
	"math/rand"
	"os"
	"testing"
)

func Test_EqualityGenericFunction(t *testing.T) {
	symbols := NewSymbolTable()

	anySymbol := symbols.Store("any")
	anyType := Type{
		identifier: anySymbol,
	}

	booleanSymbol := symbols.Store("boolean")
	booleanType := Type{
		identifier: booleanSymbol,
	}

	int32Symbol := symbols.Store("int32")
	int32Type := Type{
		identifier: int32Symbol,
	}

	int32Fn := Function{
		identifier: symbols.Store("random int32"),
		generics:   []Generic{},
		parameters: []Symbol{},
		returnType: int32Symbol,
	}

	lessThanFn := Function{
		identifier: symbols.Store("less than"),
		generics:   []Generic{},
		parameters: []Symbol{
			int32Symbol,
			int32Symbol,
		},
		returnType: booleanSymbol,
	}

	trueFn := Function{
		identifier: symbols.Store("true"),
		generics:   []Generic{},
		parameters: []Symbol{},
		returnType: booleanSymbol,
	}

	falseFn := Function{
		identifier: symbols.Store("false"),
		generics:   []Generic{},
		parameters: []Symbol{},
		returnType: booleanSymbol,
	}

	generic := Generic{
		identifier: symbols.Store("T"),
		typeSet: Types{
			mapping: map[Symbol]Type{
				symbols.Store("T"): {anySymbol},
			},
		},
	}

	equalityFn := Function{
		identifier: symbols.Store("equality"),
		generics:   []Generic{generic},
		parameters: []Symbol{
			generic.identifier,
			generic.identifier,
		},
		returnType: booleanSymbol,
	}

	functions := FunctionSet{
		set: []Function{
			trueFn,
			int32Fn,
			equalityFn,
			lessThanFn,
			falseFn,
		},
	}

	typeTree := TypeTree{
		relations: map[Symbol][]Symbol{
			booleanSymbol: {anySymbol},
			int32Symbol:   {anySymbol},
		},
	}

	generator := RndExpressionGenerator{
		functions: functions,
		types: Types{
			mapping: map[Symbol]Type{
				anySymbol:     anyType,
				booleanSymbol: booleanType,
				int32Symbol:   int32Type,
			},
		},
		typeTree: typeTree,
		factories: map[Symbol]func(parameters ...Expression) Expression{
			equalityFn.identifier: func(parameters ...Expression) Expression {
				return BinaryExpression{
					Lhs:      parameters[0],
					Operator: Equality,
					Rhs:      parameters[1],
				}
			},
			trueFn.identifier: func(parameters ...Expression) Expression {
				return ConstantBoolean{
					Value: true,
				}
			},
			falseFn.identifier: func(parameters ...Expression) Expression {
				return ConstantBoolean{
					Value: false,
				}
			},
			int32Fn.identifier: func(parameters ...Expression) Expression {
				return ConstantInt32{
					Value: int32(rand.Intn(1000)),
				}
			},
			lessThanFn.identifier: func(parameters ...Expression) Expression {
				if len(parameters) != 2 {
					panic("Equality expected two parameters")
				}

				return BinaryExpression{
					Lhs:      parameters[0],
					Operator: LessThan,
					Rhs:      parameters[1],
				}
			},
		},
	}

	generation, _ := generator.Generate(booleanType)

	emitter := ExpressionEmitter{
		writer: os.Stdout,
	}

	generation.Accept(&emitter)
}
