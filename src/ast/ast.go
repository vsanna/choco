package ast

import (
	"bytes"
	"choco/src/token"
	"strings"
)

// astを構成するNode
type Node interface {
	TokenLiteral() string // debugでのみ用いる
	String() string       // debugで用いる
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*****************
* program
******************/
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*****************
* Statement
******************/

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (node *LetStatement) statementNode() {
}
func (node *LetStatement) TokenLiteral() string {
	return node.Token.Literal
}
func (node *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(node.TokenLiteral() + " ")
	out.WriteString(node.Name.String())
	out.WriteString(" = ")
	if node.Value != nil {
		out.WriteString(node.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (node *ReturnStatement) statementNode() {
}
func (node *ReturnStatement) TokenLiteral() string {
	return node.Token.Literal
}
func (node *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(node.TokenLiteral() + " ")

	if node.ReturnValue != nil {
		out.WriteString(node.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (node *ExpressionStatement) statementNode() {
}
func (node *ExpressionStatement) TokenLiteral() string {
	return node.Token.Literal
}
func (node *ExpressionStatement) String() string {
	if node.Expression != nil {
		return node.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (node *BlockStatement) expressionNode() {}

func (node *BlockStatement) TokenLiteral() string {
	return node.Token.Literal
}
func (node *BlockStatement) String() string {
	var out bytes.Buffer

	for _, stmt := range node.Statements {
		out.WriteString(stmt.String())
	}

	return out.String()
}

/*****************
* Expression
******************/
type Identifier struct {
	Token token.Token
	Value string
}

func (node *Identifier) expressionNode() {
}
func (node *Identifier) TokenLiteral() string {
	return node.Token.Literal
}
func (node *Identifier) String() string {
	return node.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64 // long
}

func (node *IntegerLiteral) expressionNode() {
}
func (node *IntegerLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *IntegerLiteral) String() string {
	return node.Token.Literal
}

type StringLiteral struct {
	Token token.Token
	Value string // long
}

func (node *StringLiteral) expressionNode() {
}
func (node *StringLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *StringLiteral) String() string {
	return node.Token.Literal
}

type Boolean struct {
	Token token.Token
	Value bool // long
}

func (node *Boolean) expressionNode() {
}
func (node *Boolean) TokenLiteral() string {
	return node.Token.Literal
}

func (node *Boolean) String() string {
	return node.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (node *PrefixExpression) expressionNode() {}

func (node *PrefixExpression) TokenLiteral() string {
	return node.Token.Literal
}

func (node *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(node.Operator)
	out.WriteString(node.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression

	Left Expression
}

func (node *InfixExpression) expressionNode() {}

func (node *InfixExpression) TokenLiteral() string {
	return node.Token.Literal
}
func (node *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(node.Left.String())
	out.WriteString(node.Operator)
	out.WriteString(node.Right.String())
	out.WriteString(")")
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement // TODO: nullableにするにはpointerにする?
	Alternative *BlockStatement
}

func (node *IfExpression) expressionNode() {}

func (node *IfExpression) TokenLiteral() string {
	return node.Token.Literal
}
func (node *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(node.Condition.String())
	out.WriteString(" ")
	out.WriteString(node.Consequence.String())

	if node.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(node.Alternative.String())
	}

	return out.String()
}

// 関数定義の方の式
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (node *FunctionLiteral) expressionNode() {}

func (node *FunctionLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range node.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(node.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(node.Body.String())
	return out.String()
}

// 関数呼び出しの式
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (node *CallExpression) expressionNode() {}

func (node *CallExpression) TokenLiteral() string {
	return node.Token.Literal
}
func (node *CallExpression) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range node.Arguments {
		params = append(params, p.String())
	}

	out.WriteString(node.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	return out.String()
}

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (node *ArrayLiteral) expressionNode() {}

func (node *ArrayLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *ArrayLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range node.Elements {
		params = append(params, p.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(params, ","))
	out.WriteString("]")
	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (node *IndexExpression) expressionNode() {}

func (node *IndexExpression) TokenLiteral() string {
	return node.Token.Literal
}
func (node *IndexExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(node.Left.String())
	out.WriteString("[")
	out.WriteString(node.Index.String())
	out.WriteString("])")
	return out.String()
}

type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

func (node *HashLiteral) expressionNode() {}

func (node *HashLiteral) TokenLiteral() string {
	return node.Token.Literal
}
func (node *HashLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, val := range node.Pairs {
		pairs = append(pairs, key.String()+":"+val.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ","))
	out.WriteString("}")
	return out.String()
}
