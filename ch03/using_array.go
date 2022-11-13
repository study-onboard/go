package ch03

import "fmt"

func UsingArray() {
  var members [3]string
  members[0] = "show me the money"
  members[1] = "how do you turn this on"
  members[2] = "good good study, day day up"
  dumpMembers(members)

  // initialize method 1
  members = [3]string {
    "Kut Zhang",
    "Lana Chang",
    "Qing Zhang",
  }
  dumpMembers(members)

  // initialize method 2
  members = [...]string {
    "Jucylee",
    "Bowen",
    "Ray",
  }
  dumpMembers(members)

  org := [3]string {
    "Kut Zhang",
    "Lana Chang",
    "Qing Zhang",
  }
  fmt.Printf("org == members: %t\n", org == members)

  org = [3]string {
    "Jucylee",
    "Bowen",
    "Ray",
  }
  fmt.Printf("org == members: %t\n", org == members)

  // wrong syntax of arrays
  // ----------------------------------------------------
  // managers := [4]string {
  //   "Lana",
  //   "Lucy",
  //   "Lily",
  //   "Magick",
  // }
  // fmt.Printf("org == managers: %t\n", managers == org)
}

func dumpMembers(members [3]string) {
  fmt.Println("---------------------------------------------------")
  for index, member := range members {
    fmt.Printf("%d: %s\n", index, member)
  }
}
