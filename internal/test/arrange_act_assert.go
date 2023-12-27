package test

import (
	"context"
	"errors"

	"github.com/brandhoej/cuzz/internal/pipeline"
)

var (
	ErrArranging = errors.New("an error was encoutered when arranging the input")
	ErrActing    = errors.New("an error was encoutered when acting the input")
	ErrAsserting = errors.New("an error was encoutered when asserting the input")
)

type ArrangeActAssert[Parameter, Input, Output any] struct {
	arrange pipeline.Pipe[Parameter, Input]
	act     pipeline.Pipe[Input, Output]
	assert  pipeline.Pipe[Execution[Input, Output], Result]
}

func (aaa *ArrangeActAssert[Parameter, Input, Output]) Test(
	context context.Context,
	parameter Parameter,
) (Result, error) {
	var result Result

	input, err := aaa.arrange.Execute(context, parameter)
	if err != nil {
		return result, errors.Join(ErrArranging, err)
	}

	output, err := aaa.act.Execute(context, input)
	if err != nil {
		return result, errors.Join(ErrActing, err)
	}

	execute := Execution[Input, Output]{input, output}
	result, err = aaa.assert.Execute(context, execute)
	if err != nil {
		return result, errors.Join(ErrAsserting, err)
	}

	return result, nil
}
