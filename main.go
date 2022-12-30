package main

import (
	"fmt"
	"strings"
)

func main() {
	pathSplits := strings.Split("/access/assign-role", "/")

	fmt.Println(pathSplits)

}
