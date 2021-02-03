package cp5

import (
	"fmt"
	"strings"
)

func add1(r rune) rune {
	fmt.Println(r, "------")
	return r + 1
}

func AddUse() {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VNS"))
}
