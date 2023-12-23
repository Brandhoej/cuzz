package lang

import (
	"errors"
	"math/rand"
)

var ErrCannotGenerateExpressionForAbstraction = errors.New("failed generation")

type ExpressionGenerator interface {
	Generate(t Type) Expression
}

type RndExpressionGenerator struct {
	functions FunctionSet
	typeTree  TypeTree
	types     Types
	factories map[Symbol]func(parameters ...Expression) Expression
}

func (generator *RndExpressionGenerator) ChooseType(symbols []Symbol) Type {
	index := rand.Intn(len(symbols))
	symbol := symbols[index]
	t, _ := generator.types.Lookup(symbol)
	return t
}

func (generator *RndExpressionGenerator) ChooseParameterConcretion(
	generics map[Symbol]Type, parameter Symbol,
) Type {
	if generic, isGeneric := generics[parameter]; isGeneric {
		return generic
	}

	if generator.typeTree.IsAbstraction(parameter) {
		return generator.ChooseType(
			generator.typeTree.ConcretionsOf(parameter),
		)
	}

	// TODO: What about errors?
	t, _ := generator.types.Lookup(parameter)
	return t
}

func (generator *RndExpressionGenerator) ChooseParameterConcretions(
	generics map[Symbol]Type, function Function,
) []Type {
	types := make([]Type, len(function.parameters))

	for idx, parameter := range function.parameters {
		types[idx] = generator.ChooseParameterConcretion(
			generics, parameter,
		)
	}

	return types
}

func (generator *RndExpressionGenerator) ChooseGenericConcretion(generic Generic) Type {
	return generator.ChooseType(
		generic.Concretions(&generator.typeTree),
	)
}

func (generator *RndExpressionGenerator) ChooseGenericConcretions(
	generics []Generic,
) map[Symbol]Type {
	concretions := make(map[Symbol]Type, len(generics))

	for _, generic := range generics {
		concretions[generic.identifier] = generator.ChooseGenericConcretion(
			generic,
		)
	}

	return concretions
}

func (generator *RndExpressionGenerator) ChooseFunctionThatComputes(t Type) Function {
	subset := generator.functions.Computes(t, generator.typeTree)
	functions := subset.set
	index := rand.Intn(len(functions))
	return functions[index]
}

func (generator *RndExpressionGenerator) CreateExpression(symbol Symbol, parameters ...Expression) Expression {
	return generator.factories[symbol](parameters...)
}

func (generator *RndExpressionGenerator) Generate(target Type) (Expression, error) {
	// Abstractions cannot be constructed in an expression.
	if generator.typeTree.IsAbstraction(target.identifier) {
		return nil, ErrCannotGenerateExpressionForAbstraction
	}

	function := generator.ChooseFunctionThatComputes(target)

	// Exit early if there are no need for sub-expressions (Parameters).
	if function.IsEmpty() {
		return generator.CreateExpression(function.identifier), nil
	}

	// Step 1: Pick random types for generics.
	generics := generator.ChooseGenericConcretions(
		function.generics,
	)

	// Step 2: Find formal parameter types.
	var parameterTypes []Type = generator.ChooseParameterConcretions(
		generics, function,
	)

	// Step 3: Generate expression for the corresponding types.
	var parameters []Expression = make([]Expression, len(function.parameters))
	for idx, parameterType := range parameterTypes {
		var err error
		parameters[idx], err = generator.Generate(parameterType)

		if err != nil {
			return nil, err
		}
	}

	// Step 4: Finalize the function call.
	return generator.CreateExpression(
		function.identifier, parameters...,
	), nil
}
