package main

import (
	"testing"
	"testing/quick"
)

func TestMergeSort(t *testing.T) {
	f := func(slice []int) bool {
		sorted := MergeSort(slice)
		for i := 1; i < len(slice); i++ {
			if sorted[i-1] > sorted[i] {
				return false
			}
		}
		return true
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
