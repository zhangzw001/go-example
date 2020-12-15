package main

import (
	"testing"
)


//func TestRemoveSliceIndex(b *testing.T) {
//	s1 := []int{1,2,3,4,5,6,7,8,9}
//	fmt.Println(RemoveSliceIndex(s1,6))
//}
//
//func TestRemoveSliceIndex2(b *testing.T) {
//	s1 := []int{1,2,3,4,5,6,7,8,9}
//	fmt.Println(RemoveSliceIndex2(s1,6))
//}

func BenchmarkRemoveSliceIndex(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		s1 := []int{1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9}
		RemoveSliceIndex(s1,6)
	}
}

func BenchmarkRemoveSliceIndex2(b *testing.B) {
	for i:=0 ; i < b.N ; i++ {
		s1 := []int{1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9,1,2,3,4,5,6,7,8,9}
		RemoveSliceIndex2(s1,6)

	}
}
