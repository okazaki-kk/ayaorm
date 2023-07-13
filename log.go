package ayaorm

import (
	"fmt"
	"reflect"
	"strings"
)

func InterfaceJoin(values []interface{}, sep string) string {
	strs := make([]string, len(values))
	for i, v := range values {
		if reflect.TypeOf(v).Kind() == reflect.String {
			strs[i] = fmt.Sprintf("'%s'", v)
		} else {
			strs[i] = fmt.Sprintf("%v", v)
		}
	}

	if len(strs) == 0 {
		return ""
	}

	return fmt.Sprintf("[%v]", strings.Join(strs, sep))
}
