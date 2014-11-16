// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrInvalidReceiveHook = errors.New("Invalid JSON payload received over webhook")
)

type Hook struct {
	Id     int64             `json:"id"`
	Type   string            `json:"type"`
	Events []string          `json:"events"`
	Active bool              `json:"active"`
	Config map[string]string `json:"config"`
}

func (c *Client) ListRepoHooks(user, repo string) ([]*Hook, error) {
	hooks := make([]*Hook, 0, 10)
	err := c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/hooks", user, repo), nil, nil, &hooks)
	return hooks, err
}

type CreateHookOption struct {
	Type   string            `json:"type"`
	Config map[string]string `json:"config"`
	Active bool              `json:"active"`
}

func (c *Client) CreateRepoHook(user, repo string, opt CreateHookOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("POST", fmt.Sprintf("/repos/%s/%s/hooks", user, repo),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body))
	return err
}

type EditHookOption struct {
	Config map[string]string `json:"config"`
	Active bool              `json:"active"`
}

func (c *Client) EditRepoHook(user, repo string, id int64, opt EditHookOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PATCH", fmt.Sprintf("/repos/%s/%s/hooks/%d", user, repo, id),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body))
	return err
}

type PayloadAuthor struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

type PayloadCommit struct {
	Id      string         `json:"id"`
	Message string         `json:"message"`
	Url     string         `json:"url"`
	Author  *PayloadAuthor `json:"author"`
}

type PayloadRepo struct {
	Id          int64          `json:"id"`
	Name        string         `json:"name"`
	Url         string         `json:"url"`
	Description string         `json:"description"`
	Website     string         `json:"website"`
	Watchers    int            `json:"watchers"`
	Owner       *PayloadAuthor `json:"owner"`
	Private     bool           `json:"private"`
}

// Payload represents a payload information of hook.
type Payload struct {
	Secret     string           `json:"secret"`
	Ref        string           `json:"ref"`
	Commits    []*PayloadCommit `json:"commits"`
	Repo       *PayloadRepo     `json:"repository"`
	Pusher     *PayloadAuthor   `json:"pusher"`
	Before     string           `json:"before"`
	After      string           `json:"after"`
	CompareUrl string           `json:"compare_url"`
}

// ParseHook parses Gogs webhook content.
func ParseHook(raw []byte) (*Payload, error) {
	hook := new(Payload)
	if err := json.Unmarshal(raw, hook); err != nil {
		return nil, err
	}
	// it is possible the JSON was parsed, however,
	// was not from Github (maybe was from Bitbucket)
	// So we'll check to be sure certain key fields
	// were populated
	switch {
	case hook.Repo == nil:
		return nil, ErrInvalidReceiveHook
	case len(hook.Ref) == 0:
		return nil, ErrInvalidReceiveHook
	}
	return hook, nil
}

// Branch returns branch name from a payload
func (h *Payload) Branch() string {
	return strings.Replace(h.Ref, "refs/heads/", "", -1)
}
