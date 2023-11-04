package lang

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
