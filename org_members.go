// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gogs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddOrgMembershipOption struct {
	Role    string `json:"role"`
}

func (c *Client) AddOrgMembership(orgname string, username string, opt AddOrgMembershipOption) error {
	body, err := json.Marshal(&opt)
	if err != nil {
		return err
	}
	_, err = c.getResponse("PUT", fmt.Sprintf("/orgs/%s/membership/%s", orgname, username),
		http.Header{"content-type": []string{"application/json"}}, bytes.NewReader(body))
	return err
}
