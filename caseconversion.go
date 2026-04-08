package main
import (
	"fmt"
	"strings"
)

func caseConvert(s string) string {
	de := strings.Fields(s)
	for i := 0; i < len(de); i++ {
		if de[i] == "(up)" {
			de[i-1] = strings.ToUpper(de[i-1])
			de = append(de[:i], de[i+1:]...)
			i--
		}
		if de[i] == "(cap)" {
			de[i-1] = strings.Title(de[i-1])
			de = append(de[:i], de[i+1:]...)
			i--
		}
		if de[i] == "(low)" {
			de[i-1] = strings.ToLower(de[i-1])
			de = append(de[:i], de[i+1:]...)
			i--
		}
	}
	return strings.Join(de, " ")
}

func main() {
	fmt.Println(caseConvert("hello how (cap) are you (up) TODAY (low)"))
}