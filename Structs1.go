package array

import "fmt"

type Name struct {
	salary   int
	name     string
	DayOfpay int
}

func A() {
	name := Name{
		salary:   10,
		name:     "NM",
		DayOfpay: 250,
	}

	fmt.Println(name)

}
