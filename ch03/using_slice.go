package ch03

import (
	"fmt"
)

/*
Slice just like List of Java
*/
func UsingSlice() {
  members := [5]string {
    "Tom",
    "James",
    "Lana",
    "Jane",
    "Lucy",
  }

  // index base - begin, end
  // min 0
  // max len(array) 
  // begin: include
  // end: exclude
  //
  leaders := members[1:3]
  // will display: James, Lana - index: 1, 2 (exclude 3)
  for i, leader := range leaders {
    fmt.Printf("%d: %s\n", i, leader)
  }

  fmt.Println("-------------------------------------------------")
  fmt.Println()

  // define slice
  var managers [3]string = [3]string {
    "Kut Zhang",
    "Bowen Liu",
    "Ray Liu",
  }
  fmt.Printf("managers: %s\n", managers)

  // test initialize slice
  testInitializeSlice()

  // test make slice size
  testMakeSliceSize(3, managers)
  testMakeSliceSize(0, managers)
}

func testMakeSliceSize(size int, initializeItems [3]string) {
  items := make([]string, size)
  for _, item := range initializeItems {
    items = append(items, item)
  }

  fmt.Printf("------ Test for size: %d -----------------------------\n", size)
  for i, item := range items {
    fmt.Printf("%d: %s\n", i, item)
  }
  fmt.Println()
}

func testInitializeSlice() {
  var staffs []string
  staffs = append(staffs, "Mona Liu")
  staffs = append(staffs, "Leo Liu")
  staffs = append(staffs, "Sigma Liu")

  fmt.Println("----- Test initialize slice ---------------------------------")
  for i, staff := range staffs {
    fmt.Printf("%d. %s\n", i, staff)
  }
  fmt.Println()
}
