package ch08

import (
	"fmt"
	"regexp"
)

// references
// ------------------------------------------------------------
// http://c.biancheng.net/view/5124.html

func UsingRegularExpression() {
	expression := `^[0-9-].*$`
	reg := regexp.MustCompile(expression)
	if reg == nil {
		panic(fmt.Sprintf("Invild regexp: %s", expression))
	}

	id1 := "009-001-33533223-45238"
	if reg.Match([]byte(id1)) {
		fmt.Printf("id 1 matched format: %s\n", id1)
	} else {
		fmt.Printf("id 1 does not match format: %s\n", id1)
	}

	id2 := "d09-c01-b3533223-a5238"
	if reg.Match([]byte(id2)) {
		fmt.Printf("id 2 matched format: %s\n", id2)
	} else {
		fmt.Printf("id 2 does not match format: %s\n", id2)
	}
}
