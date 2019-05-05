package repository

import "fmt"

type PSQLStub struct {
	Users []User
}

func (p PSQLStub)GetAllUsers() ([]User, error){
	return p.Users,
}
