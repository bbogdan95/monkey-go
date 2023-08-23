package main

import (
	"fmt"
	"syscall/js"

	"github.com/bbogdan95/monkey-go/compiler"
	"github.com/bbogdan95/monkey-go/lexer"
	"github.com/bbogdan95/monkey-go/object"
	"github.com/bbogdan95/monkey-go/parser"
	"github.com/bbogdan95/monkey-go/vm"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("execute", js.FuncOf(execute))

	<-c
}

//export execute
func execute(this js.Value, s []js.Value) interface{} {
	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()

	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	sourceCode := fmt.Sprintf("%s", s[0])
	l := lexer.New(sourceCode)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Println("Woops! Parsing failed:")
			fmt.Println(msg)
		}
	}

	comp := compiler.NewWithState(symbolTable, constants)
	err := comp.Compile(program)
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
	return js.ValueOf(lastPopped.Inspect())
}
