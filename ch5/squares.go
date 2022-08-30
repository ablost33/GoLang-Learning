package ch5

// Experimentation w/ anonymous functions
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
