package object

import (
	"fmt"
	"bytes"
	"strings"
	"interpreter/ast"
)

type ObjectType string

// ObjectTypeの内容. データ型に相当
const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION_OBJ"
)
    
type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (o *Integer) Inspect() string  { return fmt.Sprintf("%d", o.Value) }
func (o *Integer) Type() ObjectType { return INTEGER_OBJ }
 
type Boolean struct {
	Value bool
}

func (o *Boolean) Inspect() string  { return fmt.Sprintf("%t", o.Value) }
func (o *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
 
type Null struct{}

func (o *Null) Inspect() string  { return "null" }
func (o *Null) Type() ObjectType { return NULL_OBJ }
 
// CallExpressionを評価するとこれになる
type ReturnValue struct {
	Value Object
}

func (o *ReturnValue) Inspect() string { return o.Value.Inspect() }
func (o *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

type Error struct {
	Message string
}
  
func (o *Error) Inspect() string { return "ERROR: " + o.Message }

func (o *Error) Type() ObjectType { return ERROR_OBJ }


type Function struct {
	Parameters []*ast.Identifier
	Body *ast.BlockStatement
	Env *Environment
}            
                
func (o *Function) Inspect() string { 
 	var out bytes.Buffer

	params := []string{}
	for _, p := range o.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(o.Body.String())
	out.WriteString("\n")

	return out.String()
}

func (o *Function) Type() ObjectType { return FUNCTION_OBJ }

