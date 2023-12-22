package scopesecond

import "fmt"

// declare a package level variable in the scope_second
// package named PackageVar

var PackageVar = "This is a package variable in the scope_second package"

// create an exported function in packageone called PrintMe

func PrintMe(printObject1 string){
	fmt.Println(printObject1, PackageVar)
}


