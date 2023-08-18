package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/bbogdan95/monkey-go/compiler"
	"github.com/bbogdan95/monkey-go/lexer"
	"github.com/bbogdan95/monkey-go/object"
	"github.com/bbogdan95/monkey-go/parser"
	"github.com/bbogdan95/monkey-go/repl"
	"github.com/bbogdan95/monkey-go/vm"
)

func main() {

	useRepl := flag.Bool("repl", false, "run the repl")
	flag.Parse()

	if *useRepl {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is the Monkey programming language!\n",
			user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	} else {
		var fileContent string
		var err error

		if len(os.Args) != 2 {
			fmt.Println("Woops! No argument provided")
			return
		} else {
			filename := os.Args[1]
			fileContent, err = readFile(filename)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		constants := []object.Object{}
		globals := make([]object.Object, vm.GlobalsSize)
		symbolTable := compiler.NewSymbolTable()

		for i, v := range object.Builtins {
			symbolTable.DefineBuiltin(i, v.Name)
		}

		l := lexer.New(fileContent)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, msg := range p.Errors() {
				fmt.Println("Woops! Parsing failed:")
				fmt.Println(msg)
			}
		}

		comp := compiler.NewWithState(symbolTable, constants)
		err = comp.Compile(program)
		if err != nil {
			fmt.Printf("Woops! Compilation failed:\n%s\n", err)
		}

		code := comp.Bytecode()

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Printf("Woops! Executing bytecode failed:\n%s\n", err)
		}

		lastPopped := machine.LastPoppedStackElem()
		fmt.Println(lastPopped.Inspect())
	}
}

func readFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
