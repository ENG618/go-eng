package sort_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eng618/go-eng/algo/sort"
)

func ExampleMergeSort() {
	data := []int{5, 4, 3, 2, 1}
	sorted := sort.MergeSort(data)
	fmt.Println(sorted)
	// Output: [1 2 3 4 5]
}

func ExampleMerge() {
	left := []int{1, 3, 5}
	right := []int{2, 4, 6}
	merged := sort.Merge(left, right)
	fmt.Println(merged)
	// Output: [1 2 3 4 5 6]
}

func TestMergeSort(t *testing.T) {
	t.Parallel()

	type args struct {
		d []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Base Example", args: args{d: []int{5, 4, 3, 2, 1}}, want: []int{1, 2, 3, 4, 5}},
		{name: "Empty []int", args: args{[]int{}}, want: []int{}},
		{name: "Single item in []int", args: args{[]int{5}}, want: []int{5}},
		{
			name: "Complex Example",
			args: args{[]int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}},
			want: []int{2, 3, 4, 5, 15, 19, 26, 27, 36, 38, 44, 46, 47, 48, 50},
		},
		{
			name: "Uneven heavy",
			args: args{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 1, 2, 3}},
			want: []int{1, 2, 3, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		},
		{
			name: "Uneven light",
			args: args{[]int{7, 8, 9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sort.MergeSort(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []int
		want []int
	}{
		{name: "All same elements", args: []int{1, 1, 1, 1, 1}, want: []int{1, 1, 1, 1, 1}},
		{name: "Negative numbers", args: []int{-3, -1, -2, -5, -4}, want: []int{-5, -4, -3, -2, -1}},
		{name: "Mixed positive and negative", args: []int{3, -1, 2, -5, 4}, want: []int{-5, -1, 2, 3, 4}},
		{name: "Already sorted", args: []int{1, 2, 3, 4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "Reverse sorted", args: []int{5, 4, 3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sort.MergeSort(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
		r []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "Basic1", args: args{l: []int{1}, r: []int{5}}, want: []int{1, 5}},
		{name: "Basic2", args: args{l: []int{5}, r: []int{1}}, want: []int{1, 5}},
		{name: "Basic3", args: args{l: []int{1, 2, 3, 4}, r: []int{5, 6, 7, 8, 9}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "Basic4", args: args{l: []int{1, 4, 7, 9}, r: []int{2, 3, 5, 6, 8}}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "Empty right", args: args{l: []int{1, 2, 3, 4, 5}, r: []int{}}, want: []int{1, 2, 3, 4, 5}},
		{name: "Empty left", args: args{l: []int{}, r: []int{1, 2, 3, 4, 5}}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := sort.Merge(tt.args.l, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeEdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		l    []int
		r    []int
		want []int
	}{
		{name: "All same elements", l: []int{1, 1, 1}, r: []int{1, 1, 1}, want: []int{1, 1, 1, 1, 1, 1}},
		{name: "Negative numbers", l: []int{-3, -1, -2}, r: []int{-5, -4}, want: []int{-5, -4, -3, -2, -1}},
		{name: "Mixed positive and negative", l: []int{3, -1}, r: []int{2, -5, 4}, want: []int{-5, -1, 2, 3, 4}},
		{name: "Already sorted", l: []int{1, 2, 3}, r: []int{4, 5}, want: []int{1, 2, 3, 4, 5}},
		{name: "Reverse sorted", l: []int{5, 4}, r: []int{3, 2, 1}, want: []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Sort the input slices before merging
			if got := sort.Merge(tt.l, tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}
		sort.MergeSort(data)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		left := []int{1, 3, 5, 7, 9}
		right := []int{2, 4, 6, 8, 10}
		sort.Merge(left, right)
	}
}
