// Copyright 2017 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Release represents a release API object.
type Release struct {
	ID              int64         `json:"id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Body            string        `json:"body"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	Author          *User         `json:"author"`
	Created         time.Time     `json:"created_at"`
	URL             string        `json:"url"`
	TarURL          string        `json:"tarball_url"`
	ZipURL          string        `json:"zipball_url"`
	Attachments     []*Attachment `json:"assets"`
}

// ListReleases list releases of a repository
func (c *Client) ListReleases(user, repo string) ([]*Release, error) {
	releases := make([]*Release, 0, 10)
	err := c.getParsedResponse("GET",
		fmt.Sprintf("/repos/%s/%s/releases", user, repo),
		nil, nil, &releases)
	return releases, err
}

// GetRelease get a release of a repository
func (c *Client) GetRelease(user, repo string, id int64) (*Release, error) {
	r := new(Release)
	err := c.getParsedResponse("GET",
		fmt.Sprintf("/repos/%s/%s/releases/%d", user, repo, id),
		nil, nil, &r)
	return r, err
}

// DeleteRelease delete a release from a repository
func (c *Client) DeleteRelease(user, repo string, id int64) error {
	_, err := c.getResponse("DELETE",
		fmt.Sprintf("/repos/%s/%s/releases/%d", user, repo, id),
		nil, nil)
	return err
}

// CreateReleaseOption options when creating a release
type CreateReleaseOption struct {
	// required: true
	TagName      string `json:"tag_name" binding:"Required"`
	Target       string `json:"target_commitish"`
	Title        string `json:"name"`
	Note         string `json:"body"`
	IsDraft      bool   `json:"draft"`
	IsPrerelease bool   `json:"prerelease"`
}

// CreateRelease create a release
func (c *Client) CreateRelease(user, repo string, form CreateReleaseOption) (*Release, error) {
	body, err := json.Marshal(form)
	if err != nil {
		return nil, err
	}
	r := new(Release)
	err = c.getParsedResponse("POST",
		fmt.Sprintf("/repos/%s/%s/releases", user, repo),
		jsonHeader, bytes.NewReader(body), r)
	return r, err
}
