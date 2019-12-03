package repl

import (
	"bufio"
	"fmt"
	"github.com/MoonShining/monkey-lan/evaluator"
	"github.com/MoonShining/monkey-lan/lexer"
	"github.com/MoonShining/monkey-lan/object"
	"github.com/MoonShining/monkey-lan/parser"
	"io"
)

const (
	HEADER = ">> "
	EXIT   = "exit"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(HEADER)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == EXIT {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
