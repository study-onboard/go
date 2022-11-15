package ch08

import "github.com/study-onboard/go/ch08/singleton"

func UsingSingleton() {
	singleton.GetNotificationService().Notify("Show me the money")
}
