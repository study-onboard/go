package ch03

import (
	"container/list"
	"fmt"
)

func UsingList() {
	members := list.New()
	fmt.Printf("members's type is: %T\n", members)

	members.PushBack("Kut Zhang")
	element := members.PushBack("Lana Chang")
	members.PushBack("Qing Zhang")
	members.PushBack("Qin Zhang")

	members.InsertAfter("Hahaha", element)
	members.InsertBefore("Hohoho", element)

	for member := members.Front(); member != nil; member = member.Next() {
		fmt.Printf("%s\n", member.Value)
	}
}
