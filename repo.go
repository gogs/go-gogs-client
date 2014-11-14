// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

// Permission represents a API permission.
type Permission struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

// Repository represents a API repository.
type Repository struct {
	Id          int64      `json:"id"`
	Owner       User       `json:"owner"`
	FullName    string     `json:"full_name"`
	Private     bool       `json:"private"`
	Fork        bool       `json:"fork"`
	HtmlUrl     string     `json:"html_url"`
	CloneUrl    string     `json:"clone_url"`
	SshUrl      string     `json:"ssh_url"`
	Permissions Permission `json:"permissions"`
}

// ListMyRepos lists all repositories for the authenticated user that has access to.
func (c *Client) ListMyRepos() ([]*Repository, error) {
	repos := make([]*Repository, 0, 10)
	err := c.getParsedResponse("GET", "/user/repos", nil, nil, &repos)
	return repos, err
}
