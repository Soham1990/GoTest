package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	array := []int{5, 6, 7, 2, 1, 0}
	expected := []int{0, 1, 2, 5, 6, 7}
	actual := Sort(array, 0, len(array)-1)
	for i := range actual {
		if actual[i] != expected[i] {
			t.Errorf("Test fail! want: '%d', got: '%d'", expected, actual)
		}
	}
	fmt.Println("Success")
}
