package objects

import (
	"fmt"
	"strings"
)

func asString(isRemoval bool, id uint64, props map[string]string) string {
	r := ""
	if isRemoval {
		r = "-"
	}
	propStrs := make([]string, 0, len(props))
	for k, v := range props {
		propStrs = append(propStrs, fmt.Sprintf("%s=%s", k, v))
	}
	return fmt.Sprintf("%s%x,%s", r, id, strings.Join(propStrs, ","))
}
