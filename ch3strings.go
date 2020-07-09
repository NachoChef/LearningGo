package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(comma("-12345123154567"))
}

// handles signed/unsigned int and floats
func comma(s string) string {
	n := len(s)
	const size = 3
	sign := 0
	var buf bytes.Buffer
	if s[0] == '-' {
		sign = 1
	}
	if strings.ContainsRune(s, '.') {
		n = strings.IndexRune(s, '.')
	}
	r := (n - sign) % size
	buf.WriteString(s[:r+sign])
	for i := r + sign; i < n-sign; i += size {
		if sign <= 0 || len(buf.String()) != 1 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i : i+size])
	}
	if n != len(s) {
		buf.WriteString(s[n:])
	}
	return buf.String()
}

// subtype to define sorting methods
type sortBy []rune

func (a sortBy) Len() int           { return len(a) }
func (a sortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortBy) Less(i, j int) bool { return a[i] < a[j] }

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var sl1 sortBy = strToSlice(s1)
	var sl2 sortBy = strToSlice(s2)

	sort.Sort(sl1)
	sort.Sort(sl2)

	return string(sl1) == string(sl2)
}

func strToSlice(s string) []rune {
	rSlice := make([]rune, len(s))
	for i, r := range s {
		rSlice[i] = r
	}
	return rSlice
}
