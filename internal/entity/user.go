package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Cделать проверку что введеной вещи еще нет в бд

// Переделать с User на Thing или Clothe
type User struct {
	Thing         string
	Color         string
	Number        int
	DateOfWashing string
}

var YearNow, MonthNow, DayNow = time.Now().Date()
var TimeNow = strconv.Itoa(DayNow) + "-" + strconv.Itoa(int(MonthNow)) + "-" + strconv.Itoa(YearNow)

func UsersArrToString(users []User) string {
	var sb strings.Builder
	for _, u := range users {
		sb.WriteString(fmt.Sprintf("%s %s %d шт, Стиралась %s\n", u.Thing, u.Color, u.Number, u.DateOfWashing))
	}
	res := sb.String()
	return res
}

func ThingsFromColors(users []User, color string) []string {
	res := make([]string, 0)
	for _, u := range users {
		if color == "All" {
			res = append(res, u.Thing)
		} else if color == "colored" {
			if u.Color != "black" && u.Color != "white" {
				res = append(res, u.Thing)
			}
		} else if u.Color == color {
			res = append(res, u.Thing)
		}

	}
	return res
}

func StringToUserArr(clothes string, id int64) ([]User, error) {
	var res [][]string
	var things []User

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

		u := &User{
			Thing:         res[i][0],
			Color:         res[i][1],
			Number:        num,
			DateOfWashing: TimeNow,
		}
		things = append(things, *u)

	}

	return things, nil
}
