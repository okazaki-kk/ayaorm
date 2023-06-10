package template

// ref: https://tech.buysell-technologies.com/entry/adventcalendar2022-12-06

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/okazaki-kk/ayaorm"
)

type FileInspect struct {
	PackageName         string
	StructInspect       []StructInspect
	RelationFuncInspect []RelationFuncInspect
	ValidateFuncInspect []ValidateFuncInspect
	CustomRecv          []string
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

type RelationFuncInspect struct {
	FuncName string
	Recv     string
	Args     []string
	HasMany  bool
	BelongTo bool
	HasOne bool
}

type ValidateFuncInspect struct {
	FuncName             string
	Recv                 string
	Args                 []string
	ValidatePresence     bool
	ValidateLength       bool
	ValidateNumericality bool
}

func (f RelationFuncInspect) HasManyModel() string {
	hasManyModels := strings.TrimPrefix(f.FuncName, "hasMany")
	return hasManyModels[:len(hasManyModels)-1]
}

func (f RelationFuncInspect) BelongsToModel() string {
	return strings.TrimPrefix(f.FuncName, "belongsTo")
}

func (f RelationFuncInspect) HasOneModel() string {
	return strings.TrimPrefix(f.FuncName, "hasOne")
}

func (f ValidateFuncInspect) ValidatePresenceField() string {
	return strings.TrimPrefix(f.FuncName, "validatesPresenceOf")
}

func (f ValidateFuncInspect) ValidateLengthField() string {
	return strings.TrimPrefix(f.FuncName, "validateLengthOf")
}

func (f ValidateFuncInspect) ValidateNumericalityField() string {
	return strings.TrimPrefix(f.FuncName, "validateNumericalityOf")
}

// scan file and return package name and file info
func Inspect(path string) (FileInspect, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return FileInspect{}, err
	}

	var fileInspect FileInspect
	var structInspect []StructInspect
	var relationFuncInspect []RelationFuncInspect
	var validateFuncInspect []ValidateFuncInspect

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

			if strings.HasPrefix("validateCustomRule", funcName) {
				fileInspect.CustomRecv = append(fileInspect.CustomRecv, recv)
				return true
			}

			hasMany := strings.HasPrefix(funcName, "hasMany")
			belongsTo := strings.HasPrefix(funcName, "belongsTo")
			hasOne := strings.HasPrefix(funcName, "hasOne")
			if hasMany || belongsTo || hasOne {
				relationFuncInspect = append(relationFuncInspect, RelationFuncInspect{FuncName: funcName, Recv: recv, HasMany: hasMany, BelongTo: belongsTo, HasOne: hasOne})
				return true
			}

			validatePresence := strings.HasPrefix(funcName, "validatesPresenceOf")
			validateLength := strings.HasPrefix(funcName, "validateLengthOf")
			validateNumericality := strings.HasPrefix(funcName, "validateNumericalityOf")
			if validatePresence || validateLength || validateNumericality {
				validateFuncInspect = append(validateFuncInspect, ValidateFuncInspect{FuncName: funcName, Recv: recv, ValidatePresence: validatePresence, ValidateLength: validateLength, ValidateNumericality: validateNumericality})
				return true
			}

		}
		return true
	})

	fileInspect.StructInspect = structInspect
	fileInspect.ValidateFuncInspect = validateFuncInspect
	fileInspect.RelationFuncInspect = relationFuncInspect

	return fileInspect, nil
}
