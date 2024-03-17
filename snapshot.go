package snapshot

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func MatchesInline(t *testing.T, values ...any) {
	t.Helper()

	if len(values) > 2 {
		t.Errorf("snapshot: more than 2 values provided to MatchesInline")
	}

	if len(values) == 2 {
		if !cmp.Equal(values[0], values[1]) {
			t.Errorf("snapshot doesn't match: want %v, got %v", values[1], values[0])
		}

		return
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Error("snapshot: cannot get runtime caller information")
		return
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		t.Errorf("snapshot: failed to parse file: %s", err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return false
		}
		if callExpr, ok := n.(*ast.CallExpr); ok {
			if selector, ok := callExpr.Fun.(*ast.SelectorExpr); ok && selector.Sel.Name == "MatchesInline" {
				if fset.Position(selector.Pos()).Line == line {
					expr, err := parser.ParseExpr(fmt.Sprintf("%#v", values[0]))
					if err != nil {
						panic(err)
					}
					callExpr.Args = append(callExpr.Args, expr)
					var buf bytes.Buffer
					err = format.Node(&buf, fset, f)
					if err != nil {
						panic(err)
					}
					err = os.WriteFile(file, buf.Bytes(), 0)
					if err != nil {
						panic(err)
					}
					return false
				}
			}
		}
		return true
	})
}
