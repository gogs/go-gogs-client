// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
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
