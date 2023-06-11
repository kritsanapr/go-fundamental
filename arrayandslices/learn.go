package arrayandslices

import "fmt"

func Learn() {
	fmt.Println("Learn Array and Slices")
	// Array
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println("Array 1 ", arr)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	for i := 0; i < len(arr2); i++ {
		fmt.Println("Loop Array ", arr2[i])
	}

	// Slices
	slice := []int{1, 2, 3}
	fmt.Println("slice1", slice)

	slice = append(slice, 4)
	fmt.Println("assign slice2", slice)

	slice2 := slice[1:]
	fmt.Println("slice value in slice2", slice2)

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	z := append(x, y...)
	fmt.Println("append slice", z)
}
