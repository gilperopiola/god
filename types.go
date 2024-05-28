package god

type Int32Slice []int32

func (int32Slice Int32Slice) ToIntSlice() []int {
	intSlice := make([]int, len(int32Slice))
	for i, int32Value := range int32Slice {
		intSlice[i] = int(int32Value)
	}
	return intSlice
}
