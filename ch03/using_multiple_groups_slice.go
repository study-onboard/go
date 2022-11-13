package ch03

import "fmt"


func UsingMultipleGroupsSlice() {
  groups := [][]string {
    {
      "show me the money",
    },
    {
      "how do you turn this on",
      "make love, no war",
    },
  }
  dumpGroups("Initialize", groups)

  // append
  groups[0] = append(groups[0], "good good study, day day up")
  dumpGroups("After append", groups)
}


func dumpGroups(title string, groups [][]string) {
  fmt.Println(title)
  fmt.Println("=========================================================")

  for i, group := range groups {
    fmt.Printf("--- Group: %d -------------------------------------------\n", i)
    for j, message := range group {
      fmt.Printf("message %d: %s\n", j, message)
    }
    fmt.Println()
  }

  fmt.Println()
}
