package helper

import (
	"strconv"
	"strings"
)

var (
	// table for code generator
	table       = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	table1      = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	specialUser = [1]string{"203335595"}
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" {
			oldsize := len(namedQuery)
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))

			if oldsize != len(namedQuery) {
				args = append(args, v)
				i++
			}
		}
	}

	return namedQuery, args
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}
