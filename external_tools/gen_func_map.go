package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	// Specify the target package directory (relative to the current working directory)
	targetPackage := "./src/lib"

	// Parse the package files
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, targetPackage, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Failed to parse package: %v", err)
	}

	// Open the output file for writing
	outputFile, err := os.Create("./src/lib/func_map.go")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Start generating the Go file
	fmt.Fprintln(outputFile, "package lib")
	fmt.Fprintln(outputFile, "")
	fmt.Fprintln(outputFile, "// Auto-generated by gen_func_map.go. DO NOT EDIT.")
	//fmt.Fprintln(outputFile, "import \"yourmodule/lib\"") // Replace `yourmodule` with your module name
	fmt.Fprintln(outputFile, "")
	fmt.Fprintln(outputFile, "var funcMap = map[string]interface{}{")

	// Iterate through each package
	for pkgName, pkg := range pkgs {
		fmt.Printf("Processing package: %s\n", pkgName)

		for filename, file := range pkg.Files {
			fmt.Printf("  Parsing file: %s\n", filename)
			if filename == "src\\lib\\auto_logger_util.go" {
				continue
			}
			// Iterate through declarations in the file
			for _, decl := range file.Decls {
				// Check if the declaration is a function
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					// Only include exported functions (starting with an uppercase letter)
					if funcDecl.Name.IsExported() && funcDecl.Recv == nil {
						fmt.Fprintf(outputFile, "    \"%s\": %s,\n", funcDecl.Name.Name, funcDecl.Name.Name)
					}
				}
			}
		}
	}

	// Close the map and file
	fmt.Fprintln(outputFile, "}")
	fmt.Println("func_map_gen.go has been generated successfully.")
}
