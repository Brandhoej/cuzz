package generational

import (
	"reflect"
	"testing"
)

func blueprintGeneratorSpecification[T any](name string, generator Generator[T], t *testing.T) {
	blueprint := Blueprint[T]{
		specification: generator,
	}

	actual := blueprint.Generator()
	if !reflect.DeepEqual(
		reflect.ValueOf(generator).Elem().Interface(),
		reflect.ValueOf(actual).Elem().Interface(),
	) {
		t.Error(name, "blueprint generator was", actual, "but expected", generator)
	}
}

func TestBlueprintGenerator(t *testing.T) {
	blueprintGeneratorSpecification[int]("Empty int generator specificaiton", Sequence[int](), t)
	blueprintGeneratorSpecification[int]("Int generator specificaiton", Sequence[int](1), t)
	blueprintGeneratorSpecification[struct{ value int }](
		"Struct generator specificaiton",
		Sequence[struct{ value int }](struct{ value int }{value: 1}),
		t,
	)
}
