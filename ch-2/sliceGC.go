package main

import "runtime"

type data struct {
	i, j int
}

func main() {
	var N = 40000000
	var structure []data
	for i := 0; i < N; i++ {
		value := int(i)
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	_ = structure[0]
	/*
		preventing the garbage collector from garbage collecting the structure variable too early,
		as it is not referenced or used outside of the for loop.
	*/
}
