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

  // test slice copy
  testSliceCopy()
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
  staffs := []string {
    "Mona Liu",
    "Leo Liu",
    "Sigma Liu",
  }

  fmt.Println("----- Test initialize slice ---------------------------------")
  fmt.Printf("Length: %d\n", len(staffs))
  fmt.Println("Staffs: ")
  for i, staff := range staffs {
    fmt.Printf("%d. %s\n", i, staff)
  }
  fmt.Println()
}

func testSliceCopy() {
  var items []string
  items = append(items, "Linux")
  items = append(items, "OSX")
  items = append(items, "Windows")
  items = append(items, "FreeBSD")

  var cloneItems []string
  copy(cloneItems, items)

  fmt.Println("----- Test slice copy --------------------------------------")
  for i, item := range cloneItems {
    fmt.Printf("%d. %s\n", i, item)
  }

  cloneItems = make([]string, len(items))
  copy(cloneItems, items)
  fmt.Println("----- Test slice copy --------------------------------------")
  for i, item := range cloneItems {
    fmt.Printf("%d. %s\n", i, item)
  }
}
