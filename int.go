package sliceutil

type intSlice []int

func (p intSlice) len() int                            { return len(p) }
func (p intSlice) value(index int) interface{}         { return p[index] }
func (p intSlice) newSlice() interfaceSlice            { return intSlice(make([]int, 0, len(p))) }
func (p intSlice) append(v interface{}) interfaceSlice { return append(p, v.(int)) }

func removeDuplicateInts(args []int) []int {
	result := removeDuplicate(intSlice(args))
	return result.(intSlice)
}
