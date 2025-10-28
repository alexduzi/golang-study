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

func (s *ServiceApiCall) UsersToPostsMap(users []model.User, posts []model.Post) map[int32]model.UserHighPost {
	mapUsers := map[int32]model.UserHighPost{}

	for _, user := range users {
		postsByUser := filter(posts, func(post model.Post) bool {
			return user.Id == post.UserID
		})

		mapUsers[user.Id] = model.UserHighPost{
			Id:    user.Id,
			Name:  user.Name,
			Posts: int32(len(postsByUser)),
		}
	}

	return mapUsers
}

func (s *ServiceApiCall) GetPosts(ctx context.Context) ([]model.Post, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", URL_POSTS, nil)
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

	posts := []model.Post{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func filter[E any](s []E, f func(E) bool) []E {
	filtered := make([]E, 0, len(s)) // Pre-allocate with potential capacity for efficiency
	for _, e := range s {
		if f(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
