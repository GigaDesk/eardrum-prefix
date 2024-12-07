package prefix

import "strconv"

// Takes a string and an id of a record and adds the id as a prefix to the string
func PrefixWithId(s string, id int) string{
return strconv.Itoa(id) + "-" + s
}

// Takes a prefixed string and an id of a record and removes the id as a prefix to the string
func DePrefixWithId(s string, id int) string{
return s[len(strconv.Itoa(id)) + 1: ]
}