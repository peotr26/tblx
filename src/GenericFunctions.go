package tblx

// Returns the maximum value in an array of integer.
func max(x []int) int {
	maxi := x[0]
	for _, e := range x {
		if e > maxi {
			maxi = e
		}
	}
	return maxi
}
