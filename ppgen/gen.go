package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var arches = []string{
	"386",
	"amd64",
	"arm",
	"arm64",
	"ppc64",
	"ppc64le",
}

var (
	numFileName  = filepath.Join(os.Getenv("GOROOT"), "src/syscall/zsysnum_linux_%s.go")
	funcFileName = filepath.Join(os.Getenv("GOROOT"), "src/syscall/zsyscall_linux_%s.go")
	outFileName  = "syscallpp_linux_%s.go"
)

type sc struct {
	name      string
	number    int
	argsTypes []string
}

type generator struct {
	buf      bytes.Buffer
	syscalls []sc
	arch     string
}

func (g *generator) writeHeader() {
	g.buf.WriteString("// generated file; DO NOT EDIT - use go generate in directory with source\n")
	g.buf.WriteString("\n")
	g.buf.WriteString("// +build ")
	g.buf.WriteString(g.arch)
	g.buf.WriteString(",linux")
	g.buf.WriteString("\n\n")
	g.buf.WriteString("package syscallpp")
}

func (g *generator) writeGetNameFunc() {
	g.buf.WriteString("\n")
	g.buf.WriteString("func GetName(num int) string {\n")
	g.buf.WriteString("switch num {\n")
	for _, s := range g.syscalls {
		fmt.Fprintf(&g.buf, "case %d:\n", s.number)
		fmt.Fprintf(&g.buf, "return \"%s\"\n", s.name)
	}
	g.buf.WriteString("}\n")
	g.buf.WriteString("return \"unknown\"\n")
	g.buf.WriteString("}\n")
}

func (g *generator) writeGetNumFunc() {
	g.buf.WriteString("\n")
	g.buf.WriteString("func GetNum(name string) int {\n")
	g.buf.WriteString("switch name {\n")
	for _, s := range g.syscalls {
		fmt.Fprintf(&g.buf, "case \"%s\":\n", s.name)
		fmt.Fprintf(&g.buf, "return %d\n", s.number)
	}
	g.buf.WriteString("}\n")
	g.buf.WriteString("return -1\n")
	g.buf.WriteString("}\n")
}

func (g *generator) writeGetArgsTypesFunc() {
	g.buf.WriteString("\n")
	g.buf.WriteString("func GetArgsTypes(name string) []ArgType {\n")
	g.buf.WriteString("switch name {\n")
	for _, s := range g.syscalls {
		fmt.Fprintf(&g.buf, "case \"%s\":\n", s.name)
		strTypes := fmt.Sprintf("%#v", s.argsTypes)
		strTypes = strings.Replace(strTypes, `"`, "", -1)
		strTypes = strings.Replace(strTypes, "string", "ArgType", -1)
		fmt.Fprintf(&g.buf, "return %s\n", strTypes)
	}
	g.buf.WriteString("}\n")
	g.buf.WriteString("return nil\n")
	g.buf.WriteString("}\n")
}

func parseFuncs(arch string) (map[string][]string, error) {
	funcFile := fmt.Sprintf(funcFileName, arch)
	fs := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fs, funcFile, nil, 0)
	if err != nil {
		return nil, err
	}
	funcs := make(map[string][]string)
	for _, decl := range parsedFile.Decls {
		decl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		name := strings.ToLower(decl.Name.Name)
		var ps []string
		for _, p := range decl.Type.Params.List {
			t := fmt.Sprintf("%s", p.Type)
			if strings.HasPrefix(t, "uint") {
				t = "int"
			}
			if t != "int" && t != "string" {
				t = "ptr"
			}
			switch t {
			case "int":
				t = "ARG_INT"
			case "string":
				t = "ARG_STR"
			case "ptr":
				t = "ARG_PTR"
			}
			ps = append(ps, t)
		}
		funcs[name] = ps
	}
	return funcs, nil
}

func parseSyscalls(arch string) ([]sc, error) {
	numFile := fmt.Sprintf(numFileName, arch)
	fs := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fs, numFile, nil, 0)
	if err != nil {
		return nil, err
	}
	funcs, err := parseFuncs(arch)
	if err != nil {
		return nil, err
	}
	var syscalls []sc
	for _, decl := range parsedFile.Decls {
		decl, ok := decl.(*ast.GenDecl)
		if !ok || decl.Tok != token.CONST {
			continue
		}
		for _, spec := range decl.Specs {
			vspec := spec.(*ast.ValueSpec)
			name := vspec.Names[0].Name
			if strings.HasPrefix(name, "SYS_") {
				number, ok := vspec.Names[0].Obj.Data.(int)
				if !ok {
					return nil, fmt.Errorf("Unexpected type of constant %T", number)
				}
				name := strings.ToLower(name[4:])
				args, ok := funcs[name]
				if !ok {
					continue
				}
				syscalls = append(syscalls, sc{
					name:      name,
					number:    number,
					argsTypes: args,
				})
			}
		}
	}
	return syscalls, nil
}

func writePkg(arch string) error {
	outFile := fmt.Sprintf(outFileName, arch)
	scs, err := parseSyscalls(arch)
	if err != nil {
		return err
	}
	g := &generator{
		syscalls: scs,
		arch:     arch,
	}
	g.writeHeader()
	g.writeGetNameFunc()
	g.writeGetNumFunc()
	g.writeGetArgsTypesFunc()
	src, err := format.Source(g.buf.Bytes())
	if err != nil {
		fmt.Println("generated invalid Go code")
		fmt.Println(g.buf.String())
		return err
	}
	var perm os.FileMode
	fi, err := os.Stat(outFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		perm = 0644
	} else {
		perm = fi.Mode().Perm()
	}
	if err := ioutil.WriteFile(outFile, src, perm); err != nil {
		return err
	}
	return nil
}

func main() {
	for _, arch := range arches {
		if err := writePkg(arch); err != nil {
			panic(err)
		}
	}
}
