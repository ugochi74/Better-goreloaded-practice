package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Case(s string) string {
	str := strings.Fields(s)
	cmd := ""
	num := 0
	for i := 1; i < len(str); i++ {
		if strings.HasPrefix(str[i], "(") {
			cmd = str[i]
			if i < len(str)-1 && strings.HasSuffix(str[i+1], ")") {
				temp, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
				if err != nil {
					return s
				}
				num = temp
			} else {
				num = 1
			}
			if num >= i {
				num = i
			}
			for j := i - num; j < i; j++ {
				switch cmd {
				case "(up,", "(up)", "(up),":
					str[j] = strings.ToUpper(str[j])
				case "(low,", "(low)", "(low),":
					str[j] = strings.ToLower(str[j])
				case "(cap,", "(cap)", "(cap),":
					str[j] = strings.ToUpper(string(str[j][0])) + strings.ToLower(string(str[j][1:]))
				case "(hex)", "(hex),":
					temp, err := strconv.ParseInt(str[j], 16, 64)
					if err != nil {
						return s
					}
					str[j] = strconv.FormatInt(temp, 10)
				case "(bin)", "(bin),":
					temp, err := strconv.ParseInt(str[j], 2, 64)
					if err != nil {
						return s
					}
					str[j] = strconv.FormatInt(temp, 10)
				}
			}
		}
		if strings.HasPrefix(str[i], "(") {
			str = append(str[:i], str[i+1:]...)
			i--
		} else if strings.HasSuffix(str[i], ")") {
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}
	return strings.Join(str, " ")
}

func main() {
	fmt.Println(Case("hello how (cap, 2) is you (up) and eveRYBOdy TODAY (low, 2) and today (cap), toady (up), 1E (hex), 1010 (bin)"))
}
