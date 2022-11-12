package ch02

import "fmt"

func VariableScope() {
  // global scope
  fmt.Printf("Global author name: %s\n", AuthorName)
  
  // local scope
  localScopeTest()

  // function scope
  functionVariableScope("Juggler Zhang")
}

func localScopeTest() {
  var AuthorName = "Magick"
  fmt.Printf("Local author name: %s\n", AuthorName)
}

func functionVariableScope(authorName string) {
  fmt.Printf("Function author name: %s\n", authorName)
}
