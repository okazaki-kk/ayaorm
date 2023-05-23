package ayaorm

// ref: https://tech.buysell-technologies.com/entry/adventcalendar2022-12-06

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type StructInspect struct {
	ModelName   string
	FieldKeys   []string
	FieldValues []string
}

type FuncInspect struct {
	FuncName string
	Recv     string
	Args     []string
}

func Inspect(path string) ([]StructInspect, FuncInspect) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var structInspect []StructInspect
	var funcInspect FuncInspect

	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.TypeSpec:
			s, _ := n.(*ast.TypeSpec)
			v, ok := s.Type.(*ast.StructType)
			if !ok {
				return false
			}
			var fi StructInspect
			fi.ModelName = s.Name.Name
			fi.FieldKeys = append(fi.FieldKeys, "Id")
			fi.FieldValues = append(fi.FieldValues, "int")
			// 構造体かつその名前が対象のモデルの場合
			for _, l := range v.Fields.List {
				if len(l.Names) <= 0 {
					continue
				}
				switch l.Type.(type) {
				case *ast.Ident: // intやstringのようなプリミティブな型の場合
					t, _ := l.Type.(*ast.Ident)
					fi.FieldKeys = append(fi.FieldKeys, l.Names[0].Name)
					fi.FieldValues = append(fi.FieldValues, t.Name)
				case *ast.SelectorExpr: // time.Timeやnull.Stringのような型
					t, _ := l.Type.(*ast.SelectorExpr)
					x, _ := t.X.(*ast.Ident)
					name := x.Name + "." + t.Sel.Name
					fi.FieldKeys = append(fi.FieldKeys, l.Names[0].Name)
					fi.FieldValues = append(fi.FieldValues, name)
				case *ast.StarExpr:
					t, _ := l.Type.(*ast.StarExpr)
					switch t.X.(type) {
					case *ast.Ident: // *intのようなプリミティブのポインタ型
					// 処理内容は上のIdentとほぼ同じなので省略
					case *ast.SelectorExpr: // *time.Timeのようなポインタ型
						// 処理内容は上のSelectorExprとほぼ同じなので省略
					}
				}
			}
			fi.FieldKeys = append(fi.FieldKeys, "CreatedAt")
			fi.FieldValues = append(fi.FieldValues, "time.Time")
			fi.FieldKeys = append(fi.FieldKeys, "UpdatedAt")
			fi.FieldValues = append(fi.FieldValues, "time.Time")
			structInspect = append(structInspect, fi)

		case *ast.FuncDecl:
			funcName := n.(*ast.FuncDecl).Name.Name
			recv := n.(*ast.FuncDecl).Recv.List[0].Type.(*ast.Ident).Name
			funcInspect.FuncName = funcName
			funcInspect.Recv = recv
		}
		return true
	})
	return structInspect, funcInspect
}
