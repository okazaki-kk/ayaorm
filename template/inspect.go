package template

// ref: https://tech.buysell-technologies.com/entry/adventcalendar2022-12-06

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/okazaki-kk/ayaorm"
)

type FileInspect struct {
	PackageName   string
	StructInspect []StructInspect
	FuncInspect   []FuncInspect
}

func (f FileInspect) ModelName() string {
	return f.StructInspect[0].ModelName
}

type StructInspect struct {
	PackageName string
	ModelName   string
	FieldKeys   []string
	FieldValues []string
}

func (s StructInspect) Columns() []string {
	var columns []string
	columns = append(columns, s.FieldKeys...)
	return columns
}

func (s StructInspect) SnakeCaseModelName() string {
	return ayaorm.ToSnakeCase(s.ModelName) + "s"
}

type FuncInspect struct {
	FuncName         string
	Recv             string
	Args             []string
	HasMany          bool
	BelongTo         bool
	ValidatePresence bool
	ValidateLength   bool
}

func (f FuncInspect) HasManyModel() string {
	hasManyModels := strings.TrimPrefix(f.FuncName, "hasMany")
	return hasManyModels[:len(hasManyModels)-1]
}

func (f FuncInspect) BelongsToModel() string {
	return strings.TrimPrefix(f.FuncName, "belongsTo")
}

func (f FuncInspect) ValidatePresenceField() string {
	return strings.TrimPrefix(f.FuncName, "validatesPresenceOf")
}

func (f FuncInspect) ValidateLengthField() string {
	return strings.TrimPrefix(f.FuncName, "validateLengthOf")
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
			hasMany := strings.HasPrefix(funcName, "hasMany")
			belongsTo := strings.HasPrefix(funcName, "belongsTo")
			validatePresence := strings.HasPrefix(funcName, "validatesPresenceOf")
			validateLength := strings.HasPrefix(funcName, "validateLengthOf")

			funcInspect = append(funcInspect, FuncInspect{FuncName: funcName, Recv: recv, HasMany: hasMany, BelongTo: belongsTo, ValidatePresence: validatePresence, ValidateLength: validateLength})
		}
		return true
	})

	fileInspect.StructInspect = structInspect
	fileInspect.FuncInspect = funcInspect
	return fileInspect
}
