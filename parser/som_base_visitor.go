// Code generated from ./SOM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SOM

import "github.com/antlr4-go/antlr/v4"

type BaseSOMVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSOMVisitor) VisitClassdef(ctx *ClassdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitSuperclass(ctx *SuperclassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitInstanceFields(ctx *InstanceFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitClassFields(ctx *ClassFieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitMethod(ctx *MethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitUnaryPattern(ctx *UnaryPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBinaryPattern(ctx *BinaryPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitKeywordPattern(ctx *KeywordPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitMethodBlock(ctx *MethodBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitUnarySelector(ctx *UnarySelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBinarySelector(ctx *BinarySelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBlockContents(ctx *BlockContentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLocalDefs(ctx *LocalDefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBlockBody(ctx *BlockBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitResult(ctx *ResultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitAssignation(ctx *AssignationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitAssignments(ctx *AssignmentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitEvaluation(ctx *EvaluationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitMessages(ctx *MessagesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitUnaryMessage(ctx *UnaryMessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBinaryMessage(ctx *BinaryMessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBinaryOperand(ctx *BinaryOperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitKeywordMessage(ctx *KeywordMessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitFormula(ctx *FormulaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitNestedTerm(ctx *NestedTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralArray(ctx *LiteralArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralNumber(ctx *LiteralNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralDecimal(ctx *LiteralDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitNegativeDecimal(ctx *NegativeDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralInteger(ctx *LiteralIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralDouble(ctx *LiteralDoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralSymbol(ctx *LiteralSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitLiteralString(ctx *LiteralStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitKeywordSelector(ctx *KeywordSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitSomstring(ctx *SomstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitNestedBlock(ctx *NestedBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBlockPattern(ctx *BlockPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSOMVisitor) VisitBlockArguments(ctx *BlockArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}
