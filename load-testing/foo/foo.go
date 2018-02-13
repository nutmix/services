package main

var testVar int
var PublicVar int

func myInit(i int) {
	testVar = i
	PublicVar = i + 1
}

func Get() int {
	return testVar
}
