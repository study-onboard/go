package ch03

import (
	"fmt"
	"sort"
)

func UsingMap() {
  var staffs map[string]string = map[string]string {
    "Boss": "Lana Chang",
    "IT Manager": "Kut Zhang",
    "Worker": "Qing Zhang",
  }
  for position, staff := range staffs {
    fmt.Printf("%s: %s\n", position, staff)
  }
  fmt.Println()

  var groups = map[string][]string {
    "Magical": {
      "Lana", "Qing", "Mona",
    },
    "Airo": {
      "Juggler", "Oxl", "Kut", "Lunar",
    },
  }
  printGroups("Before append", groups)

  // append member
  groups["Magical"] = append(groups["Magical"], "Borland")
  printGroups("After append", groups)


  // sort print map
  sortPrintMap(groups)

  // delete map item
  deleteMapItem()
}

// print groups
func printGroups(title string, groups map[string][]string) {
  fmt.Println(title)
  fmt.Println("================================================================")

  for group, members := range groups {
    fmt.Printf("Group: %s\n", group)
    fmt.Println("---------------------------")
    for i, member := range members {
      fmt.Printf("%d. %s\n", i, member)
    }
    fmt.Println()
  }
  
  fmt.Println()
}

// print map sortly
func sortPrintMap(groups map[string][]string) {
  printGroups("Before sort", groups)

  var keys []string
  for key := range groups {
    keys = append(keys, key)
  }

  sort.Strings(keys)
  fmt.Println("After sort")
  fmt.Println("================================================================")

  for _, key := range keys {
    fmt.Printf("Group: %s\n", key)
    fmt.Println("---------------------------")

    members := groups[key]
    sort.Strings(members)
    for i, member := range members {
      fmt.Printf("%d. %s\n", i, member)
    }
    fmt.Println()
  }

  fmt.Println()
}

// delete item of map
func deleteMapItem() {
  positions := map[string]string {
    "Boss": "Bosh King",
    "CEO": "Mana J Ku",
    "LineManager": "Kut Zhang",
  }
  
  fmt.Println("Before delete item")
  fmt.Println("-----------------------------------------------------")
  for position, staff := range positions {
    fmt.Printf("%s: %s\n", position, staff)
  }
  fmt.Println()

  delete(positions, "LineManager")
  fmt.Println("Before delete item")
  fmt.Println("-----------------------------------------------------")
  for position, staff := range positions {
    fmt.Printf("%s: %s\n", position, staff)
  }
}
