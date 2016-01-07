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
	"sort"
	"strings"
)

var osArches = map[string][]string{
	"linux": {
		"386",
		"amd64",
		"arm",
		"arm64",
		"ppc64",
		"ppc64le",
	},
	"darwin": {
		"386",
		"amd64",
		"arm",
		"arm64",
	},
	"freebsd": {
		"386",
		"amd64",
		"arm",
	},
	"openbsd": {
		"386",
		"amd64",
		"arm",
	},
}

var (
	numFileName      = filepath.Join(os.Getenv("GOROOT"), "src/syscall/zsysnum_%s_%s.go")
	funcOSFileName   = filepath.Join(os.Getenv("GOROOT"), "src/syscall/syscall_%s.go")
	funcArchFileName = filepath.Join(os.Getenv("GOROOT"), "src/syscall/zsyscall_%s_%s.go")
	outFileName      = "syscallpp_%s_%s.go"
)

type sc struct {
	name      string
	number    int
	argsTypes []string
}

type numSort []sc

func (s numSort) Len() int {
	return len(s)
}

func (s numSort) Less(i, j int) bool {
	return s[i].number < s[j].number
}

func (s numSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type generator struct {
	buf      bytes.Buffer
	syscalls []sc
	arch     string
	OS       string
}

func (g *generator) writeHeader() {
	g.buf.WriteString("// generated file; DO NOT EDIT - use go generate in directory with source\n")
	g.buf.WriteString("\n")
	g.buf.WriteString("// +build ")
	g.buf.WriteString(g.arch)
	g.buf.WriteString(",")
	g.buf.WriteString(g.OS)
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

func parseFuncsFromFile(path string) (map[string][]string, error) {
	fs := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fs, path, nil, 0)
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

func parseFuncs(OS, arch string) (map[string][]string, error) {
	funcOSFile := fmt.Sprintf(funcOSFileName, OS)
	funcArchFile := fmt.Sprintf(funcArchFileName, OS, arch)
	common, err := parseFuncsFromFile(funcOSFile)
	if err != nil {
		return nil, err
	}
	archSpec, err := parseFuncsFromFile(funcArchFile)
	if err != nil {
		return nil, err
	}
	// don't care about duplicates
	for f, args := range archSpec {
		common[f] = args
	}
	return common, nil
}

func parseSyscalls(OS, arch string) ([]sc, error) {
	numFile := fmt.Sprintf(numFileName, OS, arch)
	fs := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fs, numFile, nil, 0)
	if err != nil {
		return nil, err
	}
	funcs, err := parseFuncs(OS, arch)
	if err != nil {
		return nil, err
	}
	var syscalls []sc
	if OS == "linux" && (arch == "amd64" || arch == "386" || arch == "ppc64" || arch == "arm64") {
		scs, err := parseLinuxNumbers(arch)
		if err != nil {
			return nil, err
		}
		syscalls = scs
	} else {
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
					s := sc{
						name:   name,
						number: number,
					}
					syscalls = append(syscalls, s)
				}
			}
		}
	}
	for i := range syscalls {
		args, ok := funcs[syscalls[i].name]
		if ok {
			syscalls[i].argsTypes = args
		}
	}
	return syscalls, nil
}

func writePkg(OS, arch string) error {
	outFile := fmt.Sprintf(outFileName, OS, arch)
	scs, err := parseSyscalls(OS, arch)
	if err != nil {
		return err
	}
	sort.Sort(numSort(scs))
	g := &generator{
		syscalls: scs,
		arch:     arch,
		OS:       OS,
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
	for OS, arches := range osArches {
		for _, arch := range arches {
			if err := writePkg(OS, arch); err != nil {
				panic(err)
			}
		}
	}
}
