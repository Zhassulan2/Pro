package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type tpl struct {
	Name string
}

func main() {
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fmt.Println(os.Args[1] + file.Name())
		generate(os.Args[1] + file.Name())
	}

	fmt.Println(os.Args[2])
	return

}

func generate(fname string) {

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range node.Decls {
		g, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range g.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			_, ok = currType.Type.(*ast.StructType)
			if ok {
				os.Mkdir(os.Args[2]+"/db", 0777)
				out, _ := os.Create(os.Args[2] + "/db/" + currType.Name.Name + "Resource.go")
				dbTmpl := template.Must(template.New("db.tmpl").Funcs(funcMap).ParseFiles("db.tmpl"))
				dbTmpl.Execute(out, tpl{Name: currType.Name.Name})

				os.Mkdir(os.Args[2]+"/http", 0777)
				outHTTP, _ := os.Create(os.Args[2] + "/http/" + currType.Name.Name + ".go")
				httpTmpl := template.Must(template.New("http.tmpl").Funcs(funcMap).ParseFiles("http.tmpl"))
				httpTmpl.Execute(outHTTP, tpl{Name: currType.Name.Name})

				continue
			}

		}
	}
}
