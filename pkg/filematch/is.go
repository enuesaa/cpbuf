package filematch

import (
	"slices"
	"strings"
)

func Is(filename string, searching string) bool {
	if searching == "." || searching == "*" {
		return true
	}
	dirs := splitRmEmpty(filename, "/")
	searchingDirs := splitRmEmpty(searching, "/")

	if len(dirs) == 0 || len(searchingDirs) == 0 {
		return len(dirs) == 0 && len(searchingDirs) == 0
	}

	for _, sd := range searchingDirs {
		dir := dirs[0]
		if !isTextMatch(dir, sd) {
			return false
		}
		if len(dirs) == 1 {
			break
		}
		dirs = dirs[1:]
	}
	return true
}

func isTextMatch(text string, pattern string) bool {
	if !strings.Contains(pattern, "*") {
		return text == pattern
	}

	pattern = strings.ReplaceAll(pattern, "*", "/*/")
	patternSplit := splitRmEmpty(pattern, "/")

	anythingOk := false
	for {
		if len(patternSplit) == 0 {
			break
		}
		ps := patternSplit[0]
		if ps == "*" {
			anythingOk = true
			patternSplit = slices.Delete(patternSplit, 0, 1)
			continue
		}
		var found bool
		text, found = strings.CutPrefix(text, ps)
		if found {
			anythingOk = false
			patternSplit = slices.Delete(patternSplit, 0, 1)
			continue
		}
		if !anythingOk {
			return false
		}
		if len(text) == 1 {
			text = ""
			break
		}
		text = text[1:]
	}

	if anythingOk {
		return len(patternSplit) == 0
	}
	return len(patternSplit) == 0 && len(text) == 0
}
