package test

import (
	"context"

	"github.com/brandhoej/cuzz/internal/generational"
)

type Result struct{}

type Execution[Input, Output any] struct {
	input  Input
	output Output
}

type Case[Parameter any] interface {
	Test(context context.Context, generator generational.Generator[Parameter]) (Result, error)
}
