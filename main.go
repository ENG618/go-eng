package main

import (
	"fmt"

	"github.com/eng618/go-eng/algo/search"
	"github.com/eng618/go-eng/algo/sort"
	"github.com/eng618/go-eng/ds/list"
	"github.com/eng618/go-eng/ds/stack"
)

func main() {
	showLinkedList()
	showBinarySearch()
	showStack()
	showMergeSort()
}

func showLinkedList() {
	fmt.Println("\n\nShowing results for LinkedList")
	fmt.Println("Below are example outputs of the list package in action")

	ll := list.NewLinkedList()

	ll.PushBack(20)
	ll.PushBack(30)
	ll.PushBack(40)
	ll.PushFront(50)
	ll.PushFront(70)

	fmt.Println("Length = ", ll.Length())

	ll.Display()

	if err := ll.Delete(40); err != nil {
		fmt.Println("Error deleting 50:", err)
	}

	ll.Reverse()

	fmt.Println("Length = ", ll.Length())
	ll.Display()
	front, _ := ll.PeekFront()
	back, _ := ll.PeekBack()

	fmt.Println("Front = ", front)
	fmt.Println("Back = ", back)
}

func showBinarySearch() {
	fmt.Println("\n\nShowing results for BinarySearch")

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	if v, ok := search.BinaryLoop(xi, 5); ok {
		fmt.Println("Found 5 at index", v)
	}

	if v, ok := search.BinaryLoop(xi, 25); ok {
		fmt.Println("Found 25 at index", v)
	} else {
		fmt.Println("target number no found in slice")
	}

	fmt.Println("\n\nShowing results for BinarySearch")
	fmt.Println("5 is in xi =", search.BinaryRecursion(xi, 5))
	fmt.Println("25 is in xi =", search.BinaryRecursion(xi, 25))
}

func showStack() {
	fmt.Println("\n\nShowing results for Stack")

	s := stack.New()

	s.Push(25)
	s.Push(1)
	s.Push(2)

	if v, err := s.Pop(); err != nil {
		fmt.Println("Pop returned", v)
	}
}

func showMergeSort() {
	fmt.Println("\n\nShowing results for MergeSort")

	xi := []int{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48}

	fmt.Println("Og slice:", xi)
	fmt.Println("After merge sort:", sort.MergeSort(xi))
}
