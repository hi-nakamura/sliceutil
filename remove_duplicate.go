package sliceutil

func removeDuplicate(args interfaceSlice) interfaceSlice {
	results := args.newSlice()
	encountered := map[interface{}]bool{}
	for i := 0; i < args.len(); i++ {
		if !encountered[args.value(i)] {
			encountered[args.value(i)] = true
			results = results.append(args.value(i))
		}
	}
	return results
}
