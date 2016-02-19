package sliceutil

type stringSlice []string

func (p stringSlice) len() int                            { return len(p) }
func (p stringSlice) value(index int) interface{}         { return p[index] }
func (p stringSlice) newSlice() interfaceSlice            { return stringSlice(make([]string, 0, len(p))) }
func (p stringSlice) append(v interface{}) interfaceSlice { return append(p, v.(string)) }

func removeDuplicateStrings(args []string) []string {
	result := removeDuplicate(stringSlice(args))
	return result.(stringSlice)
}
