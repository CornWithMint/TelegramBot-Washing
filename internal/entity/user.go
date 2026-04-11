package entity

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//Cделать проверку что введеной вещи еще нет в бд

type User struct {
	Thing  string
	Color  string
	Number int
	Days   int
}

func UsersArrToString(users []User) string {
	var sb strings.Builder
	for _, u := range users {
		sb.WriteString(fmt.Sprintf("%s %s %d шт, Дней со стирки %d\n", u.Thing, u.Color, u.Number, u.Days))
	}
	res := sb.String()
	return res
}

func UsersArrToColor(users []User, color string) []string {
	res := make([]string, 0)
	for _, u := range users {
		if u.Color == color {
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
			Thing:  res[i][0],
			Color:  res[i][1],
			Number: num,
			Days:   0,
		}
		things = append(things, *u)

	}

	return things, nil
}
