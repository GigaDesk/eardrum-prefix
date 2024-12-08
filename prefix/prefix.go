package prefix

import (
	"strconv"
	"strings"
)

// Takes a string and an id of a record and adds the id as a prefix to the string
func PrefixWithId(s string, id int) string {
	return strconv.Itoa(id) + "-" + s
}

// Takes a prefixed string and an id of a record and removes the id as a prefix to the string
func DePrefixWithId(s string, id int) string {
	return s[len(strconv.Itoa(id))+1:]
}

/*
Takes a prefixed string and removes a presumed id as a prefix to the string.

Warning: user has to be sure the string was prefixed otherwise te function would yield inaccurate results
*/
func DePrefixWithoutId(s string) string {
	return s[strings.Index(s, "-")+1:]
}
