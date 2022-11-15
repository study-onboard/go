package protected

import (
	"fmt"

	"github.com/google/uuid"
)

type account struct {
	id   string
	name string
}

func NewAccount(name string) *account {
	return &account{
		id:   uuid.NewString(),
		name: name,
	}
}

func (account account) GetId() string {
	return account.id
}

func (account account) GetName() string {
	return account.name
}

func (account account) Dump() {
	fmt.Printf("account - %s\n", account)
}
