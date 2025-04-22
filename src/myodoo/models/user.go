// models/user.go
package models

import (
	"context"
	"encoding/json"
	"fmt"
	"myodoo/redis"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(user User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("user:%s", user.ID)
	return redis.Client.Set(context.Background(), key, data, 0).Err()
}

func GetUser(id string) (User, error) {
	var user User
	val, err := redis.Client.Get(context.Background(), fmt.Sprintf("user:%s", id)).Result()
	if err != nil {
		return user, err
	}

	err = json.Unmarshal([]byte(val), &user)
	return user, err
}

