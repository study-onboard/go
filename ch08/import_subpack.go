package ch08

import (
	"fmt"
	remote "sanlea/study/ch08/network"
	"sanlea/study/ch08/protected"
	_ "sanlea/study/ch08/subpack"
)

func ImportSubpack() {
	fmt.Println("-------------------------------------------------------------------")

	fmt.Println("show me the money")

	fmt.Println("-------------------------------------------------------------------")

	// test package alias
	remote.Request("http://localhost:8080/user", "POST", &map[string]any{
		"name": "Kut Zhang",
		"age":  105,
	})

	fmt.Println("-------------------------------------------------------------------")

	remote.CheckHealth("192.168.10.111", 5689)

	fmt.Println("-------------------------------------------------------------------")

	account := protected.NewAccount("Wendy Wu")
	fmt.Printf("account type: %T\n", account)
	fmt.Printf("account id: %s\n", account.GetId())
	fmt.Printf("account name: %s\n", account.GetName())
	account.Dump()
}
