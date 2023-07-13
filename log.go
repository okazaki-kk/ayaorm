package ayaorm

import (
	"fmt"
	"strings"
)

func InterfaceJoin(values []interface{}, sep string) string {
	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = fmt.Sprintf("%v", v)
	}

	if len(strs) == 0 {
		return ""
	}

	return fmt.Sprintf("[%v]", strings.Join(strs, sep))
}
