package test

type Oracle[Input, Output any] interface {
	Test(input Input, output Output) bool
}

type Active[Input, Output any] struct {
	oracle func(input Input, output Output) bool
}

func (active Active[Input, Output]) Test(input Input, output Output) bool {
	return (active.oracle)(input, output)
}

type Partition[Input, Output any] struct {
	filter func(input Input) bool
	oracle Oracle[Input, Output]
}

func (active Partition[Input, Output]) Test(input Input, output Output) bool {
	if active.filter(input) {
		return true
	}

	return active.oracle.Test(input, output)
}

type Composite[Input, Output any] struct {
	oracles []Oracle[Input, Output]
}

func (composite Composite[Input, Output]) Test(input Input, output Output) bool {
	for idx := range composite.oracles {
		if !composite.oracles[idx].Test(input, output) {
			return false
		}
	}

	return true
}
