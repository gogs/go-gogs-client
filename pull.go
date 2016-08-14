// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"time"
)

// PullRequest represents a pull reqesut API object.
type PullRequest struct {
	// Copied from issue.go
	ID        int64      `json:"id"`
	Index     int64      `json:"number"`
	State     StateType  `json:"state"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	User      *User      `json:"user"`
	Labels    []*Label   `json:"labels"`
	Milestone *Milestone `json:"milestone"`
	Assignee  *User      `json:"assignee"`
	Comments  int        `json:"comments"`

	Mergeable      *bool      `json:"mergeable"`
	HasMerged      bool       `json:"merged"`
	Merged         *time.Time `json:"merged_at"`
	MergedCommitID *string    `json:"merge_commit_sha"`
	MergedBy       *User      `json:"merged_by"`
}
