package main

import (
	"context"
	"fmt"

	prisma "github.com/jorgeAM/prisma/src/generated/prisma-client"
)

var client = prisma.New(nil)
var ctx = context.TODO()

func main() {
	users, err := getAllUsers()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", users)
}

func createNewUser(name string) (*prisma.User, error) {
	return client.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(ctx)
}

func getAllUsers() ([]prisma.User, error) {
	return client.Users(nil).Exec(ctx)
}
