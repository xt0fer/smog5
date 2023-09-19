// Code generated from ./SOM.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SOM

import "github.com/antlr4-go/antlr/v4"

// BaseSOMListener is a complete listener for a parse tree produced by SOMParser.
type BaseSOMListener struct{}

var _ SOMListener = &BaseSOMListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSOMListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSOMListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSOMListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSOMListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassdef is called when production classdef is entered.
func (s *BaseSOMListener) EnterClassdef(ctx *ClassdefContext) {}

// ExitClassdef is called when production classdef is exited.
func (s *BaseSOMListener) ExitClassdef(ctx *ClassdefContext) {}

// EnterSuperclass is called when production superclass is entered.
func (s *BaseSOMListener) EnterSuperclass(ctx *SuperclassContext) {}

// ExitSuperclass is called when production superclass is exited.
func (s *BaseSOMListener) ExitSuperclass(ctx *SuperclassContext) {}

// EnterInstanceFields is called when production instanceFields is entered.
func (s *BaseSOMListener) EnterInstanceFields(ctx *InstanceFieldsContext) {}

// ExitInstanceFields is called when production instanceFields is exited.
func (s *BaseSOMListener) ExitInstanceFields(ctx *InstanceFieldsContext) {}

// EnterClassFields is called when production classFields is entered.
func (s *BaseSOMListener) EnterClassFields(ctx *ClassFieldsContext) {}

// ExitClassFields is called when production classFields is exited.
func (s *BaseSOMListener) ExitClassFields(ctx *ClassFieldsContext) {}

// EnterMethod is called when production method is entered.
func (s *BaseSOMListener) EnterMethod(ctx *MethodContext) {}

// ExitMethod is called when production method is exited.
func (s *BaseSOMListener) ExitMethod(ctx *MethodContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSOMListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSOMListener) ExitPattern(ctx *PatternContext) {}

// EnterUnaryPattern is called when production unaryPattern is entered.
func (s *BaseSOMListener) EnterUnaryPattern(ctx *UnaryPatternContext) {}

// ExitUnaryPattern is called when production unaryPattern is exited.
func (s *BaseSOMListener) ExitUnaryPattern(ctx *UnaryPatternContext) {}

// EnterBinaryPattern is called when production binaryPattern is entered.
func (s *BaseSOMListener) EnterBinaryPattern(ctx *BinaryPatternContext) {}

// ExitBinaryPattern is called when production binaryPattern is exited.
func (s *BaseSOMListener) ExitBinaryPattern(ctx *BinaryPatternContext) {}

// EnterKeywordPattern is called when production keywordPattern is entered.
func (s *BaseSOMListener) EnterKeywordPattern(ctx *KeywordPatternContext) {}

// ExitKeywordPattern is called when production keywordPattern is exited.
func (s *BaseSOMListener) ExitKeywordPattern(ctx *KeywordPatternContext) {}

// EnterMethodBlock is called when production methodBlock is entered.
func (s *BaseSOMListener) EnterMethodBlock(ctx *MethodBlockContext) {}

// ExitMethodBlock is called when production methodBlock is exited.
func (s *BaseSOMListener) ExitMethodBlock(ctx *MethodBlockContext) {}

// EnterUnarySelector is called when production unarySelector is entered.
func (s *BaseSOMListener) EnterUnarySelector(ctx *UnarySelectorContext) {}

// ExitUnarySelector is called when production unarySelector is exited.
func (s *BaseSOMListener) ExitUnarySelector(ctx *UnarySelectorContext) {}

// EnterBinarySelector is called when production binarySelector is entered.
func (s *BaseSOMListener) EnterBinarySelector(ctx *BinarySelectorContext) {}

// ExitBinarySelector is called when production binarySelector is exited.
func (s *BaseSOMListener) ExitBinarySelector(ctx *BinarySelectorContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSOMListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSOMListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseSOMListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseSOMListener) ExitKeyword(ctx *KeywordContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseSOMListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseSOMListener) ExitArgument(ctx *ArgumentContext) {}

// EnterBlockContents is called when production blockContents is entered.
func (s *BaseSOMListener) EnterBlockContents(ctx *BlockContentsContext) {}

