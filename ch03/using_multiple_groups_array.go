package ch03

import "fmt"

func UsingMultipleGroupsArray() {
  teams := [3][2] string {
    {"Tom", "James"},
    {"Jerry", "Lucy"},
    {"Lana", "Jane"},
  }

  for i, team := range teams {
    fmt.Printf("Team %d: \n", i)
    fmt.Println("---------------------------------------")
    for j, member := range team {
      fmt.Printf("%d. %s\n", j, member) 
    }
    fmt.Println()
  }
}
