/*
A slice is just like an array which is a container to hold elements of the same data type but slice can vary in size.

-Syntax to define a slice is pretty similar to that of an array but without specifying the elements count. Hence s is a slice.
var s []int

s := arr[1:3]

- create slice using make too.
  make([] type, len, capacity)

- slice is a reference to an array

- The capacity of a slice is the number of elements it can hold and len is total element in slice

- automatic expansion of capacity

- slice is a struct
type slice struct {
    zerothElement *type
    len int
    cap int
}

- slices are still passed by value to the function but since they reference the same underneath the array, it looks like that they are passed by reference.

-Multi-dimensional slices same as array
  1 := [][]int{
    []int{1, 2},
    []int{3, 4},
    []int{5, 6},
}
// type inference like arrays
s2 := [][]int{
    {1, 2},
    {3, 4},
    {5, 6},
}

*/

package main

import (
	"fmt"
	//"reflect"
)

func main() {

	// var intSlice = make([]int, 10)        
	// var strSlice = make([]string, 10, 20) // when length and capacity is different

	// fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	// fmt.Println(reflect.ValueOf(intSlice).Kind())

	// fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	// fmt.Println(reflect.ValueOf(strSlice).Kind())

	// Creating an array
	// arr := [7]string{"This", "is", "the", "tutorial",
	// 	"of", "Go", "language"}

	// // Display array
	// fmt.Println("Array:", arr)

	// // Creating a slice
	// myslice := arr[:3]

	// // Display slice
	// fmt.Println("Slice:", myslice)
	// myslice = append(myslice, "2", "3", "4")
	// // Display length of the slice
	// fmt.Printf("Length of the slice: %d\n", len(myslice))
  // fmt.Println("Array:", arr)
	// // Display the capacity of the slice
	// fmt.Printf("\nCapacity of the slice: %d \n", cap(myslice))
	
	// myslice = append(myslice, "2", "3", "4", "2", "3", "4")
	// fmt.Println("Slice:", myslice)
  // // Display length of the slice
	// fmt.Printf("Length of the slice: %d, cap : %d", len(myslice),cap(myslice))
  // fmt.Println("\nArray:", arr)

	// // Display the capacity of the slice
	// fmt.Printf("\nCapacity of the slice: %d \n", cap(myslice))

	// fmt.Println("Comparison of slices")
	// // creating slices
	// s1 := []int{12, 34, 56}
	// // Display length of the slice
	// fmt.Printf("Length of the slice: %d", len(s1))

	// // Display the capacity of the slice
	// fmt.Printf("\nCapacity of the slice: %d \n", cap(s1))

	// var s2 []int

	// // Checking if the given slice is nil or not
	// fmt.Println(s1 == nil)
	// fmt.Println(s2 == nil)

	fmt.Println("\nMulti -D slices !!")
	// Creating multi-dimensional slice
	s3 := [][]int{{12, 34,32},
		{56, 47},
		{29, 40},
		{46, 78},
	}

	// Accessing multi-dimensional slice
	fmt.Println("Slice 1 : ", s3)

	// // Creating multi-dimensional slice
	// s4 := [][]string{
	// 	[]string{"Geeks", "for"},
	// 	[]string{"Geeks", "GFG"},
	// 	[]string{"gfg", "geek"},
	// }

	// // Accessing multi-dimensional slice
	// fmt.Println("Slice 2 : ", s4)
}
