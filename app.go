package main

import (
	"fmt"

	"github.com/ocantog/api/src/support"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(support.Reverse("!ovatsuG \n"))

	fmt.Println(cmp.Diff("Gustavo", "Gus"))
}
