package main

import (
	"fmt"
	"scopeApp/scope_second"
)

// declare a package level variable for the main
// package named myVar

var myVar = "This is my package level variable"

func main() {
	// variables scope


// declare a block variable for the main function
// called blockVar

var _blockVar = "This is my block variable"


// in the main function, print out the values of myVar,
fmt.Println(myVar)

// blockVar, and PackageVar on one line using the PrintMe
// function in packageone

scopesecond.PrintMe(_blockVar)

}