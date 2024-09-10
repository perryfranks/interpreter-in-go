package entries

import (
	"fmt"
	"os"
	"os/user"

	"github.com/perryfranks/monkey-lang/lexer"
	"github.com/perryfranks/monkey-lang/parser"
	"github.com/perryfranks/monkey-lang/repl"
)

func ReplRun() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)

	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func EvalErrors(input string) (progErrors []string, err error) {

	l := lexer.New(input)
	p := parser.New(l)
	_ = p.ParseProgram() // you must parse the program to generate errors

	return p.Errors(), nil
}
