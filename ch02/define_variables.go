package ch02

import "fmt"

func DefineVariables()  {
  // define single variable
  var message string = "I had sent you some notification for system down"
  fmt.Println("message: " + message)
  
  // define multiple variables
  var min, max int64
  min = 60
  max = 80
  fmt.Printf("min: %d, max: %d\n", min, max)

  // define multiple variables 2
  var (
    name string
    age int8
    position string
    remark string
  )
  name = "kut"
  age = 42
  position = "IT Archetech"
  remark = "OK Monitor"
  fmt.Printf("name: %s, age: %d, position: %s, remark: %s\n", name, age, position, remark)


  // short style
  manager := "Lana Change"
  company := "Sanlea comp"
  fmt.Printf("Manager: %s, company: %s\n", manager, company)
}
