package golang

import (
	"fmt"
)

type Interval interface {
	Travel()
}

type User struct {
	Name     string
	Interval Interval
}

func NewUser(name string, interval Interval) *User {
	return &User{
		Name:     name,
		Interval: interval,
	}
}

func (user User) travel() {
	fmt.Printf("%s", user.Name)
	user.Interval.Travel()
}

type Train struct{}

func (train Train) Travel() {
	fmt.Println("搭火車")
}

type Airplane struct{}

func (airplane Airplane) Travel() {
	fmt.Println("搭飛機")
}
