// Package sliceutil はsliceを取り扱うutilパッケージです.
package sliceutil

// RemoveDuplicateStrings は重複要素を削除した[]stringを返します.
func RemoveDuplicateStrings(args []string) []string {
	return removeDuplicateStrings(args)
}

// RemoveDuplicateInts は重複要素を削除した[]intを返します.
func RemoveDuplicateInts(args []int) []int {
	return removeDuplicateInts(args)
}
