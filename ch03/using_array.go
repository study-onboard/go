package ch03

import "fmt"

func UsingArray() {
  var members [3]string
  members[0] = "show me the money"
  members[1] = "how do you turn this on"
  members[2] = "good good study, day day up"
  iterateMembers(members)

  // initialize method 1
  members = [3]string {
    "Kut Zhang",
    "Lana Chang",
    "Qing Zhang",
  }
  iterateMembers(members)

  members = [...]string {
    "Jucylee",
    "Bowen",
    "Ray",
  }
  iterateMembers(members)
}

func iterateMembers(members [3]string) {
  fmt.Println("---------------------------------------------------")
  for index, member := range members {
    fmt.Printf("%d: %s\n", index, member)
  }
}
