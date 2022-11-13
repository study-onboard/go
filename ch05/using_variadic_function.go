package ch05

import "fmt"

func UsingVariadicFunction() {
	showOnMonitor(
		"Lana", "Kut", "Mona", "Bowen", "Leo", "Ray", "Albert", "Jim", "Tom", "James", "Lucy",
	)
}

func showOnMonitor(members ...string) {
	fmt.Printf("type of members: %T\n", members)

	for index, member := range members {
		fmt.Printf("%d. %s\n", index, member)
	}
}
