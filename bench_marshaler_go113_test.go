//go:build go1.13
// +build go1.13

package reflect_test

import "unsafe"

//go:linkname MapIterValue reflect.mapiterelem
func MapIterValue(it *HashIter) unsafe.Pointer
