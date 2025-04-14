package array

import "fmt"

func Maps() {
	var A map[int]string = map[int]string{

		1: "stage",
		2: "two",
		3: "out",
		4: "come",
	}
	(delete(A, 1))
	fmt.Println(A)
	fmt.Println(A[1])

	//map with  make
	B := make(map[int]string)

	B[1] = "A"
	B[2] = "Z"
	fmt.Println(B)

}
