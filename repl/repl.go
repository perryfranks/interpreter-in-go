package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/perryfranks/monkey-interpreter/evaluator"
	"github.com/perryfranks/monkey-interpreter/lexer"
	"github.com/perryfranks/monkey-interpreter/object"
	"github.com/perryfranks/monkey-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			// io.WriteString(out, program.String())
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

// func Start(in io.Reader, out io.Writer) {
// 	scanner := bufio.NewScanner(in)
//
// 	for {
// 		fmt.Fprintf(out, PROMPT)
// 		scanned := scanner.Scan()
// 		if !scanned {
// 			return
// 		}
//
// 		line := scanner.Text()
// 		l := lexer.New(line)
// 		p := parser.New(l)
//
// 		program := p.ParseProgram()
// 		if len(p.Errors()) != 0 {
// 			printParserErrors(out, p.Errors())
// 			continue
// 		}
//
// 		io.WriteString(out, program.String())
// 		io.WriteString(out, "\n")
// 	}
// }

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

// Read-lex-print-loop
// func Start(in io.Reader, out io.Writer) {
// 	scanner := bufio.NewScanner(in)
//
// 	for {
// 		fmt.Fprintf(out, PROMPT)
// 		scanned := scanner.Scan()
// 		if !scanned {
// 			return
// 		}
//
// 		line := scanner.Text()
// 		l := lexer.New(line)
//
// 		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
// 			fmt.Fprintf(out, "%+v\n", tok)
// 		}
// 	}
// }
