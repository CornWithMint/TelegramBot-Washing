package entity

import (
	"fmt"
	"strings"
)

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
