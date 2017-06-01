// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func (c *Client) ListOrgHooks(user string) ([]*Hook, error) {
	hooks := make([]*Hook, 0, 10)
	return hooks, c.getParsedResponse("GET", fmt.Sprintf("/orgs/%s/hooks", user), nil, nil, &hooks)
}

func (c *Client) CreateOrgHook(user string, opt CreateHookOption) (*Hook, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	h := new(Hook)
	return h, c.getParsedResponse("POST", fmt.Sprintf("/orgs/%s/hooks", user), jsonHeader, bytes.NewReader(body), h)
}

func (c *Client) EditOrgHook(user string, id int64, opt EditHookOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PATCH", fmt.Sprintf("/orgs/%s/hooks/%d", user, id), jsonHeader, bytes.NewReader(body))
	return err
}

func (c *Client) DeleteOrgHook(user string, id int64) error {
	_, err := c.getResponse("DELETE", fmt.Sprintf("/orgs/%s/hooks/%d", user, id), nil, nil)
	return err
}

var (
	_ Payloader = &OrgCreatePayload{}
	_ Payloader = &OrgPushPayload{}
	_ Payloader = &OrgPullRequestPayload{}
)

// _________                        __
// \_   ___ \_______   ____ _____ _/  |_  ____
// /    \  \/\_  __ \_/ __ \\__  \\   __\/ __ \
// \     \____|  | \/\  ___/ / __ \|  | \  ___/
//  \______  /|__|    \___  >____  /__|  \___  >
//         \/             \/     \/          \/

type OrgCreatePayload struct {
	Secret       string        `json:"secret"`
	Ref          string        `json:"ref"`
	RefType      string        `json:"ref_type"`
	Organization *Organization `json:"organization"`
	Repo         *Repository   `json:"repository"`
	Sender       *User         `json:"sender"`
}

func (p *OrgCreatePayload) SetSecret(secret string) {
	p.Secret = secret
}

func (p *OrgCreatePayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// ParseCreateHook parses create event hook content.
func ParseOrgCreateHook(raw []byte) (*OrgCreatePayload, error) {
	hook := new(OrgCreatePayload)
	if err := json.Unmarshal(raw, hook); err != nil {
		return nil, err
	}

	// it is possible the JSON was parsed, however,
	// was not from Gogs (maybe was from Bitbucket)
	// So we'll check to be sure certain key fields
	// were populated
	switch {
	case hook.Organization == nil:
		return nil, ErrInvalidReceiveHook
	case len(hook.Ref) == 0:
		return nil, ErrInvalidReceiveHook
	}
	return hook, nil
}

// __________             .__
// \______   \__ __  _____|  |__
//  |     ___/  |  \/  ___/  |  \
//  |    |   |  |  /\___ \|   Y  \
//  |____|   |____//____  >___|  /
//                      \/     \/

// OrgPushPayload represents a payload information of push event.
type OrgPushPayload struct {
	Secret       string           `json:"secret"`
	Ref          string           `json:"ref"`
	Before       string           `json:"before"`
	After        string           `json:"after"`
	CompareURL   string           `json:"compare_url"`
	Commits      []*PayloadCommit `json:"commits"`
	Organization *Organization    `json:"organization"`
	Repo         *Repository      `json:"repository"`
	Pusher       *User            `json:"pusher"`
	Sender       *User            `json:"sender"`
}

func (p *OrgPushPayload) SetSecret(secret string) {
	p.Secret = secret
}

func (p *OrgPushPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

// ParsePushHook parses push event hook content.
func ParseOrgPushHook(raw []byte) (*OrgPushPayload, error) {
	hook := new(OrgPushPayload)
	if err := json.Unmarshal(raw, hook); err != nil {
		return nil, err
	}

	switch {
	case hook.Repo == nil:
		return nil, ErrInvalidReceiveHook
	case len(hook.Ref) == 0:
		return nil, ErrInvalidReceiveHook
	}
	return hook, nil
}

// Branch returns branch name from a payload
func (p *OrgPushPayload) Branch() string {
	return strings.Replace(p.Ref, "refs/heads/", "", -1)
}

// __________      .__  .__    __________                                     __
// \______   \__ __|  | |  |   \______   \ ____  ________ __   ____   _______/  |_
//  |     ___/  |  \  | |  |    |       _// __ \/ ____/  |  \_/ __ \ /  ___/\   __\
//  |    |   |  |  /  |_|  |__  |    |   \  ___< <_|  |  |  /\  ___/ \___ \  |  |
//  |____|   |____/|____/____/  |____|_  /\___  >__   |____/  \___  >____  > |__|
//                                     \/     \/   |__|           \/     \/

// OrgPullRequestPayload represents a payload information of pull request event.
type OrgPullRequestPayload struct {
	Secret       string          `json:"secret"`
	Action       HookIssueAction `json:"action"`
	Index        int64           `json:"number"`
	Changes      *ChangesPayload `json:"changes,omitempty"`
	PullRequest  *PullRequest    `json:"pull_request"`
	Organization *Organization   `json:"organization"`
	Repository   *Repository     `json:"repository"`
	Sender       *User           `json:"sender"`
}

func (p *OrgPullRequestPayload) SetSecret(secret string) {
	p.Secret = secret
}

func (p *OrgPullRequestPayload) JSONPayload() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}
