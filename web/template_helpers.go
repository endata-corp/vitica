package web

import (
	"html/template"
	"reflect"
	"strings"
)

func GetTemplateHelpers() []template.FuncMap {
	return []template.FuncMap{
		template.FuncMap{"title": title},
		template.FuncMap{"upper": upper},
		template.FuncMap{"eq": eq},
	}
}

func title(str string) string {
	return strings.Title(str)
}

func upper(str string) string {
	return strings.ToUpper(str)
}

func eq(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}
	x := args[0]
	switch x := x.(type) {
	case string, int, int64, byte, float32, float64:
		for _, y := range args[1:] {
			if x == y {
				return true
			}
		}
		return false
	}

	for _, y := range args[1:] {
		if reflect.DeepEqual(x, y) {
			return true
		}
	}
	return false
}
