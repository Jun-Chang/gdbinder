package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/serenize/snaker"
)

// GenerateSuffix is suffix of the generated file.
const GenerateSuffix = "_gdbinder.go"

var tagRegex = regexp.MustCompile(`([0-9a-zA-Z,_=&\(\)\-]+)(:( )?"([0-9a-zA-Z,_=&\(\)\-]*)")?`)
var typeNm = flag.String("type", "", "a struct type name; must be set")

func main() {
	log.SetFlags(0)
	log.SetPrefix("gdbinder: ")

	flag.Parse()
	if len(*typeNm) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	dir := "."
	pkgName, structName, structType, err := searchStruct(dir)
	if err != nil {
		log.Fatal(err)
	}
	b, err := gen(pkgName, structName, structType)
	if err != nil {
		log.Fatal(err)
	}

	o := snaker.CamelToSnake(structName) + GenerateSuffix
	if err := ioutil.WriteFile(o, b, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	log.Println("generate", o)
}

func searchStruct(dir string) (pkgName string, structName string, st *ast.StructType, err error) {

	var p *build.Package
	p, err = build.Default.ImportDir(dir, 0)
	if err != nil {
		return
	}

	fs := token.NewFileSet()
	for _, f := range p.GoFiles {
		if strings.Contains(f, GenerateSuffix) {
			continue
		}

		var pf *ast.File
		pf, err = parser.ParseFile(fs, dir+"/"+f, nil, 0)
		if err != nil {
			return
		}

		pkgName = pf.Name.String()
		for _, decl := range pf.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			if genDecl.Tok != token.TYPE {
				continue
			}
			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				structName = typeSpec.Name.Name
				if structName != *typeNm {
					continue
				}
				t, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}
				st = t
				return
			}
		}
	}

	err = fmt.Errorf("%s type is not found", *typeNm)
	return
}

func searchTag(exp string) string {
	list := tagRegex.FindAllStringSubmatch(exp, -1)
	for _, v := range list {
		if v[1] != "db" {
			continue
		}
		return v[4]
	}
	return ""
}

func gen(pkgName string, structName string, structType *ast.StructType) ([]byte, error) {
	fields := []string{}
	clmns := []string{}
	for _, field := range structType.Fields.List {
		tag := field.Tag
		if tag == nil {
			continue
		}
		if clmn := searchTag(tag.Value); clmn != "" {
			fields = append(fields, field.Names[0].String())
			clmns = append(clmns, clmn)
		}
	}

	data := map[string]interface{}{}

	data["pkg"] = pkgName
	data["strct"] = structName

	var fieldsBuf bytes.Buffer
	var fieldsMax = len(fields) - 1
	for i, f := range fields {
		fieldsBuf.WriteString("&b.")
		fieldsBuf.WriteString(f)
		if i < fieldsMax {
			fieldsBuf.WriteString(", ")
		}
	}
	data["fields"] = fieldsBuf.String()
	data["fieldMax"] = fieldsMax

	data["clmns"] = strings.Join(clmns, ",")

	t, err := template.New("").Parse(GenerateTemplate)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	if err := t.Execute(w, data); err != nil {
		return nil, err
	}

	if err := w.Flush(); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
