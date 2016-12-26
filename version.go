package main

import (
	"fmt"
)

var (
	Version  string
	Revision string
)

func printVersion() {
	fmt.Println("kubeps version " + Version + ", build " + Revision)
}