// ExitBlockContents is called when production blockContents is exited.
func (s *BaseSOMListener) ExitBlockContents(ctx *BlockContentsContext) {}

// EnterLocalDefs is called when production localDefs is entered.
func (s *BaseSOMListener) EnterLocalDefs(ctx *LocalDefsContext) {}

// ExitLocalDefs is called when production localDefs is exited.
func (s *BaseSOMListener) ExitLocalDefs(ctx *LocalDefsContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseSOMListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseSOMListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterResult is called when production result is entered.
func (s *BaseSOMListener) EnterResult(ctx *ResultContext) {}

// ExitResult is called when production result is exited.
func (s *BaseSOMListener) ExitResult(ctx *ResultContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSOMListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSOMListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignation is called when production assignation is entered.
func (s *BaseSOMListener) EnterAssignation(ctx *AssignationContext) {}

// ExitAssignation is called when production assignation is exited.
func (s *BaseSOMListener) ExitAssignation(ctx *AssignationContext) {}

// EnterAssignments is called when production assignments is entered.
func (s *BaseSOMListener) EnterAssignments(ctx *AssignmentsContext) {}

// ExitAssignments is called when production assignments is exited.
func (s *BaseSOMListener) ExitAssignments(ctx *AssignmentsContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSOMListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSOMListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterEvaluation is called when production evaluation is entered.
func (s *BaseSOMListener) EnterEvaluation(ctx *EvaluationContext) {}

// ExitEvaluation is called when production evaluation is exited.
func (s *BaseSOMListener) ExitEvaluation(ctx *EvaluationContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseSOMListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseSOMListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSOMListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSOMListener) ExitVariable(ctx *VariableContext) {}

// EnterMessages is called when production messages is entered.
func (s *BaseSOMListener) EnterMessages(ctx *MessagesContext) {}

// ExitMessages is called when production messages is exited.
func (s *BaseSOMListener) ExitMessages(ctx *MessagesContext) {}

// EnterUnaryMessage is called when production unaryMessage is entered.
func (s *BaseSOMListener) EnterUnaryMessage(ctx *UnaryMessageContext) {}

// ExitUnaryMessage is called when production unaryMessage is exited.
func (s *BaseSOMListener) ExitUnaryMessage(ctx *UnaryMessageContext) {}

// EnterBinaryMessage is called when production binaryMessage is entered.
func (s *BaseSOMListener) EnterBinaryMessage(ctx *BinaryMessageContext) {}

// ExitBinaryMessage is called when production binaryMessage is exited.
func (s *BaseSOMListener) ExitBinaryMessage(ctx *BinaryMessageContext) {}

// EnterBinaryOperand is called when production binaryOperand is entered.
func (s *BaseSOMListener) EnterBinaryOperand(ctx *BinaryOperandContext) {}

// ExitBinaryOperand is called when production binaryOperand is exited.
func (s *BaseSOMListener) ExitBinaryOperand(ctx *BinaryOperandContext) {}

// EnterKeywordMessage is called when production keywordMessage is entered.
func (s *BaseSOMListener) EnterKeywordMessage(ctx *KeywordMessageContext) {}

// ExitKeywordMessage is called when production keywordMessage is exited.
func (s *BaseSOMListener) ExitKeywordMessage(ctx *KeywordMessageContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseSOMListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseSOMListener) ExitFormula(ctx *FormulaContext) {}

// EnterNestedTerm is called when production nestedTerm is entered.
func (s *BaseSOMListener) EnterNestedTerm(ctx *NestedTermContext) {}

// ExitNestedTerm is called when production nestedTerm is exited.
func (s *BaseSOMListener) ExitNestedTerm(ctx *NestedTermContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSOMListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSOMListener) ExitLiteral(ctx *LiteralContext) {}

// EnterLiteralArray is called when production literalArray is entered.
func (s *BaseSOMListener) EnterLiteralArray(ctx *LiteralArrayContext) {}

// ExitLiteralArray is called when production literalArray is exited.
func (s *BaseSOMListener) ExitLiteralArray(ctx *LiteralArrayContext) {}

// EnterLiteralNumber is called when production literalNumber is entered.
func (s *BaseSOMListener) EnterLiteralNumber(ctx *LiteralNumberContext) {}

// ExitLiteralNumber is called when production literalNumber is exited.
func (s *BaseSOMListener) ExitLiteralNumber(ctx *LiteralNumberContext) {}

// EnterLiteralDecimal is called when production literalDecimal is entered.
func (s *BaseSOMListener) EnterLiteralDecimal(ctx *LiteralDecimalContext) {}

// ExitLiteralDecimal is called when production literalDecimal is exited.
func (s *BaseSOMListener) ExitLiteralDecimal(ctx *LiteralDecimalContext) {}

// EnterNegativeDecimal is called when production negativeDecimal is entered.
func (s *BaseSOMListener) EnterNegativeDecimal(ctx *NegativeDecimalContext) {}

// ExitNegativeDecimal is called when production negativeDecimal is exited.
func (s *BaseSOMListener) ExitNegativeDecimal(ctx *NegativeDecimalContext) {}

// EnterLiteralInteger is called when production literalInteger is entered.
func (s *BaseSOMListener) EnterLiteralInteger(ctx *LiteralIntegerContext) {}

// ExitLiteralInteger is called when production literalInteger is exited.
func (s *BaseSOMListener) ExitLiteralInteger(ctx *LiteralIntegerContext) {}

// EnterLiteralDouble is called when production literalDouble is entered.
func (s *BaseSOMListener) EnterLiteralDouble(ctx *LiteralDoubleContext) {}

// ExitLiteralDouble is called when production literalDouble is exited.
func (s *BaseSOMListener) ExitLiteralDouble(ctx *LiteralDoubleContext) {}

// EnterLiteralSymbol is called when production literalSymbol is entered.
func (s *BaseSOMListener) EnterLiteralSymbol(ctx *LiteralSymbolContext) {}

// ExitLiteralSymbol is called when production literalSymbol is exited.
func (s *BaseSOMListener) ExitLiteralSymbol(ctx *LiteralSymbolContext) {}

// EnterLiteralString is called when production literalString is entered.
func (s *BaseSOMListener) EnterLiteralString(ctx *LiteralStringContext) {}

// ExitLiteralString is called when production literalString is exited.
func (s *BaseSOMListener) ExitLiteralString(ctx *LiteralStringContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseSOMListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseSOMListener) ExitSelector(ctx *SelectorContext) {}

// EnterKeywordSelector is called when production keywordSelector is entered.
func (s *BaseSOMListener) EnterKeywordSelector(ctx *KeywordSelectorContext) {}

// ExitKeywordSelector is called when production keywordSelector is exited.
func (s *BaseSOMListener) ExitKeywordSelector(ctx *KeywordSelectorContext) {}

// EnterSomstring is called when production somstring is entered.
func (s *BaseSOMListener) EnterSomstring(ctx *SomstringContext) {}

// ExitSomstring is called when production somstring is exited.
func (s *BaseSOMListener) ExitSomstring(ctx *SomstringContext) {}

// EnterNestedBlock is called when production nestedBlock is entered.
func (s *BaseSOMListener) EnterNestedBlock(ctx *NestedBlockContext) {}

// ExitNestedBlock is called when production nestedBlock is exited.
func (s *BaseSOMListener) ExitNestedBlock(ctx *NestedBlockContext) {}

// EnterBlockPattern is called when production blockPattern is entered.
func (s *BaseSOMListener) EnterBlockPattern(ctx *BlockPatternContext) {}

// ExitBlockPattern is called when production blockPattern is exited.
func (s *BaseSOMListener) ExitBlockPattern(ctx *BlockPatternContext) {}

// EnterBlockArguments is called when production blockArguments is entered.
func (s *BaseSOMListener) EnterBlockArguments(ctx *BlockArgumentsContext) {}

// ExitBlockArguments is called when production blockArguments is exited.
func (s *BaseSOMListener) ExitBlockArguments(ctx *BlockArgumentsContext) {}
