package pipeline

import "context"

type Pipe[In, Out any] struct {
	function func(context context.Context, input In) (Out, error)
}

func Adapt[In, Out any](
	function func(context context.Context, input In) (Out, error),
) Pipe[In, Out] {
	return Pipe[In, Out]{
		function,
	}
}

func Connect[In, Aux, Out any](
	first Pipe[In, Aux],
	last Pipe[Aux, Out],
) Pipe[In, Out] {
	return Pipe[In, Out]{
		function: func(context context.Context, input In) (Out, error) {
			auxVal, err := first.Execute(context, input)
			if err != nil {
				var zeroOut Out
				return zeroOut, err
			}

			return last.Execute(context, auxVal)
		},
	}
}

func (pipe Pipe[In, Out]) Execute(context context.Context, input In) (Out, error) {
	return pipe.function(context, input)
}
