//variable is the name of memeory allocation

package main

import "fmt"

func main() {
	var t = 123                  //type willbe int
	var u = "circle"             //string type
	var v = 5.6                  //float number
	var z = sample{name: "test"} //type inferred will be main.sample

	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

type sample struct {
	name string
}

// cariable declaration are of 2 types
// local and globle variable declartation
