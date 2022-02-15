package main

import (
	"fmt"
	"start/profile"
)

var isGraduated bool

func main() {
	name := "Ahmed"
	var age int
	age = 29
	fmt.Println(name)
	fmt.Println(age)

	profile.IsAged = true
	fmt.Println(profile.IsAged)
}
