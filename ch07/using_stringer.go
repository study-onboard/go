package ch07

import "fmt"

func UsingStringer() {
	commander := &Commander{
		name:    "Kut Zhang",
		comment: "Star ship commander",
	}
	fmt.Println(commander)
}

type Commander struct {
	name    string
	comment string
}

func (commander *Commander) String() string {
	return fmt.Sprintf("Commander - name: %s, comment: %s", commander.name, commander.comment)
}
