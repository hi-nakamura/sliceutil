package sliceutil

func removeDuplicate(args interfaceSlice) interfaceSlice {
	results := args.newSlice()
	for i := 0; i < args.len(); i++ {
		dup := false
		for j := 0; j < results.len(); j++ {
			if args.value(i) == results.value(j) {
				dup = true
				break
			}
		}
		if !dup {
			results = results.append(args.value(i))
		}
	}
	return results
}
