package entity

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

//Cделать проверку что введеной вещи еще нет в бд

type Thing struct {
	Thing         string
	Color         string
	Number        int
	DateOfWashing string
}

var YearNow, MonthNow, DayNow = time.Now().Date()
var TimeNow = strconv.Itoa(DayNow) + "-" + strconv.Itoa(int(MonthNow)) + "-" + strconv.Itoa(YearNow)

func ThingsArrToString(Things []Thing) string {
	var sb strings.Builder
	for _, u := range Things {
		sb.WriteString(fmt.Sprintf("%s %s %d шт, Стиралась %s\n", u.Thing, u.Color, u.Number, u.DateOfWashing))
	}
	res := sb.String()
	return res
}

func ThingsFromColors(Things []Thing, color string) []string {
	res := make([]string, 0)
	for _, u := range Things {
		switch color {
		case "black":
			if slices.Contains(Black_colored, u.Color) {
				res = append(res, u.Thing)
			}
		case "white":
			if slices.Contains(White_colored, u.Color) {
				res = append(res, u.Thing)
				fmt.Println(res)
			}
		case "colored":
			if u.Color != "black" && u.Color != "white" {
				res = append(res, u.Thing)
			}
		case "All":
			res = append(res, u.Thing)
		}
	}
	fmt.Println(res)
	return res
}

// Для AddClothes
func StringToThingArr(clothes string, id int64) ([]Thing, error) {
	var res [][]string
	var things []Thing

	for _, val := range strings.Split(clothes, ",") {
		if val == "" || !strings.Contains(val, "-") {
			return nil, errors.New("Значени введены не верно")
		} else {
			strings.Trim(val, " ")
			res = append(res, strings.Split(val, "-"))
		}
	}

	for i := range res {
		num, _ := strconv.Atoi(res[i][2])

		u := &Thing{
			Thing:         res[i][0],
			Color:         res[i][1],
			Number:        num,
			DateOfWashing: TimeNow,
		}
		things = append(things, *u)

	}

	return things, nil
}

func WashedUpdate() (Thing *Thing, id int64) {
	return nil, 0
}
