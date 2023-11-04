package lang

type Type struct {
	identifier Symbol
}

type Types struct {
	mapping map[Symbol]Type
}

func (types *Types) Lookup(symbol Symbol) (Type, bool) {
	t, exists := types.mapping[symbol]
	return t, exists
}

type Generic struct {
	identifier Symbol
	typeSet    Types
}

func (generic *Generic) Concretions(tree *TypeTree) (concretions []Symbol) {
	for _, t := range generic.typeSet.mapping {
		concretions = append(concretions, tree.ConcretionsOf(t.identifier)...)
	}

	return
}

// A tree describing the heurachical relationships between types.
type TypeTree struct {
	/* "struct -implements-> interface" relationship.
	 *   E.g., "string" and "int" can be "any" as it is "interface {}"
	 *   "string" -> "any", "int" -> "any" */
	relations map[Symbol][]Symbol
}

func (tree *TypeTree) IsAbstraction(symbol Symbol) bool {
	return !tree.IsConcretion(symbol)
}

func (tree *TypeTree) IsConcretion(symbol Symbol) bool {
	_, exists := tree.relations[symbol]
	return exists
}

func (tree *TypeTree) AbstractionsFor(concretion Symbol) (interfaces []Symbol) {
	interfaces = tree.relations[concretion]
	return
}

func (tree *TypeTree) ConcretionsOf(target Symbol) (subTypes []Symbol) {
	for concretion, abstractions := range tree.relations {
		for _, abstration := range abstractions {
			if abstration == target {
				subTypes = append(subTypes, concretion)
				break
			}
		}
	}

	return
}

func (tree *TypeTree) IsAssignable(rhs, lhs Symbol) bool {
	// Interface assigned to struct ERROR.
	if tree.IsAbstraction(rhs) && tree.IsConcretion(lhs) {
		return false
	}

	// Interface assigned to interface OK if from == to.
	// Struct assigned to struct OK if from == to.
	if lhs == rhs {
		return true
	}

	// Struct assigned to interface OK if struct implements interface.
	for _, abstraction := range tree.AbstractionsFor(rhs) {
		if abstraction == lhs {
			return true
		}
	}

	return false
}
