package main

import "fmt"

func main() {
	Person := Information{}
	Person.ContactInfo.Email = "pavelvasilev24843@gmail.com"
	Person.ContactInfo.Number = "+79276650544"
	fmt.Println(Person.Output())
	SendMessage(Person)

}

type Information struct {
	Name        string
	Age         uint
	ContactInfo PersonalContact
}

type PersonalContact struct {
	Email  string
	Number string
}

func (r Information) Output() string {
	return fmt.Sprintf("Электронная почта пользователя: %s", r.ContactInfo.Email)
}

func (m Information) Message() string {
	return fmt.Sprintf("Телефонный номер пользователя: %s", m.ContactInfo.Number)
}

type MessageforUser interface {
	Output() string
	Message() string
}

func SendMessage(n MessageforUser) {
	fmt.Println(n.Message())
}

type SetUp struct {
	Install string
}
