// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"fmt"
	"time"
)

type Signature struct {
	Email   string            `json:"email"`
	Name    string            `json:"name"`
	When 	time.Time         `json:"when"`
}

type Commit struct {
	ID      		string             	   `json:"id"`
	Author			Signature			   `json:"author"`
	Committer		Signature			   `json:"committer"`
	CommitMessage	string			   	   `json:"commit_message"`
}

func (c *Client) CommitById(user, repo, id string) (*Commit, error) {
	commit := new(Commit)
	return commit, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/commits/%s", user, repo, id), nil, nil, commit)
}
