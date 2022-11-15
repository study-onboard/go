package ch08

import (
	"fmt"
	"time"
)

func UsingTimePackage() {
	simpleUseCase()
	operateUseCase()
	timerUseCase()
}

func simpleUseCase() {
	now := time.Now()
	fmt.Printf("Time now: %T - %s\n", now, now)
	fmt.Printf("Timestamp: %T - %v\n", now.Unix(), now.Unix())

	timestamp := now.Unix()
	aTime := time.Unix(timestamp, 0)
	year := aTime.Year()
	month := aTime.Month()
	day := aTime.Day()
	hour := aTime.Hour()
	minute := aTime.Minute()
	second := aTime.Second()
	fmt.Println(aTime)
	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)

	weekday := aTime.Weekday().String()
	fmt.Printf("Weekday: %s\n", weekday)
}

func operateUseCase() {

}

func timerUseCase() {
	timer := time.Tick(time.Second)
	i := 0
	for nextTime := range timer {
		if i > 20 {
			break
		}

		fmt.Println(nextTime)
		i += 1
	}
}
