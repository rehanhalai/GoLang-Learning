package main

import (
	"fmt"
)

func main() {

	age := 20

	// if age > 18 {
	// 	fmt.Println("person is adult")
	// } else {
	// 	fmt.Println("not eligible")
	// }

	if age > 18 {
		fmt.Println("person is adult")
	} else if age >= 12 {
		fmt.Println("teenager")
	} else {
		fmt.Println("child")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	} else {
		fmt.Println("NO")
	}

	// GoLang does not have ternary operators

	// switch case
	i := 3

	switch i {
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("default")
	}

	// multiple condition switch

	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// fmt.Println("its weekand")
	// }

	// type switch

	whoAmI := func(i interface{}) {
		switch i := i.(type) {
		case int:
			fmt.Println("integer")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("other", i)
		}
	}

	whoAmI("go")
}
