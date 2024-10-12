package main

import (
	"fmt"
	"interpreter/evaluator"
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"interpreter/repl"
	"interpreter/util"
	"io"
	"os"
	"os/user"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		runRepl()
	} else if len(args) == 1 {
		runOnFile(args[0])
	} else {
		fmt.Printf("Too many arguments. Expected 0 or 1, got %d", len(args))
		os.Exit(1)
	}

}

func runOnFile(fpath string) {
	f, err := os.ReadFile(fpath)
	if err != nil {
		panic(err)
	}
	fString := string(f)

	l := lexer.New(fString)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		util.PrintParserErrors(os.Stdout, p.Errors())
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		printIfError(os.Stdout, evaluated)
	}

}

func printIfError(out io.Writer, obj object.Object) {
	if obj.Type() == object.ERROR_OBJ {
		io.WriteString(out, obj.Inspect())
	}
}

func runRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
