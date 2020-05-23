package maff

/*
	For some reason, all the default math is for floats. This is for ints.
*/

// Abs -- Get the absolute value of an integer
func Abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

// ShareSign -- See if two ints share positivity/negativity
func ShareSign(a int, b int) bool {
	if a < 0 {
		return b < 0
	}
	return b >= 0
}
