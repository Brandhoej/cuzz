package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	symbols := NewSymbolTable()
	// Types.
	anyTypeSymbol := symbols.Store("any")
	int32TypeSymbol := symbols.Store("int32")
	stringTypeSymbol := symbols.Store("string")
	booleanTypeSymbol := symbols.Store("boolean")
	types := Types{
		Relations: map[Symbol][]Symbol{
			int32TypeSymbol:   {anyTypeSymbol},
			stringTypeSymbol:  {anyTypeSymbol},
			booleanTypeSymbol: {anyTypeSymbol},
		},
	}

	// Unary operators.
	numericNegationFunctionSymbol := symbols.Store("(-)")
	logicalNegationFunctionSymbol := symbols.Store("!")
	bitwiseComplementFunctionSymbol := symbols.Store("(^)")
	// Binary operators.
	AdditionFunctionSymbol := symbols.Store("+")
	SubtractionFunctionSymbol := symbols.Store("-")
	MultiplicationFunctionSymbol := symbols.Store("*")
	DivisionFunctionSymbol := symbols.Store("/")
	RemainderFunctionSymbol := symbols.Store("%")
	ConcatenationFunctionSymbol := symbols.Store("++")
	EqualityFunctionSymbol := symbols.Store("==")
	InequalityFunctionSymbol := symbols.Store("!=")
	LessThanFunctionSymbol := symbols.Store("<")
	LessThanOrEqualFunctionSymbol := symbols.Store("<=")
	GreaterThanFunctionSymbol := symbols.Store(">")
	GreaterThanOrEqualFunctionSymbol := symbols.Store(">=")
	LogicalDisjunctionFunctionSymbol := symbols.Store("||")
	LogicalConjunctionFunctionSymbol := symbols.Store("&&")
	BitwiseDisjunctionFunctionSymbol := symbols.Store("|")
	BitwiseConjunctionFunctionSymbol := symbols.Store("&")
	BitwizeExclusiveDisjunctionFunctionSymbol := symbols.Store("^")
	BitwiseLeftShiftFunctionSymbol := symbols.Store("<<")
	BitwiseRightShiftFunctionSymbol := symbols.Store(">>")
	BitwiseClearFunctionSymbol := symbols.Store("&^")
	// Random generators.
	randomIntFuncSymbol := symbols.Store("random int")
	randomBoolFuncSymbol := symbols.Store("random bool")
	randomStringFuncSymbol := symbols.Store("random string")

	graph := FunctionSignatureGraph{
		Signatures: []FunctionSignature{
			// Numeric operators.
			CreateUnaryFunctionSignature(numericNegationFunctionSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateUnaryFunctionSignature(logicalNegationFunctionSymbol, booleanTypeSymbol, booleanTypeSymbol),
			CreateUnaryFunctionSignature(bitwiseComplementFunctionSymbol, int32TypeSymbol, int32TypeSymbol),

			CreateBinaryFunctionSignature(AdditionFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateBinaryFunctionSignature(SubtractionFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateBinaryFunctionSignature(MultiplicationFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			// TODO: Only allow non-zero division.
			/*CreateBinaryFunctionSignature(DivisionFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateBinaryFunctionSignature(RemainderFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),*/
			CreateBinaryFunctionSignature(LessThanFunctionSymbol, int32TypeSymbol, int32TypeSymbol, booleanTypeSymbol),
			CreateBinaryFunctionSignature(LessThanOrEqualFunctionSymbol, int32TypeSymbol, int32TypeSymbol, booleanTypeSymbol),
			CreateBinaryFunctionSignature(GreaterThanFunctionSymbol, int32TypeSymbol, int32TypeSymbol, booleanTypeSymbol),
			CreateBinaryFunctionSignature(GreaterThanOrEqualFunctionSymbol, int32TypeSymbol, int32TypeSymbol, booleanTypeSymbol),

			CreateBinaryFunctionSignature(ConcatenationFunctionSymbol, stringTypeSymbol, stringTypeSymbol, stringTypeSymbol),

			// Boolean operators.
			CreateBinaryFunctionSignature(LogicalDisjunctionFunctionSymbol, booleanTypeSymbol, booleanTypeSymbol, booleanTypeSymbol),
			CreateBinaryFunctionSignature(LogicalConjunctionFunctionSymbol, booleanTypeSymbol, booleanTypeSymbol, booleanTypeSymbol),

			// Bitable operators.
			CreateBinaryFunctionSignature(BitwiseDisjunctionFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateBinaryFunctionSignature(BitwiseConjunctionFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateBinaryFunctionSignature(BitwizeExclusiveDisjunctionFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			// TODO: Only allow positive amount of shifts.
			/*CreateBinaryFunctionSignature(BitwiseLeftShiftFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),
			CreateBinaryFunctionSignature(BitwiseRightShiftFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),*/
			CreateBinaryFunctionSignature(BitwiseClearFunctionSymbol, int32TypeSymbol, int32TypeSymbol, int32TypeSymbol),

			// Equality operators.
			CreateBinaryFunctionSignature(EqualityFunctionSymbol, anyTypeSymbol, anyTypeSymbol, booleanTypeSymbol),
			CreateBinaryFunctionSignature(InequalityFunctionSymbol, anyTypeSymbol, anyTypeSymbol, booleanTypeSymbol),

			// Random generators.
			CreateNullaryFunctionSignature(randomIntFuncSymbol, int32TypeSymbol),
			CreateNullaryFunctionSignature(randomBoolFuncSymbol, booleanTypeSymbol),
			CreateNullaryFunctionSignature(randomStringFuncSymbol, stringTypeSymbol),
		},
		Types: types,
	}

	graph.Dot(symbols, os.Stdout)

	unaryExpressionFactory := func(primitive Symbol, operator UnaryOperator) ExpressionFactory {
		return func(generator ExpressionGenerator) Expression {
			return UnaryExpression{
				Expression: generator.Generate(primitive),
				Operator:   operator,
			}
		}
	}

	binaryExpressionFactory := func(lhs Symbol, operator BinaryOperator, rhs Symbol) ExpressionFactory {
		return func(generator ExpressionGenerator) Expression {
			return BinaryExpression{
				Lhs:      generator.Generate(lhs),
				Operator: operator,
				Rhs:      generator.Generate(rhs),
			}
		}
	}

	generator := ExpressionGenerator{
		Factories: map[Symbol]ExpressionFactory{
			// Numeric operators.
			numericNegationFunctionSymbol:   unaryExpressionFactory(int32TypeSymbol, NumericNegation),
			logicalNegationFunctionSymbol:   unaryExpressionFactory(booleanTypeSymbol, LogicalNegation),
			bitwiseComplementFunctionSymbol: unaryExpressionFactory(int32TypeSymbol, BitwiseComplement),

			AdditionFunctionSymbol:       binaryExpressionFactory(int32TypeSymbol, Addition, int32TypeSymbol),
			SubtractionFunctionSymbol:    binaryExpressionFactory(int32TypeSymbol, Subtraction, int32TypeSymbol),
			MultiplicationFunctionSymbol: binaryExpressionFactory(int32TypeSymbol, Multiplication, int32TypeSymbol),
			DivisionFunctionSymbol:       binaryExpressionFactory(int32TypeSymbol, Division, int32TypeSymbol),
			RemainderFunctionSymbol:      binaryExpressionFactory(int32TypeSymbol, Remainder, int32TypeSymbol),

			LessThanFunctionSymbol:           binaryExpressionFactory(int32TypeSymbol, LessThan, int32TypeSymbol),
			LessThanOrEqualFunctionSymbol:    binaryExpressionFactory(int32TypeSymbol, LessThanOrEqual, int32TypeSymbol),
			GreaterThanFunctionSymbol:        binaryExpressionFactory(int32TypeSymbol, GreaterThan, int32TypeSymbol),
			GreaterThanOrEqualFunctionSymbol: binaryExpressionFactory(int32TypeSymbol, GreaterThanOrEqual, int32TypeSymbol),

			ConcatenationFunctionSymbol: binaryExpressionFactory(stringTypeSymbol, Concatenation, stringTypeSymbol),

			// Equality operators.
			EqualityFunctionSymbol:   binaryExpressionFactory(anyTypeSymbol, Equality, anyTypeSymbol),
			InequalityFunctionSymbol: binaryExpressionFactory(anyTypeSymbol, Inequality, anyTypeSymbol),

			// Boolean operators.
			LogicalDisjunctionFunctionSymbol: binaryExpressionFactory(booleanTypeSymbol, LogicalDisjunction, booleanTypeSymbol),
			LogicalConjunctionFunctionSymbol: binaryExpressionFactory(booleanTypeSymbol, LogicalConjunction, booleanTypeSymbol),

			// Bitable operators.
			BitwiseDisjunctionFunctionSymbol:          binaryExpressionFactory(int32TypeSymbol, BitwiseDisjunction, int32TypeSymbol),
			BitwiseConjunctionFunctionSymbol:          binaryExpressionFactory(int32TypeSymbol, BitwiseConjunction, int32TypeSymbol),
			BitwizeExclusiveDisjunctionFunctionSymbol: binaryExpressionFactory(int32TypeSymbol, BitwizeExclusiveDisjunction, int32TypeSymbol),
			BitwiseLeftShiftFunctionSymbol:            binaryExpressionFactory(int32TypeSymbol, BitwiseLeftShift, int32TypeSymbol),
			BitwiseRightShiftFunctionSymbol:           binaryExpressionFactory(int32TypeSymbol, BitwiseRightShift, int32TypeSymbol),
			BitwiseClearFunctionSymbol:                binaryExpressionFactory(int32TypeSymbol, BitwiseClear, int32TypeSymbol),

			// Random generators.
			randomIntFuncSymbol: func(generator ExpressionGenerator) Expression {
				return ConstantInt32{
					Value: 42,
				}
			},
			randomBoolFuncSymbol: func(generator ExpressionGenerator) Expression {
				return ConstantBoolean{
					Value: false,
				}
			},
			randomStringFuncSymbol: func(generator ExpressionGenerator) Expression {
				return ConstantString{
					Value: "Hello, World",
				}
			},
		},
		Graph: graph,
	}

	expression := generator.Generate(booleanTypeSymbol)

	emitter := &ExpressionEmitter{
		writer: os.Stdout,
	}
	expression.Accept(emitter)
}

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

type Types struct {
	Relations map[Symbol][]Symbol
}

func (types Types) Subtypes(symbol Symbol) (subtypes []Symbol) {
	subtypes = types.Relations[symbol]
	return
}

func (types Types) IsAssignable(from, to Symbol) bool {
	if from == to {
		return true
	}

	for _, symbol := range types.Subtypes(from) {
		if symbol == to {
			return true
		}
	}
	return false
}

type FunctionSignature struct {
	Symbol     Symbol
	Parameters []Symbol
	Type       Symbol
}

func CreateNullaryFunctionSignature(symbol, result Symbol) FunctionSignature {
	return FunctionSignature{
		Symbol:     symbol,
		Parameters: []Symbol{},
		Type:       result,
	}
}

func CreateUnaryFunctionSignature(symbol, before, after Symbol) FunctionSignature {
	return FunctionSignature{
		Symbol:     symbol,
		Parameters: []Symbol{before},
		Type:       after,
	}
}

func CreateBinaryFunctionSignature(symbol, lhs, rhs, result Symbol) FunctionSignature {
	return FunctionSignature{
		Symbol:     symbol,
		Parameters: []Symbol{lhs, rhs},
		Type:       result,
	}
}

type FunctionSignatureGraph struct {
	Signatures []FunctionSignature
	Types      Types
}

func (graph FunctionSignatureGraph) Computes(primitive Symbol) (selection []FunctionSignature) {
	for _, signature := range graph.Signatures {
		if graph.Types.IsAssignable(signature.Type, primitive) {
			selection = append(selection, signature)
		}
	}
	return
}

func (graph FunctionSignatureGraph) Requires(primitive Symbol) (selection []FunctionSignature) {
	for _, signature := range graph.Signatures {
		for _, parameter := range signature.Parameters {
			if graph.Types.IsAssignable(primitive, parameter) {
				selection = append(selection, signature)
				break
			}
		}
	}
	return
}

func (graph FunctionSignatureGraph) Dot(symbols SymbolTable, writer io.Writer) {
	writer.Write([]byte("digraph G {\n"))

	// Write subgraph containing all the types.
	writer.Write([]byte("subgraph cluster_0 {\n"))
	writer.Write([]byte("style=filled;\n"))
	writer.Write([]byte("color=lightgrey;\n"))
	writer.Write([]byte("node [style=filled,shape=rectanlge,color=white];\n"))

	knownTypes := make(map[Symbol]interface{})

	for _, signature := range graph.Signatures {
		for _, parameter := range signature.Parameters {
			if _, exists := knownTypes[parameter]; exists {
				continue
			}
			knownTypes[parameter] = nil

			typeName, exists := symbols.Lookup(parameter)
			if !exists {
				panic("The symbol should always exist: " + strconv.Itoa(int(parameter)))
			}

			writer.Write([]byte(fmt.Sprintln(
				parameter,
				"[",
				"label=\""+typeName+"\"",
				"]",
			)))
		}

		if _, exists := knownTypes[signature.Type]; exists {
			continue
		}
		knownTypes[signature.Type] = nil

		typeName, exists := symbols.Lookup(signature.Type)
		if !exists {
			panic("The symbol should always exist: " + strconv.Itoa(int(signature.Type)))
		}

		writer.Write([]byte(fmt.Sprintln(
			signature.Type,
			"[",
			"label=\""+typeName+"\"",
			"]",
		)))
	}

	writer.Write([]byte("label = \"Types\";\n"))
	writer.Write([]byte("}\n"))

	// Write all function signatures.
	for _, signature := range graph.Signatures {
		name, exists := symbols.Lookup(signature.Symbol)
		if !exists {
			panic("The symbol should always exist: " + strconv.Itoa(int(signature.Symbol)))
		}

		// Draw signature node.
		writer.Write([]byte(fmt.Sprintln(
			signature.Symbol,
			"[",
			"label=\""+name+"\"",
			",shape=\"rectangle\"",
			"]",
		)))

		// Draw ingoing edges (from requried types).
		for _, parameter := range signature.Parameters {
			writer.Write([]byte(fmt.Sprintln(
				parameter,
				"->",
				signature.Symbol,
			)))
		}

		// Draw outgoing edges (To resulting type).
		writer.Write([]byte(fmt.Sprintln(
			signature.Symbol,
			"->",
			signature.Type,
		)))
	}

	writer.Write([]byte("}\n"))
}

type ExpressionFactory func(generator ExpressionGenerator) Expression

type ExpressionGenerator struct {
	Factories map[Symbol]ExpressionFactory
	Graph     FunctionSignatureGraph
	Depth     uint
}

func (generator ExpressionGenerator) Generate(primitive Symbol) (expression Expression) {
	roots := generator.Graph.Computes(primitive)

	if generator.Depth >= 20 {
		for _, root := range roots {
			if len(root.Parameters) == 0 {
				factory, exists := generator.Factories[root.Symbol]
				if !exists {
					panic("Should always have atleast one factory " + strconv.Itoa(int(root.Symbol)))
				}

				return factory(generator)
			}
		}
	}
	generator.Depth += 1

	root := roots[rand.Intn(len(roots))]

	factory, exists := generator.Factories[root.Symbol]
	if !exists {
		panic("Should always have atleast one factory " + strconv.Itoa(int(root.Symbol)))
	}

	return factory(generator)
}

type ExpressionEmitter struct {
	writer io.Writer
}

func (emitter *ExpressionEmitter) write(str string) {
	emitter.writer.Write([]byte(str))
}

func (emitter *ExpressionEmitter) VisitConstantBoolean(constant ConstantBoolean) {
	emitter.write(fmt.Sprint(constant.Value))
}

func (emitter *ExpressionEmitter) VisitConstantInt32(constant ConstantInt32) {
	emitter.write(fmt.Sprint(constant.Value))
}

func (emitter *ExpressionEmitter) VisitConstantString(constant ConstantString) {
	emitter.write("\"" + fmt.Sprint(constant.Value) + "\"")
}

func (emitter *ExpressionEmitter) VisitUnary(unary UnaryExpression) {
	emitter.write("(")
	switch unary.Operator {
	case NumericNegation:
		emitter.write("-")
	case LogicalNegation:
		emitter.write("!")
	case BitwiseComplement:
		emitter.write("^")
	case Channel:
		emitter.write("<-")
	case Dereference:
		emitter.write("*")
	}
	unary.Expression.Accept(emitter)
	emitter.write(")")
}

func (emitter *ExpressionEmitter) VisitBinary(binary BinaryExpression) {
	emitter.write("(")
	binary.Lhs.Accept(emitter)
	switch binary.Operator {
	// Arithmetic
	case Addition:
		emitter.write("+")
	case Subtraction:
		emitter.write("-")
	case Multiplication:
		emitter.write("*")
	case Division:
		emitter.write("/")
	case Remainder:
		emitter.write("%")
	// Strings
	case Concatenation:
		emitter.write("+")
	// Equality
	case Equality:
		emitter.write("==")
	case Inequality:
		emitter.write("!=")
	case LessThan:
		emitter.write("<")
	case LessThanOrEqual:
		emitter.write("<=")
	case GreaterThan:
		emitter.write(">")
	case GreaterThanOrEqual:
		emitter.write(">=")
	// Logical
	case LogicalDisjunction:
		emitter.write("||")
	case LogicalConjunction:
		emitter.write("&&")
	// Bitwise
	case BitwiseDisjunction:
		emitter.write("|")
	case BitwiseConjunction:
		emitter.write("&")
	case BitwizeExclusiveDisjunction:
		emitter.write("^")
	case BitwiseLeftShift:
		emitter.write("<<")
	case BitwiseRightShift:
		emitter.write(">>")
	case BitwiseClear:
		emitter.write("&^")
	}
	binary.Rhs.Accept(emitter)
	emitter.write(")")
}

type ExpressionVisitor interface {
	VisitConstantBoolean(constant ConstantBoolean)
	VisitConstantInt32(constant ConstantInt32)
	VisitConstantString(constant ConstantString)
	VisitUnary(unary UnaryExpression)
	VisitBinary(binary BinaryExpression)
}

type BinaryOperator int

const (
	// Arithmetic
	Addition = iota
	Subtraction
	Multiplication
	Division
	Remainder
	// Strings
	Concatenation
	// Equality
	Equality
	Inequality
	LessThan
	LessThanOrEqual
	GreaterThan
	GreaterThanOrEqual
	// Logical
	LogicalDisjunction
	LogicalConjunction
	// Bitwise
	BitwiseDisjunction
	BitwiseConjunction
	BitwizeExclusiveDisjunction
	BitwiseLeftShift
	BitwiseRightShift
	BitwiseClear
)

type BinaryExpression struct {
	Lhs      Expression
	Operator BinaryOperator
	Rhs      Expression
}

func (binary BinaryExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitBinary(binary)
}

type UnaryOperator int

const (
	NumericNegation = iota
	LogicalNegation
	BitwiseComplement
	Channel
	Dereference
)

type UnaryExpression struct {
	Operator   UnaryOperator
	Expression Expression
}

func (unary UnaryExpression) Accept(visitor ExpressionVisitor) {
	visitor.VisitUnary(unary)
}

type ConstantBoolean struct {
	Value bool
}

func (constant ConstantBoolean) Accept(visitor ExpressionVisitor) {
	visitor.VisitConstantBoolean(constant)
}

type ConstantInt32 struct {
	Value int32
}

func (constant ConstantInt32) Accept(visitor ExpressionVisitor) {
	visitor.VisitConstantInt32(constant)
}

type ConstantString struct {
	Value string
}

func (constant ConstantString) Accept(visitor ExpressionVisitor) {
	visitor.VisitConstantString(constant)
}

type Expression interface {
	Accept(visitor ExpressionVisitor)
}
