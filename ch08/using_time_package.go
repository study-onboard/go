package ch08

import (
	"fmt"
	"time"
)

func UsingTimePackage() {
	simpleUseCase()
	fmt.Println("----------------------------------------------------------------------")

	operateUseCase()
	fmt.Println("----------------------------------------------------------------------")

	timerUseCase()
	fmt.Println("----------------------------------------------------------------------")

	formatUseCase()
	fmt.Println("----------------------------------------------------------------------")

	parseUseCase()
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
	now := time.Now()
	fmt.Printf("Now: %s\n", now)

	aTime := now.Add(time.Hour * 5)
	fmt.Printf("After add 5 hours: %s\n", aTime)

	aTime = now.Add(time.Hour * -5)
	fmt.Printf("After sub 5 hours: %s\n", aTime)

	hours := now.Sub(aTime) / time.Hour
	fmt.Printf("now - aTime = %d hours\n", hours)

	aTime = now.Add(time.Hour * -3)
	bTime := now.Add(time.Hour * 2)
	switch {
	case aTime.Equal(bTime):
		fmt.Println("aTime equals bTime")

	case aTime.Before(bTime):
		fmt.Println("aTime before bTime")

	case aTime.After(bTime):
		fmt.Println("aTime after bTime")
	}
}

func timerUseCase() {
	timer := time.Tick(time.Second)
	i := 0
	for nextTime := range timer {
		if i > 5 {
			break
		}

		fmt.Println(nextTime)
		i += 1
	}
}

func formatUseCase() {
	now := time.Now()
	fmt.Printf("Now for custom format: %s\n", now.Format("2006-01-02 15:04:05"))
}

func parseUseCase() {
	layout := "2006-01-02 15:04:05"
	text := "2022-11-01 21:30:45"
	aTime, err := time.Parse(layout, text)
	if err != nil {
		fmt.Printf("Invalid time text: %s\n", text)
	} else {
		fmt.Printf("aTime: %s\n", aTime)
	}

	aTime, err = time.ParseInLocation(layout, text, time.Local)
	if err != nil {
		fmt.Printf("Invalid time text: %s\n", text)
	} else {
		fmt.Printf("aTime: %s\n", aTime)
	}
}
