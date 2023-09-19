// Code generated from ./SOM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SOM

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SOMParser.
type SOMVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SOMParser#classdef.
	VisitClassdef(ctx *ClassdefContext) interface{}

	// Visit a parse tree produced by SOMParser#superclass.
	VisitSuperclass(ctx *SuperclassContext) interface{}

	// Visit a parse tree produced by SOMParser#instanceFields.
	VisitInstanceFields(ctx *InstanceFieldsContext) interface{}

	// Visit a parse tree produced by SOMParser#classFields.
	VisitClassFields(ctx *ClassFieldsContext) interface{}

	// Visit a parse tree produced by SOMParser#method.
	VisitMethod(ctx *MethodContext) interface{}

	// Visit a parse tree produced by SOMParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by SOMParser#unaryPattern.
	VisitUnaryPattern(ctx *UnaryPatternContext) interface{}

	// Visit a parse tree produced by SOMParser#binaryPattern.
	VisitBinaryPattern(ctx *BinaryPatternContext) interface{}

	// Visit a parse tree produced by SOMParser#keywordPattern.
	VisitKeywordPattern(ctx *KeywordPatternContext) interface{}

	// Visit a parse tree produced by SOMParser#methodBlock.
	VisitMethodBlock(ctx *MethodBlockContext) interface{}

	// Visit a parse tree produced by SOMParser#unarySelector.
	VisitUnarySelector(ctx *UnarySelectorContext) interface{}

	// Visit a parse tree produced by SOMParser#binarySelector.
	VisitBinarySelector(ctx *BinarySelectorContext) interface{}

	// Visit a parse tree produced by SOMParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SOMParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by SOMParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by SOMParser#blockContents.
	VisitBlockContents(ctx *BlockContentsContext) interface{}

	// Visit a parse tree produced by SOMParser#localDefs.
	VisitLocalDefs(ctx *LocalDefsContext) interface{}

	// Visit a parse tree produced by SOMParser#blockBody.
	VisitBlockBody(ctx *BlockBodyContext) interface{}

	// Visit a parse tree produced by SOMParser#result.
	VisitResult(ctx *ResultContext) interface{}

	// Visit a parse tree produced by SOMParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SOMParser#assignation.
	VisitAssignation(ctx *AssignationContext) interface{}

	// Visit a parse tree produced by SOMParser#assignments.
	VisitAssignments(ctx *AssignmentsContext) interface{}

	// Visit a parse tree produced by SOMParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SOMParser#evaluation.
	VisitEvaluation(ctx *EvaluationContext) interface{}

	// Visit a parse tree produced by SOMParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by SOMParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SOMParser#messages.
	VisitMessages(ctx *MessagesContext) interface{}

	// Visit a parse tree produced by SOMParser#unaryMessage.
	VisitUnaryMessage(ctx *UnaryMessageContext) interface{}

	// Visit a parse tree produced by SOMParser#binaryMessage.
	VisitBinaryMessage(ctx *BinaryMessageContext) interface{}

	// Visit a parse tree produced by SOMParser#binaryOperand.
	VisitBinaryOperand(ctx *BinaryOperandContext) interface{}

	// Visit a parse tree produced by SOMParser#keywordMessage.
	VisitKeywordMessage(ctx *KeywordMessageContext) interface{}

	// Visit a parse tree produced by SOMParser#formula.
	VisitFormula(ctx *FormulaContext) interface{}

	// Visit a parse tree produced by SOMParser#nestedTerm.
	VisitNestedTerm(ctx *NestedTermContext) interface{}

	// Visit a parse tree produced by SOMParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by SOMParser#literalArray.
	VisitLiteralArray(ctx *LiteralArrayContext) interface{}

	// Visit a parse tree produced by SOMParser#literalNumber.
	VisitLiteralNumber(ctx *LiteralNumberContext) interface{}

	// Visit a parse tree produced by SOMParser#literalDecimal.
	VisitLiteralDecimal(ctx *LiteralDecimalContext) interface{}

	// Visit a parse tree produced by SOMParser#negativeDecimal.
	VisitNegativeDecimal(ctx *NegativeDecimalContext) interface{}

	// Visit a parse tree produced by SOMParser#literalInteger.
	VisitLiteralInteger(ctx *LiteralIntegerContext) interface{}

	// Visit a parse tree produced by SOMParser#literalDouble.
	VisitLiteralDouble(ctx *LiteralDoubleContext) interface{}

	// Visit a parse tree produced by SOMParser#literalSymbol.
	VisitLiteralSymbol(ctx *LiteralSymbolContext) interface{}

	// Visit a parse tree produced by SOMParser#literalString.
	VisitLiteralString(ctx *LiteralStringContext) interface{}

	// Visit a parse tree produced by SOMParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by SOMParser#keywordSelector.
	VisitKeywordSelector(ctx *KeywordSelectorContext) interface{}

	// Visit a parse tree produced by SOMParser#somstring.
	VisitSomstring(ctx *SomstringContext) interface{}

	// Visit a parse tree produced by SOMParser#nestedBlock.
	VisitNestedBlock(ctx *NestedBlockContext) interface{}

	// Visit a parse tree produced by SOMParser#blockPattern.
	VisitBlockPattern(ctx *BlockPatternContext) interface{}

	// Visit a parse tree produced by SOMParser#blockArguments.
	VisitBlockArguments(ctx *BlockArgumentsContext) interface{}
}
