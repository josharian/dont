// TODO
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"strings"

	"golang.org/x/tools/go/loader"
)

type check struct {
	tfun *types.Func
	decl *ast.FuncDecl
}

func (c check) Comment() string { return c.decl.Doc.Text() }

func isDontFunc(fn *ast.FuncDecl) bool {
	if !strings.HasPrefix(fn.Name.Name, "Dont") {
		return false
	}
	if fn.Recv != nil {
		// TODO: allow methods? log an error here?
		return false
	}
	if fn.Type.Results != nil {
		// TODO: allow results? log an error here?
		return false
	}
	return true
}

func main() {
	log.SetPrefix("dont: ")
	log.SetFlags(0)

	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("TODO: usage")
	}

	conf := loader.Config{
		Fset:       token.NewFileSet(),
		ParserMode: parser.ParseComments,
	}

	if _, err := conf.FromArgs(args, true); err != nil {
		log.Fatal(err)
	}

	// Load, parse and type-check the whole program.
	iprog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Find all Dont* functions.
	var checks []check

	pkgs := iprog.InitialPackages()
	for _, pkg := range pkgs {
		scope := pkg.Pkg.Scope()
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}
				if ok := isDontFunc(fn); !ok {
					continue
				}
				name := fn.Name.Name
				obj := scope.Lookup(name)
				if obj == nil {
					log.Fatalf("could not find %q.%s", pkg.Pkg.Path(), name)
				}
				c := check{decl: fn, tfun: obj.(*types.Func)}
				checks = append(checks, c)
			}
		}
	}

	const dumpChecks = false
	if dumpChecks {
		fmt.Println("Found checks:")
		for _, c := range checks {
			fmt.Println("\t", c.tfun.FullName())
		}
	}

	var failed bool

	// Analyze all packages looking for matches.
	for _, pkg := range pkgs {
		// scope := pkg.Pkg.Scope()
		for _, f := range pkg.Files {
			for _, decl := range f.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok {
					// TODO: this misses closures. Fix.
					continue
				}
				if ok := isDontFunc(fn); ok {
					// don't check the Dont funcs themselves :)
					continue
				}
				// TODO: check all donts at once
				// TODO: In verbose mode, dump what function we're checking and
				for _, c := range checks {
					if match(c.decl.Body.List, fn.Body.List) {
						fmt.Println("MATCH", c.tfun.FullName())
						failed = true
					}
				}
			}
		}
	}

	if failed {
		os.Exit(1)
	}

	// 	for _, file := range pkg.Files {
	// 		n := xform.Transform(&pkg.Info, pkg.Pkg, file)
	// 		if n == 0 {
	// 			continue
	// 		}
	// 		filename := iprog.Fset.File(file.Pos()).Name()
}

// match reports whether c matches body.
func match(tpl, body []ast.Stmt) bool {
	fmt.Printf("check type %v %v\n", len(tpl), len(body))
	return false
}

// TODO: generate and return correspondences between variables as well
// Probably we'll need to start by trying to bind all free variables in the template
// (that is all function arguments), and then switch to pure matching mode.
func matchstmt(tpl, stmt ast.Stmt) bool {
	// AssignStmt
	// BlockStmt
	// BranchStmt
	// DeclStmt
	// DeferStmt
	// EmptyStmt
	// ExprStmt
	// ForStmt
	// GoStmt
	// IfStmt
	// IncDecStmt
	// LabeledStmt
	// RangeStmt
	// ReturnStmt
	// SelectStmt
	// SendStmt
	// SwitchStmt
	// TypeSwitchStmt

}
