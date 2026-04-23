package entity

import (
	"fmt"
	"time"
)

type Day struct {
	DayNow time.Time
}

func NewDate() *Day {
	return &Day{
		DayNow: time.Now(),
	}
}

func DaySinceLast(day time.Time) {
	newtime := time.Since(day)
	fmt.Println(newtime)
}
