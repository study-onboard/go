package ch02

import "fmt"

func UsingConst() {
  const message string = "show me the money"
  fmt.Printf("message: %s\n", message)


  // multiple consts
  const (
    slogon = "how do you turn this on?"
    manager = "Kut Zhang"
    staff = "Lana Change"
  )
  fmt.Printf("slogon: %s\nmanager: %s\nstaff: %s\n", slogon, manager, staff)

  // const generator
  const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Seturday
  )
  fmt.Printf("Sunday: %d\n", Sunday)
  fmt.Printf("Monday: %d\n", Monday)
  fmt.Printf("Tuesday: %d\n", Tuesday)
  fmt.Printf("Wednesday: %d\n", Wednesday)
  fmt.Printf("Thursday: %d\n", Thursday)
  fmt.Printf("Friday: %d\n", Friday)
  fmt.Printf("Seturday: %d\n", Seturday)
}
