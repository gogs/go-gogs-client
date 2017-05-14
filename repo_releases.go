// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Release struct {
	ID 					int64             `json:"id"`
	Publisher			User			  `json:"publisher"`
	TagName				string			  `json:"tag_name"`
	LowerTagName		string			  `json:"lower_tag_name"`
	Target				string		  	  `json:"target"`
	Title				string		  	  `json:"title"`
	Sha1				string		  	  `json:"sha1"`
	NumCommits			int		  	  	  `json:"num_commits"`
	NumCommitsBehind	int		  	  	  `json:"num_commits_behind"`
	Note				string		  	  `json:"note"`
	IsDraft				bool		  	  `json:"draft"`
	IsPrerelease		bool		  	  `json:"prerelease"`
	Created				time.Time		  `json:"created"`
}

func (c *Client) ReleaseByName(user, repo, name string) (*Release, error) {
	release := new(Release)
	return release, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/releases/%s", user, repo, name), nil, nil, release)
}

func (c *Client) ListReleases() ([]*Release, error) {
	releases := make([]*Release, 0, 10)
	err := c.getParsedResponse("GET", "/repos/%s/%s/releases", nil, nil, &releases)
	return releases, err
}

type CreateReleaseOption struct {
	TagName				string			  `json:"tag_name"`
	Target				string		  	  `json:"target"`
	Title				string		  	  `json:"title"`
	Note				string		  	  `json:"note"`
	IsDraft				bool		  	  `json:"draft"`
	IsPrerelease		bool		  	  `json:"prerelease"`
}


func (c *Client) CreateRelease(user string, repo string, opt CreateReleaseOption) (*Release, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	release := new(Release)
	return release, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/releases", user, repo),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), release)
}