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
}

func UsersArrToString(users []User) string {
	var sb strings.Builder
	for _, u := range users {
		sb.WriteString(fmt.Sprintf("Thing: %s, Color: %s, Количество: %d\n", u.Thing, u.Color, u.Number))
	}
	res := sb.String()
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
		}
		things = append(things, *u)

	}

	return things, nil
}
