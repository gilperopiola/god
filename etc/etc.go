package etc

// We wanna keep the god pkg as clean and unbloated as possible,
// that's why this pkg exists. The code here is as godly as the
// god code, just not used as often.

/* -~-~-~-~ Types & Conversions -~-~-~-~ */

type I32Slice []int32

func (i32Slice I32Slice) ToIntSlice() []int {
	iSlice := make([]int, len(i32Slice))
	for index, i32Val := range i32Slice {
		iSlice[index] = int(i32Val)
	}
	return iSlice
}
