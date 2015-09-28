// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"fmt"
	"time"
)

type Sha1 struct {
	Sha1  string             	  `json:"sha1"`
}

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

type DiffLine struct {
	LeftIdx  int		`json:"left_idx"`
	RightIdx int		`json:"right_idx"`
	Type     string		`json:"type"`
	Content  string		`json:"content"`
}

type DiffSection struct {
	Name  string		`json:"name"`
	Lines []*DiffLine	`json:"lines"`
}

type DiffFile struct {
	Name               string				`json:"name"`
	Index              int					`json:"index"`
	Addition		   int 					`json:"addition"`
	Deletion 		   int					`json:"deletion"`
	Type               string				`json:"type"`
	IsCreated          bool					`json:"created"`
	IsDeleted          bool					`json:"deleted"`
	IsBin              bool					`json:"bin"`
	Sections           []*DiffSection		`json:"sections"`
}

type Diff struct {
	TotalAddition		int 				`json:"total_addition"`
	TotalDeletion 		int					`json:"total_deletion"`
	Files				[]*DiffFile			`json:"files"`
}

func (c *Client) CommitByID(user, repo, id string) (*Commit, error) {
	commit := new(Commit)
	return commit, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/commits/%s", user, repo, id), nil, nil, commit)
}
