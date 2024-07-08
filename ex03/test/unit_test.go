package test

import (
	"piscine"
	"reflect"
	"testing"
)

func expect(arr, expect []string, t *testing.T) {
	piscine.QuickSort(arr, 0, len(arr)-1)
	if !reflect.DeepEqual(arr, expect) {
		t.Errorf("expected %s, but got %s", expect, arr)
	}
}

func TestQuickSort(t *testing.T) {
	expect([]string{"a", "b", "c", "2", "1", "0"}, []string{"0", "1", "2", "a", "b", "c"}, t)
	expect([]string{"abc", "aab", "acc"}, []string{"aab", "abc", "acc"}, t)
	expect([]string{"0", "a", ""}, []string{"", "0", "a"}, t)
	expect([]string{"0", "1", "2"}, []string{"0", "1", "2"}, t)
	expect([]string{"0"}, []string{"0"}, t)
}
