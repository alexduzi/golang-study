package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"test/exercise_users_http/model"
)

// https://jsonplaceholder.typicode.com/users

// https://jsonplaceholder.typicode.com/posts

const (
	URL_USERS string = "https://jsonplaceholder.typicode.com/users"
	URL_POSTS string = "https://jsonplaceholder.typicode.com/posts"
)

type ServiceApiCall struct {
}

func (s *ServiceApiCall) GetUsers(ctx context.Context) ([]model.User, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", URL_USERS, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	users := []model.User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *ServiceApiCall) UsersToMap(users []model.User) map[int32]int32 {
	mapUsers := map[int32]int32{}

	for _, user := range users {
		mapUsers[user.Id] = 0
	}

	return mapUsers
}

func (s *ServiceApiCall) GetPosts(ctx context.Context) ([]model.Post, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", URL_POSTS, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	posts := []model.Post{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
