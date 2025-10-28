package main

import (
	"context"
	"fmt"
	"test/exercise_users_http/service"
)

func main() {

	ctx := context.Background()

	serviceApiCall := service.ServiceApiCall{}

	users, err := serviceApiCall.GetUsers(ctx)
	if err != nil {
		panic(err)
	}

	posts, err := serviceApiCall.GetPosts(ctx)
	if err != nil {
		panic(err)
	}

	mapUsers := serviceApiCall.UsersToMap(users)

	for _, post := range posts {
		if value, ok := mapUsers[post.UserID]; ok {
			mapUsers[post.UserID] = value + 1
		}
	}

	for _, user := range users {
		if value, ok := mapUsers[user.Id]; ok {
			fmt.Printf("user: %s \t posts quanty %d\n", user.Name, value)
		}
	}
}
