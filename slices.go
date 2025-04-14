package array

import "fmt"

func Slices() {
	slice := make([]int, 2, 3)
	sli := append(slice, 54, 34, 94)
	fmt.Printf("print value of slice:%v\n", sli)

	slice_another := []string{"len", "height"}

	fmt.Println(slice_another)

	//slice range
	rangeone := slice[1:3]
	rangetwo := slice[1:]
	rangethree := slice[:2]

	fmt.Println(rangeone, rangetwo, rangethree)

}
