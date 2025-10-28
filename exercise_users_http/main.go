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

	mapUsers := serviceApiCall.UsersToPostsMap(users, posts)

	highestPost := mapUsers[users[0].Id].Posts
	userName := ""
	counterPosts := 0

	for _, user := range mapUsers {
		if highestPost < user.Posts {
			highestPost = user.Posts
			userName = user.Name
		} else if highestPost == user.Posts {
			counterPosts++
		}
	}

	if userName == "" && counterPosts == len(mapUsers) {
		fmt.Printf("all the users added the same amount of posts")
	} else {
		fmt.Printf("the user with the highest post is: %s with %d posts", userName, highestPost)
	}
}
