package main

import (
	pdk "github.com/extism/go-pdk"
)

//export cat
func cat() int32 {
	input := pdk.InputString()
	pdk.OutputString(input)
	return 0
}

func main() {}
