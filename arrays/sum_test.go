package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	got := Sum(numbers)
	want := 6
	assert(t, got, want)
}

func TestSumAll(t *testing.T) {
	set1 := []int{1, 2, 3}
	set2 := []int{1, 2, 3, 4, 5}
	got := SumAll(set1, set2)
	want := []int{6, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
