package template

// ref: https://tech.buysell-technologies.com/entry/adventcalendar2022-12-06

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

type FileInspect struct {
	PackageName   string
	StructInspect []StructInspect
	FuncInspect   []FuncInspect
}

func (f FileInspect) ModelName() string {
	return f.StructInspect[0].ModelName
}

func (f FileInspect) SnakeCaseModelName() string {
	return toSnakeCase(f.StructInspect[0].ModelName) + "s"
}

func (f FileInspect) Columns() []string {
	var columns []string
	for _, s := range f.StructInspect {
		columns = append(columns, s.FieldKeys...)
	}
	return columns
}

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

// scan file and return package name and file info
func Inspect(path string) FileInspect {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var structInspect []StructInspect
	var funcInspect []FuncInspect
	var fileInspect FileInspect

	ast.Inspect(f, func(n ast.Node) bool {
		fileInspect.PackageName = f.Name.Name
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
			funcInspect = append(funcInspect, FuncInspect{FuncName: funcName, Recv: recv})
		}
		return true
	})

	fileInspect.StructInspect = structInspect
	fileInspect.FuncInspect = funcInspect
	return fileInspect
}
