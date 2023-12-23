package lang

type Function struct {
	identifier Symbol
	generics   []Generic
	parameters []Symbol
	returnType Symbol
}

func (function *Function) Asd() {

}

func (function *Function) IsEmpty() bool {
	return len(function.parameters) == 0
}

type FunctionSet struct {
	set []Function
}

func (set *FunctionSet) Computes(t Type, types TypeTree) (subset FunctionSet) {
	for _, function := range set.set {
		if types.IsAssignable(function.returnType, t.identifier) {
			subset.set = append(subset.set, function)
		}
	}

	return
}
