// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// PullRequest represents a pull request API object.
type PullRequest struct {
	ID        int64      `json:"id"`
	Index     int64      `json:"number"`
	Poster    *User      `json:"user"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Labels    []*Label   `json:"labels"`
	Milestone *Milestone `json:"milestone"`
	Assignee  *User      `json:"assignee"`
	State     StateType  `json:"state"`
	Comments  int        `json:"comments"`

	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`

	Mergeable      bool       `json:"mergeable"`
	HasMerged      bool       `json:"merged"`
	Merged         *time.Time `json:"merged_at"`
	MergedCommitID *string    `json:"merge_commit_sha"`
	MergedBy       *User      `json:"merged_by"`

	Base      *PRBranchInfo `json:"base"`
	Head      *PRBranchInfo `json:"head"`
	MergeBase string        `json:"merge_base"`
}

type PRBranchInfo struct {
	Name       string      `json:"label"`
	Ref        string      `json:"ref"`
	Sha        string      `json:"sha"`
	RepoID     int64       `json:"repo_id"`
	Repository *Repository `json:"repo"`
}

type ListPullRequestsOptions struct {
	Page  int    `json:"page"`
	State string `json:"state"`
}

func (c *Client) ListRepoPullRequests(owner, repo string, opt ListPullRequestsOptions) ([]*PullRequest, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	prs := make([]*PullRequest, 0, 10)
	return prs, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/pulls", owner, repo), jsonHeader, bytes.NewReader(body), &prs)
}

func (c *Client) GetPullRequest(owner, repo string, index int64) (*PullRequest, error) {
	pr := new(PullRequest)
	return pr, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/pulls/%d", owner, repo, index), nil, nil, pr)
}

type CreatePullRequestOption struct {
	Head      string  `json:"head" binding:"Required"`
	Base      string  `json:"base" binding:"Required"`
	Title     string  `json:"title" binding:"Required"`
	Body      string  `json:"body"`
	Assignee  string  `json:"assignee"`
	Milestone int64   `json:"milestone"`
	Labels    []int64 `json:"labels"`
}

func (c *Client) CreatePullRequest(owner, repo string, opt CreatePullRequestOption) (*PullRequest, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	pr := new(PullRequest)
	return pr, c.getParsedResponse("POST", fmt.Sprintf("/repos/%s/%s/pulls", owner, repo),
		jsonHeader, bytes.NewReader(body), pr)
}

type EditPullRequestOption struct {
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	Assignee  string  `json:"assignee"`
	Milestone int64   `json:"milestone"`
	Labels    []int64 `json:"labels"`
	State     *string `json:"state"`
}

func (c *Client) EditPullRequest(owner, repo string, index int64, opt EditPullRequestOption) (*PullRequest, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	pr := new(PullRequest)
	return pr, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/%s/%s/issues/%d", owner, repo, index),
		jsonHeader, bytes.NewReader(body), pr)
}

func (c *Client) MergePullRequest(owner, repo string, index int64) error {
	_, err := c.getResponse("POST", fmt.Sprintf("/repos/%s/%s/pulls/%d/merge", owner, repo, index), nil, nil)
	return err
}

func (c *Client) IsPullRequestMerged(owner, repo string, index int64) (bool, error) {
	statusCode, err := c.getStatusCode("GET", fmt.Sprintf("/repos/%s/%s/pulls/%d/merge", owner, repo, index), nil, nil)

	if err != nil {
		return false, err
	}

	return statusCode == 204, nil

}
