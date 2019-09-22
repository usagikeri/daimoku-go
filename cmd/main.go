package main

import (
	"flag"
	"fmt"

	"github.com/usagikeri/daimoku/util"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	table := util.MakeTable(filename)

	fmt.Println(table)
}
