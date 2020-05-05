package runner

import (
	"interpreter/src/evaluator"
	"interpreter/src/lexer"
	"interpreter/src/object"
	"interpreter/src/parser"
	"io"
	"io/ioutil"
)

func Run(filepath string, in io.Reader, out io.Writer) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	input := string(bytes)

	env := object.NewEnvironment()
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
