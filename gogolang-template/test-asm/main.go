package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// 编译 生成抽象语法树
func main() {
	src := `
	package main

	import "fmt"
	
	var a int = 1
	
	const b int = 10
	
	func main() {
		for i := 0; i < b; i++ {
			fmt.Println("Hello world", a, b)
		}
	}
	`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Print(fset, f)
}
