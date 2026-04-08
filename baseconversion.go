package main
import (
	"fmt"
	"strconv"
	"strings"
	
)
func remove(s string) string {
	de := strings.Fields(s)
	for i := 0; i < len(de); i++ {
		if de[i] == "(hex)" {
			res, err := strconv.ParseInt(de[i-1], 16, 64)
			if err == nil {
				de[i-1] = strconv.FormatInt(res, 10)
				de = append(de[:i], de[i+1:]...)
				i--
			}
		}
		if de[i] == "(bin)" {
			res, err := strconv.ParseInt(de[i-1], 2, 64)
			if err == nil {
				de[i-1] = strconv.FormatInt(res, 10)
				de = append(de[:i], de[i+1:]...)
				i--
			}
		}

	}
	return strings.Join(de, " ")
}

func main() {
	fmt.Println(remove("It has been 1E (hex) years and 10 (bin) months"))
}