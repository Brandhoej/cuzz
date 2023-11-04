package lang

import (
	"fmt"
	"io"
)

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
