// credit: https://medium.com/better-programming/a-real-world-example-of-go-interfaces-98e89b2ddb67

package main

import "fmt"

type User struct {
	Name      string
	Email     string
	Notifiers []UserNotifier
}

type UserNotifier interface {
	SendMessage(user *User, message string) error
}

func (user *User) notify(message string) {
	for _, notifier := range user.Notifiers {
		notifier.SendMessage(user, message)
	}
}

type EmailNotifier struct {
}

func (notifier EmailNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending email to %s with content %s\n", user.Name, message)
	return err
}

type SmsNotifier struct {
}

func (notifier SmsNotifier) SendMessage(user *User, message string) error {
	_, err := fmt.Printf("Sending SMS to %s with content %s\n", user.Name, message)
	return err
}

func mainUserNotifier() {
	user := User{"Gavin Belson", "gavin.belson@hooli.com", []UserNotifier{EmailNotifier{}, SmsNotifier{}}}

	user.notify("Welcome!")
}
