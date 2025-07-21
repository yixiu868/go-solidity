package pointer

func OperSlice(slicePtr *[]int) {
	slice := *slicePtr
	for i := range slice {
		slice[i] *= 2
	}
}
