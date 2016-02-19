package sliceutil

type interfaceSlice interface {
	len() int
	value(index int) interface{}
	newSlice() interfaceSlice
	append(v interface{}) interfaceSlice
}
