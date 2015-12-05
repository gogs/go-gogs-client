// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
)

// User represents a API user.
type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"username"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

type CreateUserOption struct {
	Name        string `json:"name" binding:"Required;AlphaDashDot;MaxSize(35)"`
	Email 		string `json:"email" binding:"Required;Email;MaxSize(254)"`
	Password    string `json:"password" binding:"Required;MaxSize(255)"`
}

func (c *Client) GetUserInfo(user string) (*User, error) {
	u := new(User)
	err := c.getParsedResponse("GET", fmt.Sprintf("/users/%s", user), nil, nil, u)
	return u, err
}

func (c *Client) CreateUser(opt CreateUserOption) (*User, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	u := new(User)
	return u, c.getParsedResponse("POST", "/users", http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), u)
}
