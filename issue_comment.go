package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// CommentType is type of a comment
type CommentType int

// Comment represents a comment in commit and issue page.
type Comment struct {
	ID       int64       `json:"id"`
	Type     CommentType `json:"type"`
	Poster   *User       `json:"poster"`
	IssueID  int64       `json:"issue_id"`
	CommitID int64       `json:"commit_id"`
	Line     int64       `json:"line"`
	Content  string      `json:"content"`

	Created     time.Time `json:"created"`
	CreatedUnix int64     `json:"created_unix"`

	// Reference issue in commit message
	CommitSHA string `json:"commit_sha"`

	//Attachments []*Attachment `json:"attachments"`
}

// ListRepoIssueComments list comments on an issue
func (c *Client) ListRepoIssueComments(owner, repo string, issueID int64) ([]*Comment, error) {
	comments := make([]*Comment, 0, 10)
	return comments, c.getParsedResponse("GET", fmt.Sprintf("/repos/%s/%s/issues/%d/comments", owner, repo, issueID), nil, nil, &comments)
}

// CreateIssueCommentOption is option when creating an issue comment
type CreateIssueCommentOption struct {
	Content string `json:"content" binding:"required"`
}

// CreateIssueComment create comment on an issue
func (c *Client) CreateIssueComment(owner, repo string, issueID int64, opt CreateIssueCommentOption) (*Comment, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	comment := new(Comment)
	return comment, c.getParsedResponse("POST", fmt.Sprintf("/repos/:%s/:%s/issues/%d/comments", owner, repo, issueID),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), comment)
}

// EditIssueCommentOption is option when editing an issue comment
type EditIssueCommentOption struct {
	Content string `json:"content" binding:"required"`
}

// EditIssueComment edits an issue comment
func (c *Client) EditIssueComment(owner, repo string, issueID, commentID int64, opt EditIssueCommentOption) (*Comment, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	comment := new(Comment)
	return comment, c.getParsedResponse("PATCH", fmt.Sprintf("/repos/:%s/:%s/issues/%d/comments/%d", owner, repo, issueID, commentID),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body), comment)
}
