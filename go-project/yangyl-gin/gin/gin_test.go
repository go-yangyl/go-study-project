package gin

import (
	"fmt"
	"testing"
)

func TestParsePattern(t *testing.T) {
	str1 := "a/b/c"
	fmt.Println(ParsePattern(str1))

	str2 := "a/b/*c/d"
	fmt.Println(ParsePattern(str2))

	str3 := "a/b/:c/d"
	fmt.Println(ParsePattern(str3))
}

func TestNode_Insert(t *testing.T) {
	var nod = new(node)
	str1 := "a/:b"

	nod.Insert(str1, ParsePattern(str1), 0)
	fmt.Println(nod)
	fmt.Println(nod.Search(ParsePattern(str1), 0).pattern)

}
