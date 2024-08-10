package filematch

import (
	"slices"
	"strings"
)

func splitRmEmpty(s string, sep string) []string {
	ret := strings.Split(s, sep)
	ret = slices.DeleteFunc(ret, func(v string) bool {
		return v == ""
	})

	return ret
}
