package event_manager

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

var (
	eventRootDir      = "../event_manager"
	eventRegisterFile = "event_gen.go"
	msgRootDir        = "../event_manager"
	msgRegisterFile   = "msg_gen.go"
)

func scanFuncName(rootDir string, key string) []string {
	var toRegisterEvent []string

	_, err := gfile.ScanDirFunc(rootDir, "*.go", true, func(path string) string {
		if gstr.Contains(path, "test") {
			return ""
		}

		fSet := token.NewFileSet()
		details, err := parser.ParseFile(fSet, path, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for _, object := range details.Scope.Objects {
			if object.Kind != ast.Fun {
				continue
			}

			funcComm := object.Decl.(*ast.FuncDecl)
			if funcComm.Doc == nil {
				continue
			}

			for _, comment := range funcComm.Doc.List {
				text := comment.Text
				if gstr.Contains(text, "register") && gstr.Contains(text, key) {
					toRegisterEvent = append(toRegisterEvent, object.Name)
				}
			}
		}
		return path
	})

	if err != nil {
		panic(err)
	}
	return toRegisterEvent
}

func genEventRegister(funcNames []string, outPath string) {
	err := gfile.Remove(outPath)
	if err != nil {
		fmt.Println("clean export dir fail", err)
		return
	}

	writer := CodeWriter{content: strings.Builder{}}
	writer.appendLine("// this code is auto gen")
	writer.appendEmpty()
	writer.appendLine("package event_manager")
	writer.appendEmpty()

	if len(funcNames) > 0 {
		writer.appendEmpty()
	}

	writer.appendLine("func RegisterEventHandler() {")
	for _, name := range funcNames {
		writer.appendLine("	Register(%s)", name)
	}

	writer.appendLine("}")

	dst, err := gfile.Create(outPath)
	if err != nil {
		panic(err)
	}

	_, err = dst.WriteString(writer.content.String())
	if err != nil {
		panic(fmt.Sprintf("gen %s file fail", "event register"))
	}

	fmt.Println("gen event register success")
}

func genHandlerRegister(funcNames []string, outPath string) {
	if err := gfile.Remove(outPath); err != nil {
		fmt.Println("clean export dir fail", err)
		return
	}

	writer := CodeWriter{content: strings.Builder{}}
	writer.appendLine("// this code is auto gen")
	writer.appendEmpty()
	writer.appendLine("package event_manager")
	writer.appendEmpty()

	writer.appendLine("func RegisterMsgHandler {")
	for _, name := range funcNames {
		writer.appendLine("		Register(%s)", name)
	}

	dst, err := gfile.Create(outPath)
	if err != nil {
		fmt.Printf("create outPath %s fail", outPath)
		return
	}

	_, err = dst.WriteString(writer.content.String())
	if err != nil {
		panic(fmt.Sprintf("gen %s file fial", "event register"))
	}

	fmt.Println("gen msg register success")
}

type CodeWriter struct {
	content strings.Builder
}

func (w *CodeWriter) append(template string, params ...any) {
	w.content.WriteString(fmt.Sprintf(template, params...))
}

func (w *CodeWriter) appendLine(template string, params ...any) {
	w.content.WriteString(fmt.Sprintf(template, params...))
	w.appendEmpty()
}

func (w *CodeWriter) appendEmpty() {
	w.content.WriteString("\r\n")
}
