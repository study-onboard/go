package ch06

import (
	"fmt"

	"github.com/google/uuid"
)

func UsingStruct() {
	structSimplyUse()
	fmt.Println("---------------------------------------------------------")

	structFactoryMethodUse()
	fmt.Println("---------------------------------------------------------")

	structInherideUse()
	fmt.Println("---------------------------------------------------------")
}

// -----------------------------------------------------------------------

func structSimplyUse() {
	accountPtr := &Account{
		id:   uuid.NewString(),
		name: "Kut Zhang",
	}
	accountPtr.login("ABCabcXYZxyz")

	staffPtr := &Staff{
		id:       uuid.NewString(),
		account:  accountPtr,
		position: "IT Manager",
	}
	staffPtr.work()

	staffPtr = &Staff{
		id: uuid.NewString(),
		account: &Account{
			id:   uuid.NewString(),
			name: "Masaga",
		},
		position: "Boss",
	}
	staffPtr.work()
}

func structFactoryMethodUse() {
	accountPtr := newAccount("Lana Chang")
	accountPtr.login("xYzAbC")
	staffPtr := newStaff(accountPtr, "Worker")
	staffPtr.work()
}

func structInherideUse() {
	cat := newBlackCart("James", "Lucy Liu")
	cat.miao()
}

// -----------------------------------------------------------------------

type Account struct {
	id   string
	name string
}

func (account *Account) login(password string) {
	fmt.Printf("Account Login: %s - %s - %s\n", account.id, account.name, password)
}

func newAccount(name string) *Account {
	account := new(Account)
	account.id = uuid.NewString()
	account.name = name
	return account
}

// -----------------------------------------------------------------------

type Staff struct {
	id       string
	account  *Account
	position string
}

func (staff *Staff) work() {
	fmt.Printf("Staff work - %s(%s)[%s]\n", staff.account.name, staff.id, staff.position)
}

func newStaff(accountPtr *Account, position string) *Staff {
	staff := new(Staff)
	staff.id = uuid.NewString()
	staff.account = accountPtr
	staff.position = position
	return staff
}

// -----------------------------------------------------------------------

type Cat struct {
	name  string
	color string
}

type BlackCat struct {
	Cat
	wizardName string
}

func (cat *BlackCat) miao() {
	fmt.Printf("Cat miao - %s - %s - %s\n", cat.name, cat.color, cat.wizardName)
}

func newBlackCart(name string, wizardName string) *BlackCat {
	cat := new(BlackCat)
	cat.name = name
	cat.color = "black"
	cat.wizardName = wizardName
	return cat
}
