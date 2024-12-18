package main

import (
	"fmt"

	"github.com/Horsen121/TBD/RPBD/concurrency/scan/scan"
)

func main() {
	opened := scan.Scan("127.0.0.1")

	fmt.Println("Opened ports:")
	for i := 0; i < len(opened); i++ {
		fmt.Println(opened[i])
	}
}
